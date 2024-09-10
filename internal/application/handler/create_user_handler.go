package handler

import (
	"testGorm/internal/application/command"
	"testGorm/internal/application/validation"
	"testGorm/internal/domain/entity"
	"testGorm/internal/domain/repository"
)

type CreateUserCommandHandler interface {
	Handle(cmd command.CreateUserCommand) (*entity.User, error)
}

type createUserCommandHandlerImpl struct {
	UserRepository    repository.UserRepository
	validationService validation.ValidationService
}

func NewCreateUserCommandHandler(
	userRepository repository.UserRepository,
	validationService validation.ValidationService,
) CreateUserCommandHandler {
	return &createUserCommandHandlerImpl{
		UserRepository:    userRepository,
		validationService: validationService,
	}
}

func (h *createUserCommandHandlerImpl) Handle(cmd command.CreateUserCommand) (*entity.User, error) {
	user, err := entity.NewUser(cmd.Name, cmd.Email, cmd.Password)
	if err != nil {
		return nil, err
	}
	// Validate user
	if err := h.validationService.Validate(user); err != nil {
		return nil, err
	}

	// Save the user
	if err := h.UserRepository.Save(user); err != nil {
		return nil, err
	}

	return user, nil
}
