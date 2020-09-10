package business

import (
	"github.com/gin-gonic/gin"
	"wozaizhao.com/diancan/common"
	"wozaizhao.com/diancan/controllers/base"
	"wozaizhao.com/diancan/controllers/qiniu"
	"wozaizhao.com/diancan/models"
)

// UploadByAdmin 从管理后台上传文件
func UploadByAdmin(c *gin.Context) {
	username := c.MustGet("username").(string)
	if username != "" {
		qiniu.Upload(c)
	}
}

// PostBanner 新增banner
func PostBanner(c *gin.Context) {
	var banner common.AddBanner
	if err := c.Bind(&banner); err != nil {
		base.FailUnprocessableEntity(c, err)
		return
	}
	err := models.CreateBanner(banner)

	if err != nil {
		base.Fail(c, "add banner fail!")
		return
	}

	banners, errQueryBanners := models.QueryBanners()
	if errQueryBanners != nil {
		base.Fail(c, err)
		return
	}
	base.Ok(c, banners)
}

// DelBanner 删除banner
func DelBanner(c *gin.Context) {
	var banner common.DelBanner
	if err := c.Bind(&banner); err != nil {
		base.FailUnprocessableEntity(c, err)
		return
	}
	err := models.DeleteBanner(banner.ID)

	if err != nil {
		base.Fail(c, "delete banner fail!")
		return
	}

	banners, errQueryBanners := models.QueryBanners()
	if errQueryBanners != nil {
		base.Fail(c, err)
		return
	}
	base.Ok(c, banners)
}

// ModifyBanner 修改banner
func ModifyBanner(c *gin.Context) {
	var banner common.ModifyBanner
	if err := c.Bind(&banner); err != nil {
		base.FailUnprocessableEntity(c, err)
		return
	}
	err := models.UpdateBanner(banner)

	if err != nil {
		base.Fail(c, "modify banner fail!")
		return
	}

	banners, errQueryBanners := models.QueryBanners()
	if errQueryBanners != nil {
		base.Fail(c, err)
		return
	}
	base.Ok(c, banners)
}

// GetBonuses
func GetBonuses(c *gin.Context) {
}

// GetBonus
func GetBonus(c *gin.Context) {
}

// PostBonus
func PostBonus(c *gin.Context) {
}

// DelBonus
func DelBonus(c *gin.Context) {
}

// ModifyBonus
func ModifyBonus(c *gin.Context) {
}

// GetCategories 获取商品分类
func GetCategories(c *gin.Context) {
	categories, err := models.QueryCategorys()
	if err != nil {
		base.Fail(c, err)
		return
	}

	base.Ok(c, categories)

}

// PostCategory 增加商品分类
func PostCategory(c *gin.Context) {
	var category common.AddCategory
	if err := c.Bind(&category); err != nil {
		base.FailUnprocessableEntity(c, err)
		return
	}
	err := models.CreateCategory(category)

	if err != nil {
		base.Fail(c, "add category fail!")
		return
	}

	categories, errQueryCategorys := models.QueryCategorys()
	if errQueryCategorys != nil {
		base.Fail(c, err)
		return
	}
	base.Ok(c, categories)
}

// DelCategory 删除商品分类
func DelCategory(c *gin.Context) {

	var category common.DelCategory
	if err := c.Bind(&category); err != nil {
		base.FailUnprocessableEntity(c, err)
		return
	}
	err := models.DeleteCategory(category.ID)

	if err != nil {
		base.Fail(c, "delete category fail!")
		return
	}

	categories, errQueryCategorys := models.QueryCategorys()
	if errQueryCategorys != nil {
		base.Fail(c, err)
		return
	}
	base.Ok(c, categories)
}

// ModifyCategory 修改商品分类
func ModifyCategory(c *gin.Context) {
	var category common.ModifyCategory
	if err := c.Bind(&category); err != nil {
		base.FailUnprocessableEntity(c, err)
		return
	}
	err := models.UpdateCategory(category)

	if err != nil {
		base.Fail(c, "modify category fail!")
		return
	}

	categories, errQueryCategorys := models.QueryCategorys()
	if errQueryCategorys != nil {
		base.Fail(c, err)
		return
	}
	base.Ok(c, categories)
}

// PostConfig
func PostConfig(c *gin.Context) {
}

// DelConfig
func DelConfig(c *gin.Context) {
}

// ModifyConfig
func ModifyConfig(c *gin.Context) {
}

// GetFeedbacks
func GetFeedbacks(c *gin.Context) {
}

// GetFeedback
func GetFeedback(c *gin.Context) {
}

// PostGoods
func PostGoods(c *gin.Context) {
}

// DelGoods
func DelGoods(c *gin.Context) {
}

// ModifyGoods
func ModifyGoods(c *gin.Context) {
}

// GetOrders
func GetOrders(c *gin.Context) {
}

// GetOrder
func GetOrder(c *gin.Context) {
}
