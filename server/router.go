package server

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"wozaizhao.com/diancan/controllers/business"
	"wozaizhao.com/diancan/controllers/user"
)

//TokenAuthMiddleware ...
//JWT Authentication middleware attached to each request that needs to be authenitcated to validate the access_token in the header
func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !user.TokenValid(c) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()
	}
}

// SetupRouter 路由
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// allows all origins
	r.Use(cors.Default())

	// 小程序端接口 需要wxtoken
	wx := r.Group("/wx")
	{
		// 上传文件至七牛云
		wx.POST("upload", business.UploadByWeapp)
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
		api.POST("upload", business.Upload)
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
		// 上传文件至七牛云
		admin.POST("upload", TokenAuthMiddleware(), business.UploadByAdmin)
		// 获取当前用户
		admin.GET("currentUser", TokenAuthMiddleware(), user.GetCurrentAdmin)
		// 新增Banner
		admin.POST("banner", TokenAuthMiddleware(), business.PostBanner)
		// 删除Banner
		admin.DELETE("banner/:id", TokenAuthMiddleware(), business.DelBanner)
		// 修改Banner
		admin.PUT("banner/:id", TokenAuthMiddleware(), business.ModifyBanner)
		// 获取红包
		admin.GET("bonuses", TokenAuthMiddleware(), business.GetBonuses)
		// 获取红包
		admin.GET("bonus/:id", TokenAuthMiddleware(), business.GetBonus)
		// 新增红包
		admin.POST("bonus", TokenAuthMiddleware(), business.PostBonus)
		// 删除红包
		admin.DELETE("bonus/:id", TokenAuthMiddleware(), business.DelBonus)
		// 修改红包
		admin.PUT("bonus/:id", TokenAuthMiddleware(), business.ModifyBonus)
		// 新增分类
		admin.POST("category", TokenAuthMiddleware(), business.PostCategory)
		// 删除分类
		admin.DELETE("category/:id", TokenAuthMiddleware(), business.DelCategory)
		// 修改分类
		admin.PUT("category/:id", TokenAuthMiddleware(), business.ModifyCategory)
		// 新增配置
		admin.POST("config", TokenAuthMiddleware(), business.PostConfig)
		// 删除配置
		admin.DELETE("config/:id", TokenAuthMiddleware(), business.DelConfig)
		// 修改配置
		admin.PUT("config/:id", TokenAuthMiddleware(), business.ModifyConfig)
		// 获取反馈
		admin.GET("feedbacks", TokenAuthMiddleware(), business.GetFeedbacks)
		// 获取反馈
		admin.GET("feedback/:id", TokenAuthMiddleware(), business.GetFeedback)
		// 新增商品
		admin.POST("goods", TokenAuthMiddleware(), business.PostGoods)
		// 删除商品
		admin.DELETE("goods/:id", TokenAuthMiddleware(), business.DelGoods)
		// 修改商品
		admin.PUT("goods/:id", TokenAuthMiddleware(), business.ModifyGoods)
		// 获取订单
		admin.GET("orders", TokenAuthMiddleware(), business.GetOrders)
		// 获取订单
		admin.GET("order/:id", TokenAuthMiddleware(), business.GetOrder)
	}

	return r
}
