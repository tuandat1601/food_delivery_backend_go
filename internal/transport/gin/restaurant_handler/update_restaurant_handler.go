package restauranthandler

import (
	"food_delivery/common"
	restaurantbiz "food_delivery/internal/biz/restaurant_biz"
	"food_delivery/internal/model"
	"food_delivery/internal/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)
// UpdateRestaurant godoc
// @Summary Update infomation of restaurant
// @Description Update infomation of restaurant
// @Accept  json
// @Produce  json
// @Tags Restaurant 
// @Param name body string true "Name of restaurant"
// @Param address body string true "Address of restaurant"
// @Param phone body string true "phone number"
// @Param day_of_week body string false "Day open on week (default: 'monday to saturday')"
// @Param opening_time body string false "Time opening(default: '7AM')"
// @Param closing_time body string false "Time closing (default: '22PM')"
// @Success 200 {object} model.Restaurant "Successfully"
// @Failure 400 {object} map[string]string "fail"
// @Router /v1/users/{id}/restaurant/{restaurant_id}/update [patch]

func UpdateRestaurantById(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		res_id, errd := strconv.Atoi(c.Param("restaurant_id"))
		if errd != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errd.Error()})
			return
		}
		var data model.Restaurant
		
		
		if err := c.ShouldBind(&data); err != nil {
			return
		}
		store := repositories.NewSQLStore(db)

		business := restaurantbiz.NewUpdateRestaurant(store)
		if err := business.UpdateNewRestaurant(c.Request.Context(),id,res_id,&data);err!=nil{
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK,common.NewDataResponse(http.StatusOK, data))

	}
}
