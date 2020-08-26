package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Goods 商品
type Goods struct {
	gorm.Model
	Name             string     `json:"name" gorm:"type:varchar(120);NOT NULL;DEFAULT ''"`        // 商品名
	NameStyle        string     `json:"style" gorm:"type:varchar(60);DEFAULT ''"`                 // 商品名称样式
	CategoryID       uint       `json:"categoryID" gorm:"type:smallint(5);NOT NULL;DEFAULT '0'"`  // 商品分类
	ClickCount       uint       `gorm:"type:int(10);DEFAULT '0'"`                                 // 点击数
	GoodsNumber      uint       `json:"goodsNumber" gorm:"type:mediumint(8);DEFAULT '0'"`         // 库存
	WarnNumber       uint       `json:"warnNumber" gorm:"type:tinyint(3);DEFAULT '0'"`            // 警告库存
	MarketPrice      uint       `json:"marketPrice" gorm:"type:int(11);NOT NULL;DEFAULT '0'"`     // 市场价
	ShopPrice        uint       `json:"shopPrice" gorm:"type:int(11);NOT NULL;DEFAULT '0'"`       // 店铺价
	PromotePrice     uint       `json:"promotePrice" gorm:"type:int(11);DEFAULT '0'"`             // 活动价
	PromoteStartDate *time.Time `json:"promoteStartDate" gorm:"datetime"`                         // 活动开始时间
	PromoteEndDate   *time.Time `json:"promoteEndDate" gorm:"datetime"`                           // 活动结束时间
	Keywords         string     `json:"keywords" gorm:"type:varchar(255);DEFAULT ''"`             // 关键字
	GoodsBrief       string     `json:"goodsBrief" gorm:"type:varchar(255);NOT NULL;DEFAULT ''"`  // 商品简介
	GoodsDesc        string     `json:"goodsDesc" gorm:"type:text;NOT NULL"`                      // 商品描述
	GoodsThumb       string     `json:"goodsThumb" gorm:"type:varchar(255);NOT NULL;DEFAULT ''"`  // 缩略图
	GoodsImg         string     `json:"goodsImg" gorm:"type:varchar(255);NOT NULL;DEFAULT ''"`    // 商品图片
	SortOrder        uint       `json:"sortOrder" gorm:"type:smallint(4);NOT NULL;DEFAULT '100'"` // 排序权重
	IsOnSale         bool       `json:"isOnSale" gorm:"type:tinyint(1);NOT NULL;DEFAULT '1'"`     // 是否在售
	IsBest           bool       `json:"isBest" gorm:"type:tinyint(1);DEFAULT '0'"`                // 是否最佳
	IsNew            bool       `json:"isNew" gorm:"type:tinyint(1);DEFAULT '0'"`                 // 是否新品
	IsHot            bool       `json:"isHot" gorm:"type:tinyint(1);DEFAULT '0'"`                 // 是否热门
	IsPromote        bool       `json:"isPromote" gorm:"type:tinyint(1);DEFAULT '0'"`             // 是否活动
	Note             string     `json:"note" gorm:"type:varchar(255);DEFAULT ''"`                 // 备注
}

// -------------------------------新建-------------------------------

// CreateGoods 创建
func CreateGoods(name, namestyle string, categoryID, goodsNumber, warnNumber, marketPrice, shopPrice, promotePrice uint, promoteStartDate, promoteEndDate *time.Time, keywords, goodsBrief, goodsDesc, goodsThumb, goodsImg string, sortOrder uint, isOnSale, isBest, isNew, isHot, isPromote bool, note string) error {
	var goods = Goods{Name: name, NameStyle: namestyle, CategoryID: categoryID, GoodsNumber: goodsNumber, WarnNumber: warnNumber, MarketPrice: marketPrice, ShopPrice: shopPrice, PromotePrice: promotePrice, PromoteStartDate: promoteStartDate, PromoteEndDate: promoteEndDate, Keywords: keywords, GoodsBrief: goodsBrief, GoodsDesc: goodsDesc, GoodsThumb: goodsThumb, GoodsImg: goodsImg, SortOrder: sortOrder, IsOnSale: isOnSale, IsBest: isBest, IsNew: isNew, IsHot: isHot, IsPromote: isPromote, Note: note}
	err := DB.Create(&goods).Error
	return err
}

// -------------------------------删除-------------------------------

// -------------------------------更改-------------------------------

// -------------------------------查询单个-------------------------------

// QueryGoodsByID 根据id查goods
func QueryGoodsByID(id uint) (goods Goods, err error) {
	err = DB.Where("id = ?", id).Find(&goods).Error
	return goods, err
}

// -------------------------------查询所有-------------------------------

// QueryGoodsInCategory 商品分类下的商品
func QueryGoodsInCategory(categoryID, size, offset uint) (goods []Goods, err error) {
	err = DB.Where("category_id = ?", categoryID).Limit(size).Offset(offset).Find(&goods).Error
	return goods, err
}

// QueryGoods 所有Goods
func QueryGoods(size, offset uint) (goods []Goods, err error) {
	err = DB.Limit(size).Offset(offset).Find(&goods).Error
	return goods, err
}
