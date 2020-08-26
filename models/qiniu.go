package models

import "github.com/jinzhu/gorm"

// Qiniu 七牛云记录
type Qiniu struct {
	gorm.Model
	Key     string `json:"key" gorm:"type:varchar(255);NOT NULL;DEFAULT ''"` // 图片地址
	ImgType uint   `json:"type" gorm:"type:tinyint(1);DEFAULT '0'"`          // 订单类型
}

// -------------------------------新建-------------------------------

// CreateQiniu 创建七牛记录
func CreateQiniu(key string, imgType uint) error {
	var qiniu = Qiniu{Key: key, ImgType: imgType}
	err := DB.Create(&qiniu).Error
	return err
}

// -------------------------------删除-------------------------------

// -------------------------------更改-------------------------------

// -------------------------------查询单个-------------------------------

// QueryQiniuByID 根据id查qiniu
func QueryQiniuByID(id uint) (qiniu Qiniu, err error) {
	err = DB.Where("id = ?", id).Find(&qiniu).Error
	return qiniu, err
}

// -------------------------------查询所有-------------------------------

// QueryQinius 所有Qiniu
func QueryQinius(size, offset uint) (qinius []Qiniu, err error) {
	err = DB.Limit(size).Offset(offset).Find(&qinius).Error
	return qinius, err
}
