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
		adminRouters.GET("/changeStatus", admin.BaseController{}.ChangeStatus)
	}
}

// 管理后台的管理员的管理
func ManagerRouterInit(r *gin.Engine) {
	adminRouters := r.Group("admin/manager", middlewares.InitMiddleware, middlewares.InitAdminAuthMiddleware)
	{
		adminRouters.GET("/index", admin.ManagerController{}.Index) //Manager首页admin/manager/index
		adminRouters.GET("/add", admin.ManagerController{}.Add)
		adminRouters.POST("/doAdd", admin.ManagerController{}.DoAdd)
		adminRouters.GET("/edit", admin.ManagerController{}.Edit)
		adminRouters.POST("/doEdit", admin.ManagerController{}.DoEdit)
		adminRouters.POST("/delete", admin.ManagerController{}.Delete)
	}
}

// 管理员角色(组)管理
func RoleRouterInit(r *gin.Engine) {
	adminRouters := r.Group("admin/role", middlewares.InitMiddleware, middlewares.InitAdminAuthMiddleware)
	{
		adminRouters.GET("/index", admin.RoleController{}.Index)
		adminRouters.GET("/add", admin.RoleController{}.Add)
		adminRouters.POST("/doAdd", admin.RoleController{}.DoAdd)
		adminRouters.GET("/edit", admin.RoleController{}.Edit)
		adminRouters.POST("/doEdit", admin.RoleController{}.DoEdit)
		adminRouters.GET("/delete", admin.RoleController{}.Delete)
		adminRouters.GET("/auth", admin.RoleController{}.Auth)
		adminRouters.POST("/doAuth", admin.RoleController{}.DoAuth)
	}
}

// 管理员权限管理 的 展示
func AccessRouterInit(r *gin.Engine) {
	adminRouters := r.Group("admin/access", middlewares.InitMiddleware, middlewares.InitAdminAuthMiddleware)
	{
		adminRouters.GET("/index", admin.AccessController{}.Index)
		adminRouters.GET("/add", admin.AccessController{}.Add)
		adminRouters.POST("/doAdd", admin.AccessController{}.DoAdd)
		adminRouters.GET("/edit", admin.AccessController{}.Edit)
		adminRouters.POST("/doEdit", admin.AccessController{}.DoEdit)

	}
}

// 轮播图 管理
func FocusRouterInit(r *gin.Engine) {
	adminRouters := r.Group("admin/focus", middlewares.InitMiddleware, middlewares.InitAdminAuthMiddleware)
	{
		adminRouters.GET("/index", admin.FocusController{}.Index)
		adminRouters.GET("/add", admin.FocusController{}.Add)
		adminRouters.POST("/doAdd", admin.FocusController{}.DoAdd)
		adminRouters.GET("/edit", admin.FocusController{}.Edit)
		adminRouters.POST("/doEdit", admin.FocusController{}.DoEdit)
		//删除轮播图
		adminRouters.GET("/delete", admin.FocusController{}.Delete)

	}
}

// 商品图 管理GoodsController
func GoodsRouterInit(r *gin.Engine) {
	adminRouters := r.Group("admin/goods", middlewares.InitMiddleware, middlewares.InitAdminAuthMiddleware)
	{
		adminRouters.GET("/index", admin.GoodsController{}.Index)
		adminRouters.GET("/add", admin.GoodsController{}.Add)
		adminRouters.POST("/doAdd", admin.GoodsController{}.DoAdd)
		adminRouters.GET("/edit", admin.GoodsController{}.Edit)
		adminRouters.POST("/doEdit", admin.GoodsController{}.DoEdit)
		//删除商品图
		adminRouters.GET("/delete", admin.GoodsController{}.Delete)
		//获取商品类型属性/admin/goods/goodsTypeAttribute
		adminRouters.GET("/goodsTypeAttribute", admin.GoodsController{}.GoodsTypeAttribute)

	}
}

// 商品图类型的 管理
func GoodsTypeRouterInit(r *gin.Engine) {
	adminRouters := r.Group("admin/goodsType", middlewares.InitMiddleware, middlewares.InitAdminAuthMiddleware)
	{
		adminRouters.GET("/index", admin.GoodsTypeController{}.Index)
		adminRouters.GET("/add", admin.GoodsTypeController{}.Add)
		adminRouters.POST("/doAdd", admin.GoodsTypeController{}.DoAdd)
		adminRouters.GET("/edit", admin.GoodsTypeController{}.Edit)
		adminRouters.POST("/doEdit", admin.GoodsTypeController{}.DoEdit)
		//删除商品类型
		adminRouters.GET("/delete", admin.GoodsTypeController{}.Delete)

	}
}

// 商品有哪些种类的类型 管理 /admin/goodsTypeAttribute GoodsTypeAttributeController
func GoodsTypeAttributeRouterInit(r *gin.Engine) {
	adminRouters := r.Group("admin/goodsTypeAttribute", middlewares.InitMiddleware, middlewares.InitAdminAuthMiddleware)
	{
		adminRouters.GET("/index", admin.GoodsTypeAttributeController{}.Index)
		adminRouters.GET("/add", admin.GoodsTypeAttributeController{}.Add)
		adminRouters.POST("/doAdd", admin.GoodsTypeAttributeController{}.DoAdd)
		adminRouters.GET("/edit", admin.GoodsTypeAttributeController{}.Edit)
		adminRouters.POST("/doEdit", admin.GoodsTypeAttributeController{}.DoEdit)
		//删除商品类型种类
		adminRouters.GET("/delete", admin.GoodsTypeAttributeController{}.Delete)

	}
}

// 商品信息的管理
func GoodsInfoRouterInit(r *gin.Engine) {
	adminRouters := r.Group("admin/goodsinfo", middlewares.InitMiddleware, middlewares.InitAdminAuthMiddleware)
	{
		adminRouters.GET("/index", admin.GoodsInfoController{}.Index)
		adminRouters.GET("/add", admin.GoodsInfoController{}.Add)
		adminRouters.POST("/doAdd", admin.GoodsInfoController{}.DoAdd) //商品信息添加页面 提交的部分 的控制器/路由
		adminRouters.GET("/edit", admin.GoodsInfoController{}.Edit)
		adminRouters.POST("/doEdit", admin.GoodsInfoController{}.DoEdit)
		//删除商品信息
		adminRouters.GET("/delete", admin.GoodsInfoController{}.Delete)
		//上传图片的路由/admin/goodsinfo/static/goodsUpload/
		adminRouters.POST("/static/goodsUpload/", admin.GoodsInfoController{}.ImageUpload)
		//上传商品相册图片/admin/goods/imageUpload
		adminRouters.POST("/imageUpload", admin.GoodsInfoController{}.ImageUpload)
		adminRouters.GET("/changeGoodsImageColor", admin.GoodsInfoController{}.ChangeGoodsImageColor)
		adminRouters.GET("/removeGoodsImage", admin.GoodsInfoController{}.RemoveGoodsImage)

	}
}

// /admin/setting/index
func GoodsSettingRouterInit(r *gin.Engine) {
	adminRouters := r.Group("admin/setting", middlewares.InitMiddleware, middlewares.InitAdminAuthMiddleware)
	{
		adminRouters.GET("/index", admin.GoodsSettingController{}.Index)
		///admin/setting/doEdit
		adminRouters.POST("/doEdit", admin.GoodsSettingController{}.DoEdit)

	}
}

func GoodsNavRouterInit(r *gin.Engine) {
	adminRouters := r.Group("admin/nav", middlewares.InitMiddleware, middlewares.InitAdminAuthMiddleware)
	{
		adminRouters.GET("/", admin.NavController{}.Index)
		adminRouters.GET("/add", admin.NavController{}.Add)
		adminRouters.POST("/doAdd", admin.NavController{}.DoAdd)
		adminRouters.GET("/edit", admin.NavController{}.Edit)
		adminRouters.POST("/doEdit", admin.NavController{}.DoEdit)
		adminRouters.GET("/delete", admin.NavController{}.Delete)

	}
}

// 商品信息的管理（静态资源部分）
/* func GoodsInfoUploadRouterInit(r *gin.Engine) {
	adminRouters := r.Group("static/goodsUpload", middlewares.InitMiddleware, middlewares.InitAdminAuthMiddleware)
	{
		//上传图片的路由
		adminRouters.POST("/imageUpload", admin.GoodsInfoController{}.ImageUpload)

	}
} */
