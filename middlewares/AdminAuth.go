package middlewares

import (
	"fake-market/models"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func InitAdminAuthMiddleware(c *gin.Context) {
	// 获取cookie
	cookie111, _ := c.Request.Cookie("admin_cookie")
	if cookie111 == nil {
		c.HTML(200, "admin/public/error.html", gin.H{
			"gotourl": "/admin/login",
			"message": "没有cookie,要登录",
		})
	} else {
		//判断是否用户名和密码是否存在已经在登录界面判断过了
		//获取 Url 路径去掉 Get 传值
		pathname := strings.Split(c.Request.URL.String(), "?")[0]
		fmt.Println(pathname)
		//判断redis中是否有cookie对应的session
		userinfo := models.CookieRedisStore{}.Get(cookie111.Value)
		superCheck := []models.Manager{}
		models.DB.Where("Username = ?", cookie111.Value).Find(&superCheck)
		if userinfo != nil && superCheck[0].IsSuper != 1 {

			//拼接sql查询语句，gorm自带的拼接 拼接 带引号的值会出错
			//sql := fmt.Sprintf("SELECT  `manager`.username, `access`.url FROM `manager` INNER JOIN `role_access` ON  `manager`.role_id=`role_access`.role_id AND `manager`.username = '%s' INNER JOIN `access` ON `role_access`.access_id=`access`.id", cookie111.Value)
			//执行原生函数获取name和可访问的url的对应值.Joins("Company").Joins("Manager").Joins("Account").First(&user, 1)

			nameurls := []models.NameUrl{}
			//finallyValue := "'" + cookie111.Value + "'"
			finallyValue := fmt.Sprintf("SELECT manager.username , access.url FROM manager left join role_access on manager.role_id = role_access.manager_id left join access on access.id = role_access.access_id "+"WHERE manager.username ="+"'%v'", cookie111.Value)
			models.DB.Raw(finallyValue).Scan(&nameurls)

			name_url_map := make(map[string]string)
			for _, t := range nameurls {
				// 把结构体添加到map中
				name_url_map[t.Url] = t.Username
			}
			// 查询 map 中是否含有 key 为 "foo" 的项
			_, ok := name_url_map[pathname]
			if !ok {
				//c.Redirect(http.StatusFound, c.Request.Referer())
				c.HTML(200, "admin/public/error.html", gin.H{
					"gotourl": c.Request.Referer(),
					"message": "没有权限",
				})
			}

		} else if userinfo != nil && superCheck[0].IsSuper == 1 {
			models.SugarLogger.Infof("超级管理员登录! 用户名 = %s for URL %s", cookie111.Value, strings.Split(c.Request.URL.String(), "?")[0])

		}

	}
}

/*      舍去下面的部分           */
/*舍去下面 （gprca这部分）*/
/*                舍去gprca这部分，用 models.DB.Exec 执行原生函数解决了问题                 */
// 查询有哪些角色对应manager
/* func QueryRolesByHeaders(c *gin.Context, header http.Header) (roles []string, err error) {

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
*/
