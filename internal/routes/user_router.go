package routes

import (
	"food_delivery/internal/transport/gin/users_handler"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(imncomingRoutes *gin.Engine, db *gorm.DB) {
	imncomingRoutes.POST("v1/users/signup", users_handler.SignUp(db))
	imncomingRoutes.POST("v1/users/signin", users_handler.SignIn(db))
}
