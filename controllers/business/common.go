package business

import (
	"github.com/gin-gonic/gin"
	"wozaizhao.com/diancan/controllers/base"
	"wozaizhao.com/diancan/controllers/qiniu"
	"wozaizhao.com/diancan/models"
)

// Upload 上传
func Upload(c *gin.Context) {
	qiniu.Upload(c)
}

// GetGoodsList
func GetGoodsList(c *gin.Context) {
}

// GetGoods
func GetGoods(c *gin.Context) {
}

// GetBanners 获取banners
func GetBanners(c *gin.Context) {
	banners, err := models.QueryBanners()
	if err != nil {
		base.Fail(c, err)
		return
	}

	base.Ok(c, banners)
}

// GetBanner
func GetBanner(c *gin.Context) {
}

// GetCategory
func GetCategory(c *gin.Context) {
}

// GetConfigs
func GetConfigs(c *gin.Context) {
}

// GetConfig
func GetConfig(c *gin.Context) {
}
