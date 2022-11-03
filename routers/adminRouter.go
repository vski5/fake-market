package routers

import (
	"fake-market/controllers/admin"

	"github.com/gin-gonic/gin"
)

func AdminRouterInit(r *gin.Engine) {
	adminRouters := r.Group("admin")
	{
		adminRouters.GET("/login", admin.AdminController{}.Login)
		adminRouters.GET("/test2", admin.AdminController{}.Back)
	}
}
