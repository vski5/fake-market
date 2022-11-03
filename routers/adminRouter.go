package routers

import (
	"fake-market/controllers/admin"

	"github.com/gin-gonic/gin"
)

func AdminRouterInit(r *gin.Engine) {
	adminRouters := r.Group("admin")
	{
		adminRouters.GET("/login", admin.AdminController{}.Login)
		adminRouters.POST("/dologin", admin.AdminController{}.Dolog)
	}
}
