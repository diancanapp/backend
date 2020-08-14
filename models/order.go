package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Order 订单
type Order struct {
	gorm.Model
	UserID      uint       `json:"user_id" gorm:"type:mediumint(8);NOT NULL;DEFAULT '0'"`    // 用户ID
	OrderStatus uint       `json:"order_status" gorm:"type:tinyint(1);NOT NULL;DEFAULT '0'"` // 订单状态
	PayStatus   uint       `json:"pay_status" gorm:"type:tinyint(1);NOT NULL;DEFAULT '0'"`   // 支付状态
	PayID       uint       `json:"pay_id" gorm:"type:mediumint(8);NOT NULL;DEFAULT '0'"`     // 支付ID
	OrderAmount uint       `json:"order_amount" gorm:"type:int(11);NOT NULL;DEFAULT '0'"`    // 订单金额
	Bonus       uint       `json:"bonus" gorm:"type:int(11);NOT NULL;DEFAULT '0'"`           // 红包
	BonusID     uint       `json:"bonus_id" gorm:"type:mediumint(8);NOT NULL;DEFAULT '0'"`   // 红包ID
	MoneyPaid   uint       `json:"money_paid" gorm:"type:int(11);NOT NULL;DEFAULT '0'"`      // 支付金额
	ConfirmTime *time.Time `json:"confirm_time" gorm:"datetime"`                             // 下单时间
	PayTime     *time.Time `json:"pay_time" gorm:"datetime"`                                 // 支付时间
	PayNote     string     `json:"pay_note" gorm:"type:varchar(255);NOT NULL;DEFAULT ''"`    // 备注
	OrderType   uint       `json:"order_type" gorm:"type:tinyint(1);NOT NULL;DEFAULT '0'"`   // 订单类型
}

// -------------------------------新建-------------------------------

// -------------------------------删除-------------------------------

// -------------------------------更改-------------------------------

// -------------------------------查询单个-------------------------------

// -------------------------------查询所有-------------------------------
