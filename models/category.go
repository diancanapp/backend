package models

import (
	"github.com/jinzhu/gorm"
	"wozaizhao.com/diancan/common"
)

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
func CreateCategory(c common.AddCategory) error {
	var cate = Category{Name: c.Name, CategoryDesc: c.CategoryDesc, SortOrder: c.SortOrder, CategoryImg: c.CategoryImg, IsShow: c.IsShow, IsHot: c.IsHot}
	err := DB.Create(&cate).Error
	if err != nil {
		common.Log("CreateCategory Error", err)
	} else {
		common.Log("CreateCategory Success", err)
	}
	return err
}

// -------------------------------删除-------------------------------

// DeleteCategory 删除Category
func DeleteCategory(id uint) (err error) {
	var category Category
	err = DB.Where("id = ?", id).Delete(&category).Error
	return err
}

// -------------------------------更改-------------------------------

// UpdateCategory 更新Category
func UpdateCategory(c common.ModifyCategory) (err error) {
	var category Category
	category.ID = c.ID
	err = DB.Model(&category).Updates(Category{Name: c.Name, CategoryDesc: c.CategoryDesc, SortOrder: c.SortOrder, CategoryImg: c.CategoryImg, IsShow: c.IsShow, IsHot: c.IsHot}).Error
	return err
}

// -------------------------------查询单个-------------------------------

// QueryCategoryByID 根据id查category
func QueryCategoryByID(id uint) (category Category, err error) {
	err = DB.Where("id = ?", id).Find(&category).Error
	return category, err
}

// -------------------------------查询所有-------------------------------

// QueryCategorys 所有Category
func QueryCategorys() (categories []Category, err error) {
	err = DB.Find(&categories).Error
	return categories, err
}
