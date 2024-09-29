package rest

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"testGorm/internal/application/services"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var userDTO struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := ctx.BindJSON(&userDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := c.userService.RegisterUser(userDTO.Name, userDTO.Email, userDTO.Password)
	if err != nil {
		var invalidValidationError *validator.InvalidValidationError
		if !errors.As(err, &invalidValidationError) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, user)
}

func (c *UserController) GetUserByID(ctx *gin.Context) {
	id := ctx.Param("id")

	user, err := c.userService.FindByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	ctx.JSON(http.StatusOK, user)
}
func (c *UserController) GetUserByEmail(ctx *gin.Context) {
	email := ctx.Query("email")

	user, err := c.userService.FindByEmail(email)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	ctx.JSON(http.StatusOK, user)
}
func (c *UserController) GetAll(ctx *gin.Context) {
	users, err := c.userService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, users)
}

func (c *UserController) Login(ctx *gin.Context) {
	var loginDTO struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := ctx.BindJSON(&loginDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	token, err := c.userService.Login(loginDTO.Email, loginDTO.Password)
	var invalidValidationError *validator.InvalidValidationError
	if !errors.As(err, &invalidValidationError) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
