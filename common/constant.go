package common

// 订单状态
const (
	OrderConfirmed = 1 // 已提交
	OrderPaid      = 2 // 已支付
)

// 红包发放的方式
const (
	SendByUser  = 1 // 按用户发放
	SendByGoods = 2 // 按商品发放
	SendByOrder = 3 // 按订单发放
	SendByPrint = 4 // 线下发放
)

// 反馈 留言类型
const (
	MsgMessage   = 1 // 留言
	MsgComplaint = 2 // 投诉
	MsgEnquiry   = 3 // 询问
	MsgCustomer  = 4 // 售后
	MsgComment   = 5 // 评论
)

// 反馈 消息状态
const (
	MsgUnread = 1 // 未读
	MsgRead   = 2 // 已读
)

// 七牛云图片类型
const (
	ImgBanner = 1 // Banner图片
	ImgGoods  = 2 // 商品图片
)

// 用户角色
const (
	User  = "user"  // 普通用户
	Admin = "admin" // 管理员
)
