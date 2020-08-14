package models

import "github.com/jinzhu/gorm"

// Banner banner
type Banner struct {
	gorm.Model
	ImgSrc    string `json:"img_src" gorm:"type:varchar(255);NOT NULL;DEFAULT ''"`      // 图片地址
	Link      string `json:"link" gorm:"type:varchar(255);NOT NULL;DEFAULT ''"`         // 链接地址
	SortOrder uint   `json:"sort_order" gorm:"type:smallint(4);NOT NULL;DEFAULT '100'"` // 排序权重
}

// -------------------------------新建-------------------------------

// -------------------------------删除-------------------------------

// -------------------------------更改-------------------------------

// -------------------------------查询单个-------------------------------

// -------------------------------查询所有-------------------------------
