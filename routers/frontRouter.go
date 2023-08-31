package routers

import (
	"fake-market/controllers/defaults"
	"fake-market/middlewares"

	"github.com/gin-gonic/gin"
)

// 前台front/index/某某.html
func FrontRouterInit(r *gin.Engine) {
	adminRouters := r.Group("front", middlewares.InitMiddleware)
	{
		adminRouters.GET("/index", defaults.DefaultController{}.Index)

	}
}
