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
// @Summary Lấy thông tin người dùng theo ID
// @Description Lấy thông tin người dùng theo ID
// @Produce json
// @Tags User
// @Param id path int true "ID của người dùng"
// @Success 200 {object} model.User "Thông tin người dùng"
// @Failure 400 {object} map[string]string "Lỗi lấy thông tin người dùng"
// @Router /v1/users/:id [get]
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