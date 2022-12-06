package api

import (
	"fake-market/models"

	"github.com/hunterhug/go_image"
)

/*
按宽度进行比例缩放，输入和输出都是图片字节数组:
func ScaleB2B(InRaw []byte, width int) (OutRaw []byte, err error)
按宽度进行比例缩放，输入输出都是文件:
func ScaleF2F(filename string, savepath string, width int) (err error)
按宽度和高度进行比例缩放，输入和输出都是图片字节数组:
func ThumbnailB2B(InRaw []byte, width int, height int) (OutRaw []byte, err error)
按宽度和高度进行比例缩放，输入和输出都是文件:
func ThumbnailF2F(filename string, savepath string, width int, height int) (err error)
检测图像文件真正文件类型,并返回真实文件名,参数为图像文件位置
func RealImageName(filename string) (filerealname string, err error)
文件改名,如果force为假,且新的文件名已经存在,那么抛出错误
func ChangeImageName(oldname string, newname string, force bool) (err error)
*/
type EditImage struct {
	filename string //图像位置
	savepath string //保存位置
	width    int    //宽度
	height   int    //高度
}

// 将某一图片文件进行缩放后存入另外的文件中,记得先实例化
func (image EditImage) TestImage() {

	//按照宽度和高度进行等比例缩放
	err := go_image.ThumbnailF2F(image.filename, image.savepath, image.width, image.height)
	if err != nil {
		models.SugarLogger.Errorf("生成按宽度高度缩放图 : Error = %s", err.Error())
	}

}
