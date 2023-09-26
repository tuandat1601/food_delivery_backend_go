package menuhandler

import (
	"food_delivery/common"
	menubiz "food_delivery/internal/biz/menu_biz"
	"food_delivery/internal/model"
	"food_delivery/internal/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateMenu(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var data model.Menu
		data.RestaurantID=id

		if err := c.ShouldBind(&data); err != nil {
			return
		}
		store := repositories.NewSQLStore(db)

		business := menubiz.NewCreateMenu(store)
		if err := business.CreateNewMenu(c.Request.Context(),&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, common.NewDataResponse(http.StatusOK, data))

	}
}
