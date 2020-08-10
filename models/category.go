package models

// Category 商品分类
type Category struct {
	Name         string `json:"name" gorm:"type:varchar(90);NOT NULL;DEFAULT ''"`          // 商品分类名
	CategoryDesc string `json:"goods_desc" gorm:"type:text;NOT NULL"`                      // 分类描述
	SortOrder    uint   `json:"sort_order" gorm:"type:smallint(1);NOT NULL;DEFAULT '50'"`  // 排序权重
	Style        string `json:"style" gorm:"type:varchar(60);NOT NULL;DEFAULT ''"`         // 分类名样式
	CategoryImg  string `json:"category_img" gorm:"type:varchar(255);NOT NULL;DEFAULT ''"` // 分类图
	IsShow       bool   `json:"is_show" gorm:"type:tinyint(1);NOT NULL;DEFAULT '0'"`       // 是否展示
	IsHot        bool   `json:"is_hot" gorm:"type:tinyint(1);NOT NULL;DEFAULT '0'"`        // 是否热门
}
