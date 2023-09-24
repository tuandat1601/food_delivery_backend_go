package routes

import (
	"food_delivery/internal/middleware"
	restauranthandler "food_delivery/internal/transport/gin/restaurant_handler"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RestaurantRoutes(imncomingRoutes *gin.Engine, db *gorm.DB) {
	imncomingRoutes.POST("v1/users/:id/restaurant/create", middleware.DeserializeUser(db), restauranthandler.CreateRestaurant(db))
	imncomingRoutes.GET("v1/users/:id/restaurant/:restaurant_id", restauranthandler.GetRestaurantById(db))
	imncomingRoutes.PATCH("v1/users/:id/restaurant/:restaurant_id/update", restauranthandler.UpdateRestaurantById(db))
	imncomingRoutes.PATCH("v1/users/:id/restaurant/:restaurant_id/deleted", restauranthandler.DeleteRestaurantById(db))
}
