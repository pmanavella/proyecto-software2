package users

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"errors"
	dao "users-api/dao"
	domain "users-api/domain/users"
	"utils" // Ajustar la importación para apuntar al paquete utils fuera de users-api
)

type Repository interface {
	GetAll() ([]dao.User, error)
	GetByID(id int64) (dao.User, error)
	GetByUsername(username string) (dao.User, error)
	Create(user dao.User) (int64, error)
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

func NewService(mainRepo, cacheRepo, memcachedRepo Repository, tokenizer Tokenizer) *Service {
	return &Service{
		mainRepository:      mainRepo,
		cacheRepository:     cacheRepo,
		memcachedRepository: memcachedRepo,
		tokenizer:           tokenizer,
	}
}

func (s *Service) GetAll() ([]domain.User, error) {
	users, err := s.mainRepository.GetAll()
	if err != nil {
		return nil, err
	}

	var domainUsers []domain.User
	for _, user := range users {
		domainUsers = append(domainUsers, domain.User{
			ID:       user.ID,
			Username: user.Username,
			Password: user.Password,
		})
	}

	return domainUsers, nil
}

func (s *Service) GetByID(id int64) (domain.User, error) {
	user, err := s.mainRepository.GetByID(id)
	if err != nil {
		return domain.User{}, err
	}

	return domain.User{
		ID:       user.ID,
		Username: user.Username,
		Password: user.Password,
	}, nil
}

func (s *Service) Create(user domain.User) (int64, error) {
	// Encriptar la contraseña antes de guardar
	user.Password = utils.EncryptPassword(user.Password)

	daoUser := dao.User{
		Username: user.Username,
		Password: user.Password,
	}

	return s.mainRepository.Create(daoUser)
}

func (s *Service) Update(user domain.User) error {
	// Encriptar la contraseña antes de actualizar
	user.Password = utils.EncryptPassword(user.Password)

	daoUser := dao.User{
		ID:       user.ID,
		Username: user.Username,
		Password: user.Password,
	}

	return s.mainRepository.Update(daoUser)
}

func (s *Service) Delete(id int64) error {
	return s.mainRepository.Delete(id)
}

func (s *Service) Login(username string, password string) (domain.LoginResponse, error) {
	user, err := s.mainRepository.GetByUsername(username)
	if (err != nil) {
		return domain.LoginResponse{}, err
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return domain.LoginResponse{}, errors.New("invalid credentials")
	}

	token, err := s.tokenizer.GenerateToken(user.Username, user.ID)
	if err != nil {
		return domain.LoginResponse{}, err
	}

	return domain.LoginResponse{
		UserID:   user.ID,
		Username: user.Username,
		Token:    token,
	}, nil
}

func Hash(input string) string {
	hash := md5.Sum([]byte(input))
	return hex.EncodeToString(hash[:])
}

func (s *Service) convertUser(user dao.User) domain.User {
	return domain.User{
		ID:       user.ID,
		Username: user.Username,
		Password: user.Password,
	}
}