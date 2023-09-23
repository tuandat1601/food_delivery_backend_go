package users_handler

import (
	"food_delivery/common"
	"food_delivery/internal/biz/users_biz"
	"food_delivery/internal/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetUserById(db *gorm.DB) func(c * gin.Context){
	return func (c * gin.Context){
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store:=repositories.NewSQLStore(db)
		business:=users_biz.NewGetUser(store)
		
		result ,err:=business.GetUserById(c.Request.Context(),id);
		if err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"err":err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK,common.NewDataResponse(http.StatusOK, result))

	}
}