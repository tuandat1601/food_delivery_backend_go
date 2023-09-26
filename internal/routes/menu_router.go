package routes

import (
	menuhandler "food_delivery/internal/transport/gin/menu_handler"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func MenuRoutes(imncomingRoutes *gin.Engine, db *gorm.DB) {
	imncomingRoutes.POST("v1/:restaurant_id/menu/create", menuhandler.CreateMenu(db))
}