package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"wozaizhao.com/diancan/controllers/business"
	"wozaizhao.com/diancan/controllers/qiniu"
	"wozaizhao.com/diancan/controllers/user"
)

// SetupRouter 路由
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// allows all origins
	r.Use(cors.Default())

	// 小程序端接口 需要wxtoken
	wx := r.Group("/wx")
	{
		// 推送用户信息,后端保存
		wx.POST("userinfo", user.SaveUserInfo)
		// 当前用户
		wx.GET("currentUser", user.GetCurrentUser)
		// 获取我的红包
		wx.GET("userBonus", business.GetMyBonus)
		// 下单
		wx.POST("order", business.MakeOrder)
		// 支付订单
		wx.POST("pay", business.PayOrder)
		// 获取我的订单
		wx.GET("orders", business.GetMyOrders)
		// 获取我的订单
		wx.GET("order/:id", business.GetMyOrder)
		// 提交反馈
		wx.POST("feedback", business.SubmitFeedback)

	}

	// 一般接口，不需要token
	api := r.Group("api")
	{
		// 上传文件至七牛云
		api.POST("upload", qiniu.Upload)
		// wx登录，code换token
		api.POST("wxlogin", user.WxLogin)
		// 后台登录
		api.POST("login", user.Login)
		// 获取商品(分类下的)
		api.GET("goods", business.GetGoodsList)
		// 获取商品
		api.GET("goods/:id", business.GetGoods)
		// 获取Banner
		api.GET("banners", business.GetBanners)
		// 获取Banner
		api.GET("banner/:id", business.GetBanner)
		// 获取分类
		api.GET("categories", business.GetCategories)
		// 获取分类
		api.GET("category/:id", business.GetCategory)
		// 获取配置
		api.GET("configs", business.GetConfigs)
		// 获取配置
		api.GET("config/:id", business.GetConfig)

	}

	// 管理后台接口, 需要jwtToken
	admin := r.Group("/admin")
	{
		// 获取当前用户
		admin.GET("currentUser", user.GetCurrentAdmin)
		// 新增Banner
		admin.POST("banner", business.PostBanner)
		// 删除Banner
		admin.DELETE("banner/:id", business.DelBanner)
		// 修改Banner
		admin.PUT("banner/:id", business.ModifyBanner)
		// 获取红包
		admin.GET("bonuses", business.GetBonuses)
		// 获取红包
		admin.GET("bonus/:id", business.GetBonus)
		// 新增红包
		admin.POST("bonus", business.PostBonus)
		// 删除红包
		admin.DELETE("bonus/:id", business.DelBonus)
		// 修改红包
		admin.PUT("bonus/:id", business.ModifyBonus)
		// 新增分类
		admin.POST("category", business.PostCategory)
		// 删除分类
		admin.DELETE("category/:id", business.DelCategory)
		// 修改分类
		admin.PUT("category/:id", business.ModifyCategory)
		// 新增配置
		admin.POST("config", business.PostConfig)
		// 删除配置
		admin.DELETE("config/:id", business.DelConfig)
		// 修改配置
		admin.PUT("config/:id", business.ModifyConfig)
		// 获取反馈
		admin.GET("feedbacks", business.GetFeedbacks)
		// 获取反馈
		admin.GET("feedback/:id", business.GetFeedback)
		// 删除反馈
		admin.DELETE("feedback/:id", business.DelFeedback)
		// 修改反馈
		admin.PUT("feedback/:id", business.ModifyFeekback)
		// 新增商品
		admin.POST("goods", business.PostGoods)
		// 删除商品
		admin.DELETE("goods/:id", business.DelGoods)
		// 修改商品
		admin.PUT("goods/:id", business.ModifyGoods)
		// 获取订单
		admin.GET("orders", business.GetOrders)
		// 获取订单
		admin.GET("order/:id", business.GetOrder)
		// 删除订单
		admin.DELETE("order/:id", business.DelOrder)
		// 修改订单
		admin.PUT("order/:id", business.ModifyOrder)
	}

	return r
}
