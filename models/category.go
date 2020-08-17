package models

import "github.com/jinzhu/gorm"

// Category 商品分类
type Category struct {
	gorm.Model
	Name         string `json:"name" gorm:"type:varchar(90);NOT NULL;DEFAULT ''"`         // 商品分类名
	CategoryDesc string `json:"cate_desc" gorm:"type:text;NOT NULL"`                      // 分类描述
	SortOrder    uint   `json:"sort_order" gorm:"type:smallint(1);NOT NULL;DEFAULT '50'"` // 排序权重
	Style        string `json:"style" gorm:"type:varchar(60);DEFAULT ''"`                 // 分类名样式
	CategoryImg  string `json:"category_img" gorm:"type:varchar(255);DEFAULT ''"`         // 分类图
	IsShow       bool   `json:"is_show" gorm:"type:tinyint(1);DEFAULT '1'"`               // 是否展示
	IsHot        bool   `json:"is_hot" gorm:"type:tinyint(1);DEFAULT '0'"`                // 是否热门
}

// -------------------------------新建-------------------------------

// CreateCategory 创建商品分类
func CreateCategory(name string, categoryDesc string, categoryImg string, isShow bool, isHot bool) error {
	var cate = Category{Name: name, CategoryDesc: categoryDesc, CategoryImg: categoryImg, IsShow: isShow, IsHot: isHot}
	err := DB.Create(&cate).Error
	return err
}

// -------------------------------删除-------------------------------

// -------------------------------更改-------------------------------

// -------------------------------查询单个-------------------------------

// -------------------------------查询所有-------------------------------
