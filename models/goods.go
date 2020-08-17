package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Goods 商品
type Goods struct {
	gorm.Model
	Name             string     `json:"name" gorm:"type:varchar(120);NOT NULL;DEFAULT ''"`         // 商品名
	NameStyle        string     `json:"style" gorm:"type:varchar(60);DEFAULT ''"`                  // 商品名称样式
	CategoryID       uint       `json:"category_id" gorm:"type:smallint(5);NOT NULL;DEFAULT '0'"`  // 商品分类
	ClickCount       uint       `gorm:"type:int(10);DEFAULT '0'"`                                  // 点击数
	GoodsNumber      uint       `json:"goods_number" gorm:"type:mediumint(8);DEFAULT '0'"`         // 库存
	WarnNumber       uint       `json:"warn_number" gorm:"type:tinyint(3);DEFAULT '0'"`            // 警告库存
	MarketPrice      uint       `json:"market_price" gorm:"type:int(11);NOT NULL;DEFAULT '0'"`     // 市场价
	ShopPrice        uint       `json:"shop_price" gorm:"type:int(11);NOT NULL;DEFAULT '0'"`       // 店铺价
	PromotePrice     uint       `json:"promote_price" gorm:"type:int(11);DEFAULT '0'"`             // 活动价
	PromoteStartDate *time.Time `json:"promote_start_date" gorm:"datetime"`                        // 活动开始时间
	PromoteEndDate   *time.Time `json:"promote_end_date" gorm:"datetime"`                          // 活动结束时间
	Keywords         string     `json:"keywords" gorm:"type:varchar(255);DEFAULT ''"`              // 关键字
	GoodsBrief       string     `json:"goods_brief" gorm:"type:varchar(255);NOT NULL;DEFAULT ''"`  // 商品简介
	GoodsDesc        string     `json:"goods_desc" gorm:"type:text;NOT NULL"`                      // 商品描述
	GoodsThumb       string     `json:"goods_thumb" gorm:"type:varchar(255);NOT NULL;DEFAULT ''"`  // 缩略图
	GoodsImg         string     `json:"goods_img" gorm:"type:varchar(255);NOT NULL;DEFAULT ''"`    // 商品图片
	SortOrder        uint       `json:"sort_order" gorm:"type:smallint(4);NOT NULL;DEFAULT '100'"` // 排序权重
	IsOnSale         bool       `json:"is_on_sale" gorm:"type:tinyint(1);NOT NULL;DEFAULT '1'"`    // 是否在售
	IsBest           bool       `json:"is_best" gorm:"type:tinyint(1);DEFAULT '0'"`                // 是否最佳
	IsNew            bool       `json:"is_new" gorm:"type:tinyint(1);DEFAULT '0'"`                 // 是否新品
	IsHot            bool       `json:"is_hot" gorm:"type:tinyint(1);DEFAULT '0'"`                 // 是否热门
	IsPromote        bool       `json:"is_promote" gorm:"type:tinyint(1);DEFAULT '0'"`             // 是否活动
	Note             string     `json:"note" gorm:"type:varchar(255);DEFAULT ''"`                  // 备注
}

// -------------------------------新建-------------------------------

// CreateGoods 创建
func CreateGoods(name string, namestyle string, categoryid uint, goodsnumber uint, warnnumber uint, marketprice uint, shopprice uint, promoteprice uint, promotestartdate *time.Time, promoteenddate *time.Time, keywords string, goodsbrief string, goodsdesc string, goodsthumb string, goodsimg string, sortorder uint, isonsale bool, isbest bool, isnew bool, ishot bool, ispromote bool, note string) error {
	var goods = Goods{Name: name, NameStyle: namestyle, CategoryID: categoryid, GoodsNumber: goodsnumber, WarnNumber: warnnumber, MarketPrice: marketprice, ShopPrice: shopprice, PromotePrice: promoteprice, PromoteStartDate: promotestartdate, PromoteEndDate: promoteenddate, Keywords: keywords, GoodsBrief: goodsbrief, GoodsDesc: goodsdesc, GoodsThumb: goodsthumb, GoodsImg: goodsimg, SortOrder: sortorder, IsOnSale: isonsale, IsBest: isbest, IsNew: isnew, IsHot: ishot, IsPromote: ispromote, Note: note}
	err := DB.Create(&goods).Error
	return err
}

// -------------------------------删除-------------------------------

// -------------------------------更改-------------------------------

// -------------------------------查询单个-------------------------------

// -------------------------------查询所有-------------------------------
