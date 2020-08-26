package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Bonus 红包
type Bonus struct {
	gorm.Model
	Name           string     `json:"name" gorm:"type:varchar(90);NOT NULL;DEFAULT ''"`        // 红包名
	Money          uint       `json:"money" gorm:"type:int(11);NOT NULL;DEFAULT '0'"`          // 红包金额
	SendType       uint       `json:"sendType" gorm:"type:tinyint(1);NOT NULL;DEFAULT '1'"`    // 红包类型
	MinAmount      uint       `json:"minAmount" gorm:"type:int(11);NOT NULL;DEFAULT '0'"`      // 最小金额
	MaxAmount      uint       `json:"maxAmount" gorm:"type:int(11);NOT NULL;DEFAULT '0'"`      // 最大金额
	SendStartDate  *time.Time `json:"sendStartDate" gorm:"datetime"`                           // 发送开始时间
	SendEndDate    *time.Time `json:"sendEndDate" gorm:"datetime"`                             // 发送结束时间
	UseStartDate   *time.Time `json:"useStartDate" gorm:"datetime"`                            // 使用开始时间
	UseEndDate     *time.Time `json:"useEndDate" gorm:"datetime"`                              // 使用结束时间
	MinOrderAmount uint       `json:"minOrderAmount" gorm:"type:int(11);NOT NULL;DEFAULT '0'"` // 最小使用金额
}

// -------------------------------新建-------------------------------

// CreateBonus 创建红包
func CreateBonus(name string, money, sendType, minAmount, maxAmount uint, sendStartDate, sendEndDate, useStartDate, useEndDate *time.Time, minOrderAmount uint) error {
	var bonus = Bonus{Name: name, Money: money, SendType: sendType, MinAmount: minAmount, MaxAmount: maxAmount, SendStartDate: sendStartDate, SendEndDate: sendEndDate, UseStartDate: useStartDate, UseEndDate: useEndDate, MinOrderAmount: minOrderAmount}
	err := DB.Create(&bonus).Error
	return err
}

// -------------------------------删除-------------------------------

// -------------------------------更改-------------------------------

// -------------------------------查询单个-------------------------------

// -------------------------------查询所有-------------------------------
