package auth

import (
	dto "aseprayana-skeleton-go/dto/auth"
	m "aseprayana-skeleton-go/middlewares"
	r "aseprayana-skeleton-go/repositories/auth"
)

type AuthService interface {
	Profile(req dto.ProfileRequest) (*dto.ProfileResponse, error)
	Signin(req dto.AuthLoginRequest) (*dto.AuthLoginResponse, error)
	Signup(req dto.AuthRegisterRequest) (dto.AuthRegisterResponse, error)
}

type authService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewAuthService(Repo r.DomainRepository, jwtS m.JWTService) AuthService {
	return &authService{
		Repo: Repo,
		jwt:  jwtS,
	}
}
