package users_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)
// @Summary Đăng xuất người dùng
// @Description Đăng xuất người dùng và hủy mã thông báo truy cập
// @Produce json
// @Tags User
// @Success 200 {object} map[string]string "logout success"
// @Router /v1/users/logout [get]
func LogOut(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ctx.SetCookie("access_token", "", -1, "/", "localhost", false, true)
		ctx.SetCookie("refresh_token", "", -1, "/", "localhost", false, true)
		ctx.SetCookie("logged_in", "", -1, "/", "localhost", false, true)
		ctx.JSON(http.StatusOK, gin.H{"status": "logout success"})

	}
}
