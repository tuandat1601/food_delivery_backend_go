package users_handler

import (
	"food_delivery/common"
	"food_delivery/internal/biz/users_biz"
	"food_delivery/internal/model"
	"food_delivery/internal/repositories"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	
	"gorm.io/gorm"
)
// @Summary Đăng ký người dùng
// @Description Đăng ký một người dùng mới
// @Param user body model.User true "Thông tin người dùng"
// @Accept json
// @Produce json
// @Tags User
// @Success 200 {object} model.User "Người dùng đã đăng ký thành công"
// @Failure 400 {object} map[string]string "Lỗi đăng ký người dùng"
// @Router /v1/users/signup [post]
func SignUp(db *gorm.DB) func( c *gin.Context) {
	return func(c *gin.Context) {
		var data model.User
		if err := c.ShouldBind(&data); err != nil {
			return 
		}
		store:=repositories.NewSQLStore(db)
		business:=users_biz.NewSignUpUser(store)
		if err:=business.SignUpUser(c.Request.Context(),&data);err!=nil{
			log.Println(err)
			c.JSON(http.StatusBadRequest,gin.H{
				"err":err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK,common.NewDataResponse(http.StatusOK, data))
	}

}