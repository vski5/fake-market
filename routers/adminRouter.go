package routers

import (
	"fake-market/controllers/admin"
	"fake-market/middlewares"

	"github.com/gin-gonic/gin"
)

// 登录管理后台
func AdminRouterInit(r *gin.Engine) {
	adminRouters := r.Group("admin", middlewares.InitMiddleware)
	{
		adminRouters.GET("/login", admin.AdminController{}.Login)
		adminRouters.POST("/dologin", admin.AdminController{}.Dolog)
	}
}

// 管理后台的管理
func ManagerRouterInit(r *gin.Engine) {
	adminRouters := r.Group("admin/manager", middlewares.InitMiddleware)
	{
		adminRouters.GET("/index", admin.ManagerController{}.Index) //Manager首页admin/manager/index
		adminRouters.GET("/add", admin.ManagerController{}.Add)
		adminRouters.GET("/edit", admin.ManagerController{}.Edit)
		adminRouters.GET("/delete", admin.ManagerController{}.Delete)
	}
}

// 轮播图管理
func FocusRouterInit(r *gin.Engine) {
	adminRouters := r.Group("admin/focus", middlewares.InitMiddleware)
	{
		adminRouters.GET("/index", admin.FocusController{}.Index)
		adminRouters.GET("/add", admin.FocusController{}.Add)
		adminRouters.GET("/edit", admin.FocusController{}.Edit)
		adminRouters.GET("/delete", admin.FocusController{}.Delete)
	}
}
