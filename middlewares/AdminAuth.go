package middlewares

import (
	"fake-market/models"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/storyicon/grbac"
	"gorm.io/gorm"
)

func InitAdminAuthMiddleware(c *gin.Context) {
	// 获取cookie
	cookie111, _ := c.Request.Cookie("admin_cookie")
	if cookie111 == nil {
		c.Redirect(302, "/admin/login")
	} else {
		//获取 Url 路径去掉 Get 传值
		pathname := strings.Split(c.Request.URL.String(), "?")[0]
		//判断redis中是否有cookie对应的session
		userinfo := models.CookieRedisStore{}.Get(cookie111.Value)
		//类型断言，先类型断言判断是否为string，确定之后才能进行下一步
		//userinfoStr, ok := userinfo.(string)
		if userinfo != nil {
			name_url := []models.NameUrl{}
			//执行原生函数获取name和可访问的url的对应值
			models.DB.Exec("SELECT  `manager`.username, `access`.url FROM `manager` INNER JOIN `role_access` ON  `manager`.role_id=`role_access`.role_id AND `manager`.username = '?' INNER JOIN `access` ON `role_access`.access_id=`access`.id", cookie111.Name).Find(&name_url)
			name_url_map := make(map[string]string)
			for _, value1 := range name_url {
				name_url_map[value1.Username] = value1.Url

			}
			roles := models.MapString2Slice(name_url_map)
			models.InSliceOK(roles, pathname)

		} else {
			c.HTML(200, "admin/public/error.html", gin.H{
				"gotourl": "用户名或者密码异常",
				"message": "/admin/login",
			})
		}
	}
}

/*      舍去下面的部分           */
/*                舍去gprca这部分，用 models.DB.Exec 执行原生函数解决了问题                 */
// 查询有哪些角色对应manager
func QueryRolesByHeaders(c *gin.Context, header http.Header) (roles []string, err error) {

	//获取cookie
	cookie111, _ := c.Request.Cookie("admin_cookie")
	if cookie111 == nil {
		c.Redirect(302, "/admin/login")
	} else {
		//获取 Url 路径去掉 Get 传值
		//pathname := strings.Split(c.Request.URL.String(), "?")[0]
		//判断redis中是否有cookie对应的session
		userinfo := models.CookieRedisStore{}.Get(cookie111.Value)
		//类型断言，先类型断言判断是否为string，确定之后才能进行下一步
		//userinfoStr, ok := userinfo.(string)
		if userinfo != nil {
			name_url := []models.NameUrl{}
			//执行原生函数获取name和可访问的url的对应值
			models.DB.Exec("SELECT  `manager`.username, `access`.url FROM `manager` INNER JOIN `role_access` ON  `manager`.role_id=`role_access`.role_id AND `manager`.username = '?' INNER JOIN `access` ON `role_access`.access_id=`access`.id", cookie111.Name).Find(&name_url)
			name_url_map := make(map[string]string)
			for _, value1 := range name_url {
				name_url_map[value1.Username] = value1.Url

			}
			roles := models.MapString2Slice(name_url_map)
			return roles, err

		} else {
			c.HTML(200, "admin/public/error.html", gin.H{
				"gotourl": "用户名或者密码异常",
				"message": "/admin/login",
			})
		}

	}
	return roles, err
}

type MySQLLoader struct {
	session *gorm.DB
}

// 数据库的连接在models.DB，返回一个包裹了连接的结构体用于制造method
func NewMySQLLoader() *MySQLLoader {
	loader := &MySQLLoader{}
	loader.session = models.DB
	return loader
}

func (loader *MySQLLoader) LoadRules() (rules grbac.Rules, err error) {
	// 在这里实现你的逻辑
	// ...
	// 你可以从数据库或文件加载授权规则
	// 但是你需要以 grbac.Rules 的格式返回你的身份验证规则
	// 提示：你还可以将此函数绑定到golang结构体

	return
}
func Authorization() gin.HandlerFunc {
	// 在这里，通过“grbac.WithLoader”接口使用自定义Loader功能
	// 并指定应每30分钟调用一次LoadAuthorizationRules函数以获取最新的身份验证规则。

	rbac, err := grbac.New(grbac.WithLoader(LoadAuthorizationRules, time.Minute*30))
	if err != nil {
		panic(err)
	}
	return func(c *gin.Context) {
		roles, err := QueryRolesByHeaders(c.Request.Header)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		state, _ := rbac.IsRequestGranted(c.Request, roles)
		if !state.IsGranted() {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}
