package validation

type EmailValidationService interface {
	IsValidEmail(email string) bool
}
