package middlewares

import (
	"encoding/json"
	"fake-market/models"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func InitAdminAuthMiddleware(c *gin.Context /* , username string */) {

	//判断用户【管理员】是否登录

	/* admincUsername := c.PostForm("username") */
	cookie111, _ := c.Request.Cookie("admin_cookie")
	if cookie111 == nil {
		c.Redirect(302, "/admin/login")
	} else {
		//获取 Url 路径去掉 Get 传值
		pathname := strings.Split(c.Request.URL.String(), "?")[0]

		userinfo := models.CookieRedisStore{}.Get(cookie111.Value)

		/* userinfo := models.CookieRedisStore{}.Get("admin") */
		//先类型断言判断是否为string，确定之后才能进行下一步
		//类型断言
		userinfoStr, ok := userinfo.(string)
		if ok {
			if len(userinfoStr) > 0 {
				var u []models.Manager
				err333 := json.Unmarshal([]byte(userinfoStr), &u)

				//如果验证成功
				if len(u) > 0 && u[0].Username != "" && err333 != nil {
					/*
					   u总为[]
					   过一段时间修补，先用着
					*/
					if pathname != "/admin/login" && pathname != "/admin/doLogin" && pathname != "/admin/captcha" {
						c.Redirect(302, "/admin/login")
					} else { //用户登录成功 权限判断
						urlPath := strings.Replace(pathname, "/admin/", "", 1)

						if u[0].IsSuper == 0 {

							// 1、根据 当前角色名 获取 权限id
							managerName := []models.Manager{}
							models.DB.Where("Username=?", userinfoStr).Find(&managerName)
							managerRoleId := managerName[0].RoleId
							//用RoleId获取accessId
							rolrAccessIdList := []models.RoleAccess{}
							models.DB.Where("RoleId=?", managerRoleId).Find(&rolrAccessIdList)
							//历遍rolrAccessIdList里的AccessId
							// 二维数组的遍历
							rolrAccessIdMap := make(map[int]int)
							for _, value1 := range rolrAccessIdList {
								rolrAccessIdMap[value1.AccessId] = value1.AccessId
							}

							//用accessId获取Access表里的url
							accessUrlList := []models.Access{}
							models.DB.Where("Url=?", models.Map2Slice(rolrAccessIdMap)).Find(&accessUrlList)
							// 二维数组的遍历
							rolrAccessUrlMap := make(map[string]string)
							for _, value1 := range accessUrlList {
								rolrAccessUrlMap[value1.Url] = value1.Url
							}

							//判断访问的url对应的权限
							if models.InSliceOK(models.MapString2Slice(rolrAccessUrlMap), urlPath) {
								c.String(200, "没有权限")
								c.Abort()
							}
						}

						fmt.Println("成功跳过验证登录的中间件")
					}
				} else { //如果验证失败
					if pathname != "/admin/login" && pathname != "/admin/doLogin" && pathname != "/admin/captcha" {
						/* c.Redirect(302, "/admin/login") */
					}
				}

			} else {
				c.HTML(200, "admin/public/error.html", gin.H{
					"gotourl": "用户名或者密码异常",
					"message": "/admin/login",
				})
			}
		}
	}
}
