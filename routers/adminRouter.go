package routers

import (
	"fake-market/controllers/admin"
	"fake-market/middlewares"

	"github.com/gin-gonic/gin"
)

// 登录管理后台/admin/login
func AdminRouterInit(r *gin.Engine) {
	adminRouters := r.Group("admin", middlewares.InitMiddleware)
	{
		adminRouters.GET("/login", admin.AdminController{}.Login)
		adminRouters.GET("/captcha", admin.LoginCaptchaController{}.DoCaptchaMake) //生成 验证码
		/* adminRouters.GET("/verify", admin.LoginCaptchaController{}.DoCaptcha)      //检验验证码 */
		adminRouters.POST("/dologin", admin.AdminController{}.Dolog)
	}
}

// 管理后台的管理员的管理
func ManagerRouterInit(r *gin.Engine) {
	adminRouters := r.Group("admin/manager", middlewares.InitMiddleware, middlewares.InitAdminAuthMiddleware)
	{
		adminRouters.GET("/index", admin.ManagerController{}.Index) //Manager首页admin/manager/index
		adminRouters.GET("/add", admin.ManagerController{}.Add)
		adminRouters.GET("/edit", admin.ManagerController{}.Edit)
		adminRouters.GET("/delete", admin.ManagerController{}.Delete)
	}
}

// 轮播图管理
func FocusRouterInit(r *gin.Engine) {
	adminRouters := r.Group("admin/focus", middlewares.InitMiddleware, middlewares.InitAdminAuthMiddleware)
	{
		adminRouters.GET("/index", admin.FocusController{}.Index)
		adminRouters.GET("/add", admin.FocusController{}.Add)
		adminRouters.GET("/edit", admin.FocusController{}.Edit)
		adminRouters.GET("/delete", admin.FocusController{}.Delete)
	}
}

// 管理员权限管理
func RoleRouterInit(r *gin.Engine) {
	adminRouters := r.Group("admin/role", middlewares.InitMiddleware, middlewares.InitAdminAuthMiddleware)
	{
		adminRouters.GET("/index", admin.RoleController{}.Index)
		adminRouters.GET("/add", admin.RoleController{}.Add)
		adminRouters.POST("/doAdd", admin.RoleController{}.DoAdd)
		adminRouters.GET("/edit", admin.RoleController{}.Edit)
		adminRouters.POST("/doEdit", admin.RoleController{}.DoEdit)
		adminRouters.GET("/delete", admin.RoleController{}.Delete)
	}
}
