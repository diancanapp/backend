package models

import (
	"github.com/jinzhu/gorm"
	"wozaizhao.com/diancan/common"
)

// Order 订单
type Order struct {
	gorm.Model
	UserID      uint   `json:"user_id" gorm:"type:mediumint(8);NOT NULL;DEFAULT '0'"` // 用户ID
	OrderStatus uint   `json:"order_status" gorm:"type:tinyint(1);DEFAULT '0'"`       // 订单状态
	PayID       uint   `json:"pay_id" gorm:"type:mediumint(8);DEFAULT '0'"`           // 支付ID
	OrderAmount uint   `json:"order_amount" gorm:"type:int(11);NOT NULL;DEFAULT '0'"` // 订单金额
	Bonus       uint   `json:"bonus" gorm:"type:int(11);DEFAULT '0'"`                 // 红包
	BonusID     uint   `json:"bonus_id" gorm:"type:mediumint(8);DEFAULT '0'"`         // 红包ID
	MoneyPaid   uint   `json:"money_paid" gorm:"type:int(11);DEFAULT '0'"`            // 支付金额
	OrderNote   string `json:"pay_note" gorm:"type:varchar(255);DEFAULT ''"`          // 备注
	OrderType   uint   `json:"order_type" gorm:"type:tinyint(1);DEFAULT '0'"`         // 订单类型
}

// -------------------------------新建-------------------------------

// CreateOrder 创建订单
func CreateOrder(userID uint, orderAmount uint, bonus uint, bonusID uint, orderNote string) error {
	var order = Order{UserID: userID, OrderStatus: common.OrderConfirmed, OrderAmount: orderAmount, Bonus: bonus, BonusID: bonusID, OrderNote: orderNote}
	err := DB.Create(&order).Error
	return err
}

// -------------------------------删除-------------------------------

// -------------------------------更改-------------------------------

// 支付订单

// -------------------------------查询单个-------------------------------

// -------------------------------查询所有-------------------------------
