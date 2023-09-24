package users_handler

import (
	
	"food_delivery/internal/biz/users_biz"
	"food_delivery/internal/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)
// @Summary Xóa người dùng theo ID
// @Description Xóa người dùng theo ID
// @Produce json
// @Tags User
// @Param id path int true "ID của người dùng"
// @Success 200 {object} map[string]string "Deleted successfully"
// @Failure 400 {object} map[string]string "Fail"
// @Router /v1/users/:id [patch]
func DeletedUserById(db *gorm.DB) func(c * gin.Context){
	return func (c * gin.Context){
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store:=repositories.NewSQLStore(db)
		business:=users_biz.NewDeletedUser(store)
		
		
		if err:=business.DeletedUseById(c.Request.Context(),id);err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"err":err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK,gin.H{
			"message":"Deleted successfully",
		})

	}
}