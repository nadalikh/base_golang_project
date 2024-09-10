package services

import (
	"testGorm/internal/application/command"
	"testGorm/internal/application/handler"
	"testGorm/internal/application/validation"
	"testGorm/internal/domain/entity"
	"testGorm/internal/domain/repository"
)

type UserService interface {
	RegisterUser(name, email, password string) (*entity.User, error)
	FindByID(id string) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	Login(email, password string) (string, error)
}

type userServiceImpl struct {
	userRepository    repository.UserRepository
	createUserHandler handler.CreateUserCommandHandler
	loginUserHandler  handler.LoginUserCommandHandler
	validationService validation.ValidationService
}

func NewUserService(userRepository repository.UserRepository,
	createUserHandler handler.CreateUserCommandHandler,
	validationService validation.ValidationService,
	loginUserHandler handler.LoginUserCommandHandler,
) UserService {
	return &userServiceImpl{
		userRepository:    userRepository,
		createUserHandler: createUserHandler,
		loginUserHandler:  loginUserHandler,
		validationService: validationService,
	}
}

func (s *userServiceImpl) RegisterUser(name, email, password string) (*entity.User, error) {
	// Save the user
	cmd := command.CreateUserCommand{Name: name, Email: email, Password: password}
	return s.createUserHandler.Handle(cmd)
}

func (s *userServiceImpl) FindByID(id string) (*entity.User, error) {
	return s.userRepository.FindByID(id)
}
func (s *userServiceImpl) FindByEmail(email string) (*entity.User, error) {
	return s.userRepository.FindByEmail(email)
}
func (s *userServiceImpl) Login(email, password string) (string, error) {
	cmd := command.LoginUserCommand{Email: email, Password: password}
	//user, err := s.userRepository.FindByEmail(email)
	//if err != nil {
	//	return "", err
	//}
	//return s.loginUserHandler.Handle(cmd, user)
	return s.loginUserHandler.Handle(cmd)

}
