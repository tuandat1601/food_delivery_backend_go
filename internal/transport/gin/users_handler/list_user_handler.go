package users_handler

import (
	"food_delivery/common"
	"food_delivery/internal/biz/users_biz"
	"food_delivery/internal/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListUser(db *gorm.DB) func(c * gin.Context){
	return func (c * gin.Context){
		
		store:=repositories.NewSQLStore(db)
		business:=users_biz.NewGetListUser(store)
		conditions := map[string]interface{}{"is_deleted": false}
		result ,err:=business.ListUser(c.Request.Context(),conditions);
		if err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"err":err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK,common.NewDataResponse(http.StatusOK, result))

	}
}