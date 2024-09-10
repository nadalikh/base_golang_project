package validation

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
)

type ValidationService interface {
	Validate(interface{}) error
	RegisterCustomValidations()
}

type validationServiceImpl struct {
	validator *validator.Validate
}

func NewValidationService() ValidationService {
	vs := &validationServiceImpl{
		validator: validator.New(),
	}
	vs.RegisterCustomValidations() // Register custom validations on initialization
	return vs
}

func (s *validationServiceImpl) Validate(obj interface{}) error {
	err := s.validator.Struct(obj)
	if err != nil {
		var invalidValidationError *validator.InvalidValidationError
		if errors.As(err, &invalidValidationError) {
			return err
		}
		for _, err := range err.(validator.ValidationErrors) {
			return fmt.Errorf("این فیل ردیه شده است %s is not valid: %s", err.Field(), err.Tag())
		}
	}
	return nil
}

func (s *validationServiceImpl) RegisterCustomValidations() {
	s.validator.RegisterValidation("existedUser", func(fl validator.FieldLevel) bool {
		_ = fl.Field().String()
		return true
	})
}
