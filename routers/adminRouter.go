package routers

import (
	"fake-market/controllers/admin"
	"fake-market/middlewares"

	"github.com/gin-gonic/gin"
)

func AdminRouterInit(r *gin.Engine) {
	adminRouters := r.Group("admin", middlewares.InitMiddleware)
	{
		adminRouters.GET("/login", admin.AdminController{}.Login)
		adminRouters.POST("/dologin", admin.AdminController{}.Dolog)
	}
}
