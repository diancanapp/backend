package models

import "github.com/jinzhu/gorm"

// Config 配置
type Config struct {
	gorm.Model
	Key   string `json:"key" gorm:"type:varchar(255);NOT NULL;DEFAULT ''"`   // 键
	Value string `json:"value" gorm:"type:varchar(255);NOT NULL;DEFAULT ''"` // 值
}

// -------------------------------新建-------------------------------

// CreateConfig 创建配置
func CreateConfig(key string, value string) error {
	var config = Config{Key: key, Value: value}
	err := DB.Create(&config).Error
	return err
}

// -------------------------------删除-------------------------------

// -------------------------------更改-------------------------------

// -------------------------------查询单个-------------------------------

// -------------------------------查询所有-------------------------------
