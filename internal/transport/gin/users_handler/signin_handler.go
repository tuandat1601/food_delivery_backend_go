package users_handler

import (
	"food_delivery/common"
	"food_delivery/config"
	"food_delivery/internal/biz/users_biz"
	"food_delivery/internal/model"
	"food_delivery/internal/repositories"
	"food_delivery/internal/util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SignIn godoc
// @Summary Sign in
// @Description Sign in a user and return an access token
// @Accept json
// @Produce json
// @Param user body model.User true "User credentials"
// @Success 200 {object} common.DataResponseLogin "Sign-in successful"
// @Failure 400 {object} map[string]string "Lỗi đăng ký người dùng"

// @Router /v1/users/signin [post]

func SignIn(db *gorm.DB) func(c *gin.Context) {

	return func(c *gin.Context) {

		var data model.User
		if err := c.ShouldBind(&data); err != nil {
			return
		}

		store := repositories.NewSQLStore(db)
		business := users_biz.NewSignInUser(store)
		if err := business.SignInUser(c.Request.Context(), &data); err != nil {

			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		}
		config, _ := config.LoadConfig(".")
		log.Println(data.Id,data.Email,data.UserName)
		access_token, err := util.CreateToken(config.AccessTokenExpiresIn, map[string]interface{}{
			"id":      data.Id,
			"email":  data.Email,
			"username":data.UserName,
		    }, config.AccessTokenPrivateKey)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
			return
		}
		refresh_token, err := util.CreateToken(config.RefreshTokenExpiresIn, map[string]interface{}{
			"id":      data.Id,
			"email":  data.Email,
			"username":data.UserName,
		    }, config.RefreshTokenPrivateKey)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
			return
		}
		c.SetCookie("access_token", access_token, config.AccessTokenMaxAge*60, "/", "localhost", false, true)
		c.SetCookie("refresh_token", refresh_token, config.RefreshTokenMaxAge*60, "/", "localhost", false, true)
		c.SetCookie("logged_in", "true", config.AccessTokenMaxAge*60, "/", "localhost", false, false)
		response := common.DataResponseLogin{
			Status:      200,
			AccessToken: access_token,
		    }
		c.JSON(http.StatusOK, response)
	}

}
