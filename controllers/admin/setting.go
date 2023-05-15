package admin
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GoodsSettingController struct {
	BaseController
}

func (con GoodsSettingController) Index(c *gin.Context){
	c.HTML(http.StatusOK, "admin/setting/index.html", gin.H{})
}