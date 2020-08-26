package models

import (
	"github.com/jinzhu/gorm"
	"wozaizhao.com/diancan/common"
	"wozaizhao.com/diancan/util"
)

// Admin 用户
type Admin struct {
	gorm.Model
	UserName string `json:"username" gorm:"type:varchar(255);NOT NULL;DEFAULT ''"`
	Password string `json:"password" gorm:"type:varchar(255);DEFAULT ''"`
	Role     uint   `json:"role" gorm:"type:tinyint(1);DEFAULT '0'"`
}

// -------------------------------新建-------------------------------

// CreateAdmin 创建用户
func CreateAdmin(userName, password string, role uint) error {
	passwordMd5 := util.Md5(password)
	var admin = Admin{UserName: userName, Password: passwordMd5, Role: role}
	if err := DB.Create(&admin).Error; err != nil {
		common.Log("CreateAdmin", err)
		return err
	}
	return nil
}

// -------------------------------删除-------------------------------

// -------------------------------更改-------------------------------

// -------------------------------查询单个-------------------------------

// GetAdminByUserName 根据username获取Admin
func GetAdminByUserName(userName string) (Admin, error) {
	var admin Admin
	if err := DB.Where("user_name = ?", userName).First(&admin).Error; err != nil {
		common.Log("GetAdminByUserName", err)
		return admin, err
	}
	return admin, nil
}

// -------------------------------查询所有-------------------------------
