package models

import "github.com/jinzhu/gorm"

// Feedback 反馈
type Feedback struct {
	gorm.Model
	UserID     uint   `json:"user_id" gorm:"type:mediumint(8);NOT NULL;DEFAULT '0'"`    // 用户ID
	UserName   string `json:"user_name" gorm:"type:varchar(255);DEFAULT ''"`            // 用户名
	UserPhone  string `json:"user_phone" gorm:"type:varchar(255);NOT NULL;DEFAULT ''"`  // 用户手机
	MsgTitle   string `json:"msg_title" gorm:"type:varchar(255);NOT NULL;DEFAULT ''"`   // 消息标题
	MsgType    uint   `json:"msg_type" gorm:"type:tinyint(1);NOT NULL;DEFAULT '0'"`     // 消息类型
	MsgStatus  uint   `json:"msg_status" gorm:"type:tinyint(1);NOT NULL;DEFAULT '0'"`   // 消息状态
	MsgContent string `json:"msg_content" gorm:"type:varchar(255);NOT NULL;DEFAULT ''"` // 消息内容
	MessageImg string `json:"message_img" gorm:"type:varchar(255);DEFAULT ''"`          // 消息图片
	OrderID    uint   `json:"order_id" gorm:"type:mediumint(8);DEFAULT '0'"`            // 订单ID
}

// -------------------------------新建-------------------------------

// -------------------------------删除-------------------------------

// -------------------------------更改-------------------------------

// -------------------------------查询单个-------------------------------

// -------------------------------查询所有-------------------------------
