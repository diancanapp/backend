package models

import "github.com/jinzhu/gorm"

// OrderGoods 订单商品
type OrderGoods struct {
	gorm.Model
	OrderID     uint   `json:"order_id" gorm:"type:mediumint(8);NOT NULL;DEFAULT '0'"`    // 订单ID
	GoodsID     uint   `json:"goods_id" gorm:"type:mediumint(8);NOT NULL;DEFAULT '0'"`    // 商品ID
	GoodsName   string `json:"goods_name" gorm:"type:varchar(255);NOT NULL;DEFAULT ''"`   // 商品名称
	GoodsNumber uint   `json:"goods_number" gorm:"type:smallint(5);NOT NULL;DEFAULT '0'"` // 商品数量
	MarketPrice uint   `json:"market_price" gorm:"type:int(11);NOT NULL;DEFAULT '0'"`     // 市场价
	GoodsPrice  uint   `json:"goods_price" gorm:"type:int(11);NOT NULL;DEFAULT '0'"`      // 实际价
}

// -------------------------------新建-------------------------------

// -------------------------------删除-------------------------------

// -------------------------------更改-------------------------------

// -------------------------------查询单个-------------------------------

// -------------------------------查询所有-------------------------------
