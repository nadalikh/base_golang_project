package command

type LoginUserCommand struct {
	Email    string `json:"email" validate:"required,email,existedUser"`
	Password string `json:"password" validate:"required,min=6"`
}
