package admin

import (
	"fake-market/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GoodsSettingController struct {
	BaseController
}

func (con GoodsSettingController) Index(c *gin.Context) {
	setting := models.Setting{}
	models.DB.First(&setting)
	c.HTML(http.StatusOK, "admin/setting/index.html", gin.H{
		"setting": setting,
	})
}

func (con GoodsSettingController) DoEdit(c *gin.Context) {
	setting := models.Setting{Id: 1}
	models.DB.Find(&setting)
	if err := c.ShouldBind(&setting); err != nil {
		fmt.Println(err)
		con.Error(c, "修改数据失败,请重试", "/admin/setting/index")
		return
	} else {
		// 上传图片 logo
		//goodsImg, _ := models.UploadImg(c, "goods_img")
		siteLogo, err1 := models.UploadImg(c, "site_logo")
		if len(siteLogo) > 0 && err1 == nil {
			setting.SiteLogo = siteLogo
		}
		//上传图片 no_picture
		noPicture, err2 := models.UploadImg(c, "no_picture")
		if len(noPicture) > 0 && err2 == nil {
			setting.NoPicture = noPicture
		}

		err3 := models.DB.Save(&setting).Error
		if err3 != nil {
			con.Error(c, "修改数据失败", "/admin/setting/index")
			return
		}

		con.Success(c, "修改数据成功", "/admin/setting/index")
	}
}
