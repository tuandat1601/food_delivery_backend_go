package users_handler

import (
	
	"food_delivery/internal/biz/users_biz"
	"food_delivery/internal/model"
	"food_delivery/internal/repositories"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	
	"gorm.io/gorm"
)

func UpdateUserById(db *gorm.DB) func(c * gin.Context){
	return func (c * gin.Context){
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store:=repositories.NewSQLStore(db)
		business:=users_biz.NewUpdateUser(store)
		var data model.User
		if err := c.ShouldBind(&data); err != nil {
			return 
		}
		
		if err:=business.UpdateUserById(c.Request.Context(),id,&data);err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"err":err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK,gin.H{
			"message":"update successfully",
		})

	}
}