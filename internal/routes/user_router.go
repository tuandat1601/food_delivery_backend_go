package routes

import (
	"food_delivery/internal/middleware"
	"food_delivery/internal/transport/gin/users_handler"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(imncomingRoutes *gin.Engine, db *gorm.DB) {
	imncomingRoutes.POST("v1/users/signup", users_handler.SignUp(db))
	imncomingRoutes.POST("v1/users/signin", users_handler.SignIn(db))
	imncomingRoutes.GET("v1/users/logout", users_handler.LogOut(db))
	imncomingRoutes.GET("v1/users/:id", users_handler.GetUserById(db))
	imncomingRoutes.PATCH("v1/users/:id", users_handler.DeletedUserById(db))
	imncomingRoutes.GET("v1/users/list",middleware.DeserializeUser(db), users_handler.ListUser(db))
	imncomingRoutes.GET("v1/authen", middleware.DeserializeUser(db))
	imncomingRoutes.PATCH("v1/users/:id/update", middleware.DeserializeUser(db),users_handler.UpdateUserById(db))
}
