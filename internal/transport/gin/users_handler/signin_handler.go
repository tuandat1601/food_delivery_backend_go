package users_handler

import (
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
		c.JSON(http.StatusOK, gin.H{"status": "success", "access_token": access_token})
	}

}
