package repositories

import (
	dto "aseprayana-skeleton-go/dto/auth"
	"aseprayana-skeleton-go/entity"

	"gorm.io/gorm"
)

type DomainRepository interface {
	Signup(req dto.AuthRegisterRequest) (dto.AuthRegisterResponse, error)
	Signin(req dto.AuthLoginRequest) (*entity.User, error)
	Profile(req dto.ProfileRequest) (*entity.User, error)
}

type authRepository struct {
	DB *gorm.DB
}

func NewAuthRepository(DB *gorm.DB) DomainRepository {
	return &authRepository{
		DB: DB,
	}
}
