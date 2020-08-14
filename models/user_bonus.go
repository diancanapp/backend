package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// UserBonus 用户红包
type UserBonus struct {
	gorm.Model
	BonusID  uint       `json:"bonus_id" gorm:"type:smallint(5);NOT NULL;DEFAULT '0'"` // 红包ID
	BonusSn  string     `json:"bonus_sn" gorm:"type:varchar(60);NOT NULL;DEFAULT ''"`  // 红包码
	UserID   uint       `json:"user_id" gorm:"type:smallint(5);NOT NULL;DEFAULT '0'"`  // 用户ID
	UsedTime *time.Time `json:"used_time" gorm:"datetime"`                             // 使用时间
	OrderID  uint       `json:"order_id" gorm:"type:smallint(5);DEFAULT '0'"`          // 订单ID
}

// -------------------------------新建-------------------------------

// -------------------------------删除-------------------------------

// -------------------------------更改-------------------------------

// -------------------------------查询单个-------------------------------

// -------------------------------查询所有-------------------------------
