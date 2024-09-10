package http

import (
	"github.com/gin-gonic/gin"
	"testGorm/internal/interfaces/rest"
)

func SetupRoutes(userController *rest.UserController) *gin.Engine {
	r := gin.Default()
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"error": "صفحه یافت نشد"})
	})
	// Define routes
	userRoute(r, userController)

	return r
}

func userRoute(engine *gin.Engine, userController *rest.UserController) {
	engine.POST("/users", userController.CreateUser)
	engine.GET("/users/:id", userController.GetUserByID)
	engine.GET("/users/", userController.GetUserByEmail)
	engine.POST("/login", userController.Login)
}
