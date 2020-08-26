package models

import "github.com/jinzhu/gorm"

// Category 商品分类
type Category struct {
	gorm.Model
	Name         string `json:"name" gorm:"type:varchar(90);NOT NULL;DEFAULT ''"`        // 商品分类名
	CategoryDesc string `json:"categoryDesc" gorm:"type:text;NOT NULL"`                  // 分类描述
	SortOrder    uint   `json:"sortOrder" gorm:"type:smallint(1);NOT NULL;DEFAULT '50'"` // 排序权重
	Style        string `json:"style" gorm:"type:varchar(60);DEFAULT ''"`                // 分类名样式
	CategoryImg  string `json:"categoryImg" gorm:"type:varchar(255);DEFAULT ''"`         // 分类图
	IsShow       bool   `json:"isShow" gorm:"type:tinyint(1);DEFAULT '1'"`               // 是否展示
	IsHot        bool   `json:"isHot" gorm:"type:tinyint(1);DEFAULT '0'"`                // 是否热门
}

// -------------------------------新建-------------------------------

// CreateCategory 创建商品分类
func CreateCategory(name, categoryDesc string, sortOrder uint, categoryImg string, isShow, isHot bool) error {
	var cate = Category{Name: name, CategoryDesc: categoryDesc, SortOrder: sortOrder, CategoryImg: categoryImg, IsShow: isShow, IsHot: isHot}
	err := DB.Create(&cate).Error
	return err
}

// -------------------------------删除-------------------------------

// -------------------------------更改-------------------------------

// -------------------------------查询单个-------------------------------

// QueryCategoryByID 根据id查category
func QueryCategoryByID(id uint) (category Category, err error) {
	err = DB.Where("id = ?", id).Find(&category).Error
	return category, err
}

// -------------------------------查询所有-------------------------------

// QueryCategorys 所有Category
func QueryCategorys(size, offset uint) (categories []Category, err error) {
	err = DB.Limit(size).Offset(offset).Find(&categories).Error
	return categories, err
}
