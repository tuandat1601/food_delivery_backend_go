package restauranthandler

import (
	"food_delivery/internal/model"
	"food_delivery/internal/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateRestaurant(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data model.Restaurants
		if err := c.ShouldBind(&data); err != nil {
			return
		}
		store := repositories.NewSQLStore(db)

		business := repositories.NewRestaurantUsecase(store)
		if err := business.CreateNewRestaurant(c.Request.Context(),&data);err!=nil{
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
	}
}
