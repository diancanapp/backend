package models

import (
	"github.com/jinzhu/gorm"
	"wozaizhao.com/diancan/common"
)

// User 用户
type User struct {
	gorm.Model
	OpenID          string `json:"openId" gorm:"type:varchar(255);NOT NULL;DEFAULT ''"`
	UnionID         string `json:"unionId" gorm:"type:varchar(255);DEFAULT ''"`
	NickName        string `json:"nickName" gorm:"type:varchar(255);DEFAULT ''"`
	Gender          int    `json:"gender" gorm:"type:tinyint(1);DEFAULT '0'"`
	City            string `json:"city" gorm:"type:varchar(20);DEFAULT ''"`
	Province        string `json:"province" gorm:"type:varchar(20);DEFAULT ''"`
	Country         string `json:"country" gorm:"type:varchar(20);DEFAULT ''"`
	AvatarURL       string `json:"avatarUrl" gorm:"type:varchar(255);DEFAULT ''"`
	Language        string `json:"language" gorm:"type:varchar(20);DEFAULT ''"`
	PhoneNumber     string `json:"phoneNumber" gorm:"type:varchar(11);DEFAULT ''"`
	PurePhoneNumber string `json:"purePhoneNumber" gorm:"type:varchar(11);DEFAULT ''"`
	CountryCode     string `json:"countryCode"  gorm:"type:varchar(20);DEFAULT ''"`
	SessionKey      string `json:"session_key"`
	Token           string `json:"token"`
}

// -------------------------------新建-------------------------------

// CreateUser 创建用户
func CreateUser(openID string, sessionKey string, token string) error {
	var user = User{OpenID: openID, SessionKey: sessionKey, Token: token}
	err := DB.Create(&user).Error
	return err
}

// -------------------------------删除-------------------------------

// -------------------------------更改-------------------------------

// UpdateSessionKey 更新用户sessionKey
func UpdateSessionKey(openID string, sessionKey string, token string) error {
	var user User
	err := DB.Model(&user).Where("open_id = ?", openID).Updates(User{SessionKey: sessionKey, Token: token}).Error
	return err
}

// -------------------------------查询单个-------------------------------

// GetUserByOpenID 根据openid获取用户
func GetUserByOpenID(openID string) (User, error) {
	var user User
	if err := DB.Where("open_id = ?", openID).First(&user).Error; err != nil {
		common.Log("GetUserByOpenID", err)
		return user, err
	}
	return user, nil
}

// UserNotExistByOpenID 用户是否存在
func UserNotExistByOpenID(openID string) (res bool) {
	var user User
	res = DB.Where("open_id = ?", openID).First(&user).RecordNotFound()
	return
}

// GetSessionKeyByToken 根据token获取sessionKey
func GetSessionKeyByToken(token string) (string, error) {
	var user User
	if err := DB.Where("token = ?", token).First(&user).Error; err != nil {
		common.Log("GetSessionKeyByToken", err)
		return user.Token, err
	}
	return user.Token, nil
}

// -------------------------------查询所有-------------------------------
