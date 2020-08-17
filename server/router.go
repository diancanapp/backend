package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"wozaizhao.com/diancan/controllers/auxiliary"
	"wozaizhao.com/diancan/controllers/user"
)

// SetupRouter 路由
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// allows all origins
	r.Use(cors.Default())

	// 小程序端接口
	wx := r.Group("/wx")
	{
		// wx登录，code换token
		wx.POST("login", user.WxLogin)
		// 推送用户信息,后端保存
		wx.POST("userinfo", user.SaveUserInfo)
		// 检测token是否过期
		// wx.GET("checktoken", user.CheckToken)
		// 绑定帐号
		// wx.POST("bindcustomer", business.BindCustomer)
		// 更新绑定信息
		// wx.POST("updatebind", business.UpdateCustomer)
		// 当前用户
		// wx.GET("currentuser", business.GetCurrentCustomer)
		// 下单
		// wx.POST("order", business.Order)
		// 获取我的订单
		// wx.GET("orders", business.GetMyOrders)
		// 订单详情
		// wx.GET("order/:id", business.GetOrderDetail)
		// 客户评价
		// wx.POST("comment", business.CommetOrder)

	}

	// 一般接口
	api := r.Group("api")
	{
		api.POST("upload", auxiliary.Upload)
	}

	//管理后台接口
	// admin := r.Group("/admin")
	// {
	// 	admin
	// }

	return r
}
