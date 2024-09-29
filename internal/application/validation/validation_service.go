package validation

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"testGorm/internal/domain/repository"
)

type ValidationService interface {
	Validate(interface{}) error
	RegisterCustomValidations()
}

type validationServiceImpl struct {
	validator      *validator.Validate
	userRepository repository.UserRepository
}

func NewValidationService(userRepository repository.UserRepository) ValidationService {
	vs := &validationServiceImpl{
		validator:      validator.New(),
		userRepository: userRepository,
	}
	vs.RegisterCustomValidations() // Register custom validations on initialization
	return vs
}

// Validate function with better error handling for localization/custom messages
func (s *validationServiceImpl) Validate(obj interface{}) error {
	err := s.validator.Struct(obj)
	if err != nil {
		var invalidValidationError *validator.InvalidValidationError
		if errors.As(err, &invalidValidationError) {
			return err
		}
		for _, err := range err.(validator.ValidationErrors) {
			return fmt.Errorf("%s  %s", translateField(err.Field()), translateTag(err.Tag()))
		}
	}
	return nil
}

// A function to register custom validation rules
func (s *validationServiceImpl) RegisterCustomValidations() {
	// Custom validator example: "existedUser"
	s.validator.RegisterValidation("existedUser", func(fl validator.FieldLevel) bool {
		email := fl.Field().String()

		// Custom logic to check if the user exists in the database
		// Simulating with a hardcoded check for demo purposes.
		isTook, err := s.userRepository.IsEmailTaken(email)
		if err != nil {
			panic(err)
		}
		return isTook
	})
}

// Helper function to translate field names if needed
func translateField(field string) string {
	// Custom logic to translate fields or keep as is
	return field
}

// Helper function to provide custom error messages for each validation tag
func translateTag(tag string) string {
	switch tag {
	case "required":
		return "این فیلد اجباری است."
	case "email":
		return "ایمیل نامعتبر است."
	case "existedUser":
		return "این ایمیل قبلاً ثبت شده است."
	case "min":
		return "حداقل تعداد کاراکتر رعابت نشده است در"
	default:
		return tag
	}
}
