package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// UserBonus 用户红包
type UserBonus struct {
	gorm.Model
	BonusID  uint       `json:"bonusID" gorm:"type:smallint(5);NOT NULL;DEFAULT '0'"`       // 红包ID
	BonusSn  string     `json:"bonusSn" gorm:"unique;type:varchar(60);NOT NULL;DEFAULT ''"` // 红包码
	UserID   uint       `json:"userID" gorm:"type:smallint(5);NOT NULL;DEFAULT '0'"`        // 用户ID
	UsedTime *time.Time `json:"usedTime" gorm:"datetime"`                                   // 使用时间
	OrderID  uint       `json:"orderID" gorm:"type:smallint(5);DEFAULT '0'"`                // 订单ID
}

// -------------------------------新建-------------------------------

// CreateUserBonus 创建用户红包
func CreateUserBonus(bonusID uint, bonusSn string, userID uint) error {
	var userBonus = UserBonus{BonusID: bonusID, BonusSn: bonusSn, UserID: userID}
	err := DB.Create(&userBonus).Error
	return err
}

// -------------------------------删除-------------------------------

// -------------------------------更改-------------------------------

// -------------------------------查询单个-------------------------------

// QueryUserBonusByUserID 获取用户的红包
func QueryUserBonusByUserID(userID uint) (userBonus []UserBonus, err error) {
	err = DB.Where("userID = ?", userID).Find(&userBonus).Error
	return userBonus, err
}

// -------------------------------查询所有-------------------------------

// QueryBonus 所有Bonus
func QueryBonus(size, offset uint) (bonus []UserBonus, err error) {
	err = DB.Limit(size).Offset(offset).Find(&bonus).Error
	return bonus, err
}
