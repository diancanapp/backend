package models

import "github.com/jinzhu/gorm"

// Banner banner
type Banner struct {
	gorm.Model
	ImgSrc    string `json:"imgSrc" gorm:"type:varchar(255);NOT NULL;DEFAULT ''"`      // 图片地址
	Link      string `json:"link" gorm:"type:varchar(255);NOT NULL;DEFAULT ''"`        // 链接地址
	SortOrder int    `json:"sortOrder" gorm:"type:smallint(4);NOT NULL;DEFAULT '100'"` // 排序权重
}

// -------------------------------新建-------------------------------

// CreateBanner 创建Banner
func CreateBanner(imgSrc, link string, sortOrder int) error {
	var banner = Banner{ImgSrc: imgSrc, Link: link, SortOrder: sortOrder}
	err := DB.Create(&banner).Error
	return err
}

// -------------------------------删除-------------------------------

// DeleteBanner 删除Banner
func DeleteBanner(id uint) (err error) {
	var banner Banner
	err = DB.Where("id = ?", id).Delete(&banner).Error
	return err
}

// -------------------------------更改-------------------------------

// UpdateBanner 更新Banner
func UpdateBanner(id uint, imgSrc string, link string, sortOrder int) (err error) {
	var banner Banner
	banner.ID = id
	err = DB.Model(&banner).Updates(Banner{ImgSrc: imgSrc, Link: link, SortOrder: sortOrder}).Error
	return err
}

// -------------------------------查询单个-------------------------------

// QueryBannerByID 根据id查banner
func QueryBannerByID(id uint) (banner Banner, err error) {
	err = DB.Where("id = ?", id).Find(&banner).Error
	return banner, err
}

// -------------------------------查询所有-------------------------------

// QueryBanners 所有Banner
func QueryBanners(size, offset uint) (banners []Banner, err error) {
	err = DB.Limit(size).Offset(offset).Find(&banners).Error
	return banners, err
}
