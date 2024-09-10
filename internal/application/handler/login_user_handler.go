package handler

import (
	"errors"
	"testGorm/internal/application/command"
	"testGorm/internal/application/validation"
	"testGorm/internal/domain/repository"
	"testGorm/internal/infrastructure/auth"
)

type LoginUserCommandHandler interface {
	Handle(cmd command.LoginUserCommand) (string, error)
}

type loginUserCommandHandlerImpl struct {
	UserRepository    repository.UserRepository
	validationService validation.ValidationService
	jwtService        auth.JWTService
}

func NewLoginUserCommandHandler(
	userRepository repository.UserRepository,
	validationService validation.ValidationService,
	jwtService auth.JWTService,
) LoginUserCommandHandler {
	return &loginUserCommandHandlerImpl{
		UserRepository:    userRepository,
		validationService: validationService,
		jwtService:        jwtService,
	}
}

func (h *loginUserCommandHandlerImpl) Handle(cmd command.LoginUserCommand) (string, error) {

	if err := h.validationService.Validate(cmd); err != nil {
		return "", err
	}

	user, err := h.UserRepository.FindByEmail(cmd.Email)
	if err != nil || user == nil {
		return "", errors.New("invalid email or password")
	}

	if err := user.VerifyPassword(cmd.Password); err != nil {
		return "", errors.New("invalid email or password")
	}

	// Generate JWT
	token, err := h.jwtService.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
