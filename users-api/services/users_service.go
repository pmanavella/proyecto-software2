package users

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	dao "users-api/dao"
	domain "users-api/domain/users"
	errores "users-api/extra"
)

type Repository interface {
	GetAll() ([]dao.User, error)
	GetUserByID(id int64) (dao.User, error)
	Create(user dao.User) (int64, error)
	GetUserByEmail(email string) (dao.User, error)
    Update(user dao.User) error
    Delete(id int64) error
}

type Tokenizer interface {
	GenerateToken(username string, userID int64) (string, error)
}

type Service struct {
	mainRepository      Repository
	cacheRepository     Repository
	memcachedRepository Repository
	tokenizer           Tokenizer
}

func NewService(mainRepository, cacheRepository, memcachedRepository Repository, tokenizer Tokenizer) Service {
	return Service{
		mainRepository:      mainRepository,
		cacheRepository:     cacheRepository,
		memcachedRepository: memcachedRepository,
		tokenizer:           tokenizer,
	}
}

func (service Service) GetAll() ([]domain.User, error) {
    users, err := service.mainRepository.GetAll()
    if err != nil {
        return nil, err
    }

    var domainUsers []domain.User
    for _, user := range users {
        domainUsers = append(domainUsers, domain.User{
            ID:       user.UserID,
            Username: user.Username,
            Password: user.Password,
            Nombre:   user.Nombre,
            Apellido: user.Apellido,
            Email:    user.Email,
            Admin:    user.Admin,
        })
    }

    return domainUsers, nil
}

func (service Service) GetByID(id int64) (domain.User, error) {
	// Intentar obtener el usuario desde el repositorio de caché
	user, err := service.cacheRepository.GetUserByID(id)
	if err != nil {
		fmt.Printf("warning: error getting user from cache repository: %s\n", err.Error())

		// Si no está en el caché repository, intentar obtenerlo desde memcached
		user, err = service.memcachedRepository.GetUserByID(id)
		if err != nil {
			fmt.Printf("warning: error getting user from memcached repository: %s\n", err.Error())

			// Si tampoco está en memcached, buscar en el repositorio principal (base de datos)
			user, err = service.mainRepository.GetUserByID(id)
			if err != nil {
				// Si el usuario no se encuentra en ninguno de los repositorios, devolver un error
				return domain.User{}, errores.NewBadRequestApiError("user not found")
			}

			// Guardar el usuario en el repositorio de caché y en memcached si fue encontrado en la base de datos principal
			if _, err := service.cacheRepository.Create(user); err != nil {
				fmt.Printf("warning: error caching user in cache repository: %s\n", err.Error())
			}
			if _, err := service.memcachedRepository.Create(user); err != nil {
				fmt.Printf("warning: error caching user in memcached repository: %s\n", err.Error())
			}
		} else {
			// Sí se encuentra en memcached, guardarlo en el caché repository
			if _, err := service.cacheRepository.Create(user); err != nil {
				fmt.Printf("warning: error caching user in cache repository: %s\n", err.Error())
			}
		}
	}

	// Verificar si el ID del usuario es 0, lo que indica que no se encontró el usuario
	if user.UserID == 0 {
		return domain.User{}, errores.NewBadRequestApiError("user not found")
	}

	// Aquí asigna los valores correspondientes del usuario al DTO
	userDto := domain.User{
		ID:       user.UserID,
		Username: user.Username,
		Password: user.Password,
		Nombre:   user.Nombre,
		Apellido: user.Apellido,
		Email:    user.Email,
		Admin:    user.Admin,
	}

	// Devolver el usuario
	return userDto, nil
}


func (service Service) Delete(id int64) error {
    if err := service.mainRepository.Delete(id); err != nil {
        return errores.NewInternalServerApiError("error deleting user", err)
    }

    // Eliminar de caché (advertencia si falla, sin detener el flujo)
    if err := service.cacheRepository.Delete(id); err != nil {
        fmt.Printf("warning: error deleting user from cache: %s\n", err.Error())
    }

    // Eliminar de memcached (advertencia si falla, sin detener el flujo)
    if err := service.memcachedRepository.Delete(id); err != nil {
        fmt.Printf("warning: error deleting user from memcached: %s\n", err.Error())
    }

    return nil
}

func (service Service) Login(email string, password string) (domain.LoginResponse, error) {
	passwordHash := Hash(password)

	user, err := service.cacheRepository.GetUserByEmail(email)
	if err != nil {
		fmt.Printf("warning: error getting user from cache repository: %s\n", err.Error())

		user, err = service.memcachedRepository.GetUserByEmail(email)
		if err != nil {
			fmt.Printf("warning: error getting user from memcached repository: %s\n", err.Error())

			user, err = service.mainRepository.GetUserByEmail(email)
			if err != nil {
				return domain.LoginResponse{}, fmt.Errorf("error getting user by email from main repository: %w", err)
			}

			if _, err := service.cacheRepository.Create(user); err != nil {
				fmt.Printf("warning: error caching user in cache repository: %s\n", err.Error())
			}

			if _, err := service.memcachedRepository.Create(user); err != nil {
				fmt.Printf("warning: error caching user in memcached repository: %s\n", err.Error())
			}
		} else {
			fmt.Printf("Guardando en cache local: %v\n", user)
			if _, err := service.cacheRepository.Create(user); err != nil {
				fmt.Printf("warning: error caching user in cache repository: %s\n", err.Error())
			}
		}
	} else {
		fmt.Printf("User found in cache: %v\n", user)
	}

	if user.Password != passwordHash {
		return domain.LoginResponse{}, fmt.Errorf("invalid credentials")
	}

	token, err := service.tokenizer.GenerateToken(user.Email, user.UserID)
	if err != nil {
		return domain.LoginResponse{}, fmt.Errorf("error generating token: %w", err)
	}

	return domain.LoginResponse{
		UserID:   user.UserID,
		Username: user.Username,
		Token:    token,
		Admin:    user.Admin,
	}, nil
}

func Hash(input string) string {
	hash := md5.Sum([]byte(input))
	return hex.EncodeToString(hash[:])
}

//create user

func (service Service) Create(user domain.User) (int64, error) {
    user.Password = Hash(user.Password)

    newUser := dao.User{
        Username: user.Username,
        Password: user.Password,
        Nombre:   user.Nombre,
        Apellido: user.Apellido,
        Email:    user.Email,
        Admin:    user.Admin,
    }

    id, err := service.mainRepository.Create(newUser)
    if err != nil {
        return 0, err
    }

    return id, nil
}

func (service Service) Update(user domain.User) error {
    user.Password = Hash(user.Password)

    updatedUser := dao.User{
        UserID:   user.ID,
        Username: user.Username,
        Password: user.Password,
        Nombre:   user.Nombre,
        Apellido: user.Apellido,
        Email:    user.Email,
        Admin:    user.Admin,
    }

    if err := service.mainRepository.Update(updatedUser); err != nil {
        return err
    }

    // Actualizar en caché (advertencia si falla, sin detener el flujo)
    if err := service.cacheRepository.Update(updatedUser); err != nil {
        fmt.Printf("warning: error updating user in cache: %s\n", err.Error())
    }

    // Actualizar en memcached (advertencia si falla, sin detener el flujo)
    if err := service.memcachedRepository.Update(updatedUser); err != nil {
        fmt.Printf("warning: error updating user in memcached: %s\n", err.Error())
    }

    return nil
}