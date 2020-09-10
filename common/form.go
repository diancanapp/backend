package common

import "time"

// AddCategory form
type AddCategory struct {
	Name         string `json:"name" binding:"required"`
	CategoryDesc string `json:"categoryDesc" binding:"required"`
	SortOrder    uint   `json:"sortOrder"`
	Style        string `json:"style"`
	CategoryImg  string `json:"categoryImg"`
	IsShow       bool   `json:"isShow"`
	IsHot        bool   `json:"isHot"`
}

// ModifyCategory form
type ModifyCategory struct {
	ID           uint   `json:"ID" binding:"required"`
	Name         string `json:"name"`
	CategoryDesc string `json:"categoryDesc"`
	SortOrder    uint   `json:"sortOrder"`
	Style        string `json:"style"`
	CategoryImg  string `json:"categoryImg"`
	IsShow       bool   `json:"isShow"`
	IsHot        bool   `json:"isHot"`
}

// DelCategory form
type DelCategory struct {
	ID uint `json:"ID" binding:"required"`
}

// AddBanner form
type AddBanner struct {
	ImgSrc    string `json:"imgSrc" binding:"required"`
	Link      string `json:"link"`
	SortOrder int    `json:"sortOrder"`
}

// ModifyBanner form
type ModifyBanner struct {
	ID        uint   `json:"ID" binding:"required"`
	ImgSrc    string `json:"imgSrc"`
	Link      string `json:"link"`
	SortOrder int    `json:"sortOrder"`
}

// DelBanner form
type DelBanner struct {
	ID uint `json:"ID" binding:"required"`
}

// AddGoods form
type AddGoods struct {
	Name             string     `json:"name" binding:"required"`
	NameStyle        string     `json:"style"`
	CategoryID       uint       `json:"categoryID" binding:"required"`
	GoodsNumber      uint       `json:"goodsNumber"`
	WarnNumber       uint       `json:"warnNumber"`
	MarketPrice      uint       `json:"marketPrice" binding:"required"`
	ShopPrice        uint       `json:"shopPrice" binding:"required"`
	PromotePrice     uint       `json:"promotePrice"`
	PromoteStartDate *time.Time `json:"promoteStartDate"`
	PromoteEndDate   *time.Time `json:"promoteEndDate"`
	Keywords         string     `json:"keywords"`
	GoodsBrief       string     `json:"goodsBrief" binding:"required"`
	GoodsDesc        string     `json:"goodsDesc" binding:"required"`
	GoodsThumb       string     `json:"goodsThumb" binding:"required"`
	GoodsImg         string     `json:"goodsImg" binding:"required"`
	SortOrder        uint       `json:"sortOrder"`
	IsOnSale         bool       `json:"isOnSale"`
	IsBest           bool       `json:"isBest"`
	IsNew            bool       `json:"isNew"`
	IsHot            bool       `json:"isHot"`
	IsPromote        bool       `json:"isPromote"`
	Note             string     `json:"note"`
}

// ModifyGoods form
type ModifyGoods struct {
	ID               uint       `json:"ID" binding:"required"`
	Name             string     `json:"name"`
	NameStyle        string     `json:"style"`
	CategoryID       uint       `json:"categoryID"`
	GoodsNumber      uint       `json:"goodsNumber"`
	WarnNumber       uint       `json:"warnNumber"`
	MarketPrice      uint       `json:"marketPrice"`
	ShopPrice        uint       `json:"shopPrice"`
	PromotePrice     uint       `json:"promotePrice"`
	PromoteStartDate *time.Time `json:"promoteStartDate"`
	PromoteEndDate   *time.Time `json:"promoteEndDate"`
	Keywords         string     `json:"keywords"`
	GoodsBrief       string     `json:"goodsBrief"`
	GoodsDesc        string     `json:"goodsDesc"`
	GoodsThumb       string     `json:"goodsThumb"`
	GoodsImg         string     `json:"goodsImg"`
	SortOrder        uint       `json:"sortOrder"`
	IsOnSale         bool       `json:"isOnSale"`
	IsBest           bool       `json:"isBest"`
	IsNew            bool       `json:"isNew"`
	IsHot            bool       `json:"isHot"`
	IsPromote        bool       `json:"isPromote"`
	Note             string     `json:"note"`
}

// DelGoods form
type DelGoods struct {
	ID uint `json:"ID" binding:"required"`
}

// OrderGoods struct
type OrderGoods struct {
	GoodsID     uint   `json:"goodsID"`
	GoodsName   string `json:"goodsName"`
	GoodsNumber uint   `json:"goodsNumber"`
	MarketPrice uint   `json:"marketPrice"`
	GoodsPrice  uint   `json:"goodsPrice"`
}

// PostOrder form
type PostOrder struct {
	UserID      uint         `json:"userID" binding:"required"`
	OrderAmount uint         `json:"orderAmount" binding:"required"`
	Bonus       uint         `json:"bonus"`
	BonusID     uint         `json:"bonusID"`
	Goods       []OrderGoods `json:"goods" binding:"required"`
	OrderNote   string       `json:"orderNote"`
}
