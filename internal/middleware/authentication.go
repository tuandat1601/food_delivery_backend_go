package middleware

import (
	"food_delivery/config"
	"food_delivery/internal/biz/users_biz"
	"food_delivery/internal/repositories"
	"food_delivery/internal/util"
	"log"

	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

func DeserializeUser(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var access_token string
		cookie, err := ctx.Cookie("access_token")

		authorizationHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			access_token = fields[1]
		} else if err == nil {
			access_token = cookie
		}

		if access_token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You are not logged in"})
			return
		}

		config, _ := config.LoadConfig(".")
		sub, err := util.ValidateToken(access_token, config.AccessTokenPublicKey)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
			return
		}
		subs:=sub.(map[string]interface{})
		store:=repositories.NewSQLStore(db)
		business:=users_biz.NewGetUser(store)
		iduser :=subs["id"].(float64)
		user ,err:=business.GetUserById(ctx.Request.Context(),int(iduser));
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "The user belonging to this token no logger exists","ccc":sub})
			return
		}
		log.Println(user)
		// ctx.Set("currentUser", user)
		ctx.Next()
	}
}