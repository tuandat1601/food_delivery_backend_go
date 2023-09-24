package restauranthandler

import (
	"food_delivery/common"
	restaurantbiz "food_delivery/internal/biz/restaurant_biz"

	"food_delivery/internal/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)
// GetRestaurant godoc
// @Summary Get infomation restaurant by ID
// @Description Get infomation restaurant by ID from database
// @Accept json
// @Produce json
// @Tags Restaurant
// @Param id path int true "ID of User"
// @Param restaurant_id path int true "ID of Restaurant"
// @Success 200 {object} model.Restaurant "Infomation of restaurant"
// @Failure 400 {object} map[string]string "Fail "
// @Router /v1/users/{id}/restaurant/{restaurant_id} [get]
func GetRestaurantById(db *gorm.DB) func(c *gin.Context) {
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

		store := repositories.NewSQLStore(db)

		business := restaurantbiz.NewGetRestaurant(store)
		result, err := business.GetNewRestaurant(c.Request.Context(), id, res_id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, common.NewDataResponse(http.StatusOK, result))

	}
}
