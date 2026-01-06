package router

import (
	"github.com/gin-gonic/gin"
	"{{.ModuleName}}/internal/handler"
)

func RegisterUserRoutes(engine *gin.Engine, userHandler *UserHandler) {

	engine.POST("/user", userHandler.CreateUser)

	engine.GET("/user/:id", userHandler.GetUserByID)

	engine.PUT("/user/:id", userHandler.UpdateUser)

	engine.DELETE("/user/:id", userHandler.DeleteUser)
}
