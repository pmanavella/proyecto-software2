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
    GetUserByID(id int64) (dao.User, error)
    CreateUser(user dao.User) (int64, error)
    GetUserByEmail(email string) (dao.User, error)
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

func (service Service) GetUserByID(id int64) (domain.User, error) {
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
            if _, err := service.cacheRepository.CreateUser(user); err != nil {
                fmt.Printf("warning: error caching user in cache repository: %s\n", err.Error())
            }
            if _, err := service.memcachedRepository.CreateUser(user); err != nil {
                fmt.Printf("warning: error caching user in memcached repository: %s\n", err.Error())
            }
        } else {
            // Sí se encuentra en memcached, guardarlo en el caché repository
            if _, err := service.cacheRepository.CreateUser(user); err != nil {
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

            if _, err := service.cacheRepository.CreateUser(user); err != nil {
                fmt.Printf("warning: error caching user in cache repository: %s\n", err.Error())
            }

            if _, err := service.memcachedRepository.CreateUser(user); err != nil {
                fmt.Printf("warning: error caching user in memcached repository: %s\n", err.Error())
            }
        } else {
            fmt.Printf("Guardando en cache local: %v\n", user)
            if _, err := service.cacheRepository.CreateUser(user); err != nil {
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

func (service Service) CreateUser(registro domain.User) (int64, error) {
    // Hashear la contraseña
    passwordHash := Hash(registro.Password)

    // Crear nuevo usuario en la base de datos principal
    nuevoUser := dao.User{
        Username: registro.Username,
        Password: passwordHash,
        Nombre:   registro.Nombre,
        Apellido: registro.Apellido,
        Email:    registro.Email,
        Admin:    registro.Admin,
    }

    // Intentar crear el usuario en el repositorio principal
    id, err := service.mainRepository.CreateUser(nuevoUser)
    if err != nil {
        return 0, errores.NewInternalServerApiError("error creating user", err)
    }

    // Asignar el ID al nuevo usuario
    nuevoUser.UserID = id

    // Almacenar en caché (advertencia si falla, sin detener el flujo)
    if _, err := service.cacheRepository.CreateUser(nuevoUser); err != nil {
        fmt.Printf("warning: error caching new user: %s\n", err.Error())
    }

    // Almacenar en memcached (advertencia si falla, sin detener el flujo)
    if _, err := service.memcachedRepository.CreateUser(nuevoUser); err != nil {
        fmt.Printf("warning: error saving new user in memcached: %s\n", err.Error())
    }

    return id, nil
}