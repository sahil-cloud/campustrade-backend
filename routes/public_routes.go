package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sahil-cloud/backend/controllers"
)

func PublicRoutes(router *gin.Engine) {
	router.POST("/login", controllers.Login())
	// router.POST("/signup", controllers.SignUp())
}
