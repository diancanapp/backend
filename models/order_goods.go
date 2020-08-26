package models

import "github.com/jinzhu/gorm"

// OrderGoods 订单商品
type OrderGoods struct {
	gorm.Model
	OrderID     uint   `json:"orderID" gorm:"type:mediumint(8);NOT NULL;DEFAULT '0'"`    // 订单ID
	GoodsID     uint   `json:"goodsID" gorm:"type:mediumint(8);NOT NULL;DEFAULT '0'"`    // 商品ID
	GoodsName   string `json:"goodsName" gorm:"type:varchar(255);NOT NULL;DEFAULT ''"`   // 商品名称
	GoodsNumber uint   `json:"goodsNumber" gorm:"type:smallint(5);NOT NULL;DEFAULT '0'"` // 商品数量
	MarketPrice uint   `json:"marketPrice" gorm:"type:int(11);NOT NULL;DEFAULT '0'"`     // 市场价
	GoodsPrice  uint   `json:"goodsPrice" gorm:"type:int(11);NOT NULL;DEFAULT '0'"`      // 实际价
}

// -------------------------------新建-------------------------------

// CreateOrderGoods 创建订单商品
func CreateOrderGoods(orderID, goodsID uint, goodsName string, goodsNumber, marketPrice, goodsPrice uint) error {
	var orderGoods = OrderGoods{OrderID: orderID, GoodsID: goodsID, GoodsName: goodsName, GoodsNumber: goodsNumber, MarketPrice: marketPrice, GoodsPrice: goodsPrice}
	err := DB.Create(&orderGoods).Error
	return err
}

// -------------------------------删除-------------------------------

// -------------------------------更改-------------------------------

// -------------------------------查询单个-------------------------------

// QueryOrderGoodsByID 根据id查orderGoods
func QueryOrderGoodsByID(id uint) (orderGoods OrderGoods, err error) {
	err = DB.Where("id = ?", id).Find(&orderGoods).Error
	return orderGoods, err
}

// -------------------------------查询所有-------------------------------

// QueryOrderGoodsInOrder 查询订单所属商品
func QueryOrderGoodsInOrder(orderID uint) (orderGoods []OrderGoods, err error) {
	err = DB.Where("order_id = ?", orderID).Find(&orderGoods).Error
	return orderGoods, err
}

// QueryOrderGoods 所有OrderGoods
func QueryOrderGoods(size, offset uint) (orderGoods []OrderGoods, err error) {
	err = DB.Limit(size).Offset(offset).Find(&orderGoods).Error
	return orderGoods, err
}
