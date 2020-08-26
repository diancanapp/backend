package models

import (
	"github.com/jinzhu/gorm"
	"wozaizhao.com/diancan/common"
)

// Order 订单
type Order struct {
	gorm.Model
	UserID      uint   `json:"userID" gorm:"type:mediumint(8);NOT NULL;DEFAULT '0'"` // 用户ID
	OrderStatus uint   `json:"orderStatus" gorm:"type:tinyint(1);DEFAULT '0'"`       // 订单状态
	PayID       uint   `json:"payID" gorm:"type:mediumint(8);DEFAULT '0'"`           // 支付ID
	OrderAmount uint   `json:"orderAmount" gorm:"type:int(11);NOT NULL;DEFAULT '0'"` // 订单金额
	Bonus       uint   `json:"bonus" gorm:"type:int(11);DEFAULT '0'"`                // 红包
	BonusID     uint   `json:"bonusID" gorm:"type:mediumint(8);DEFAULT '0'"`         // 红包ID
	MoneyPaid   uint   `json:"moneyPaid" gorm:"type:int(11);DEFAULT '0'"`            // 支付金额
	OrderNote   string `json:"orderNote" gorm:"type:varchar(255);DEFAULT ''"`        // 备注
	OrderType   uint   `json:"orderType" gorm:"type:tinyint(1);DEFAULT '0'"`         // 订单类型
}

// -------------------------------新建-------------------------------

// CreateOrder 创建订单
func CreateOrder(userID, orderAmount, bonus, bonusID uint, orderNote string) error {
	var order = Order{UserID: userID, OrderStatus: common.OrderConfirmed, OrderAmount: orderAmount, Bonus: bonus, BonusID: bonusID, OrderNote: orderNote}
	err := DB.Create(&order).Error
	return err
}

// -------------------------------删除-------------------------------

// -------------------------------更改-------------------------------

// 支付订单

// -------------------------------查询单个-------------------------------

// QueryOrderByID 根据id查order
func QueryOrderByID(id uint) (order Order, err error) {
	err = DB.Where("id = ?", id).Find(&order).Error
	return order, err
}

// -------------------------------查询所有-------------------------------

// QueryOrders 所有Order
func QueryOrders(size, offset uint) (orders []Order, err error) {
	err = DB.Limit(size).Offset(offset).Find(&orders).Error
	return orders, err
}
