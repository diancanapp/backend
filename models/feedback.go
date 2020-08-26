package models

import "github.com/jinzhu/gorm"

// Feedback 反馈
type Feedback struct {
	gorm.Model
	UserID     uint   `json:"userID" gorm:"type:mediumint(8);NOT NULL;DEFAULT '0'"`    // 用户ID
	UserName   string `json:"userName" gorm:"type:varchar(255);DEFAULT ''"`            // 用户名
	UserPhone  string `json:"userPhone" gorm:"type:varchar(255);NOT NULL;DEFAULT ''"`  // 用户手机
	MsgTitle   string `json:"msgTitle" gorm:"type:varchar(255);NOT NULL;DEFAULT ''"`   // 消息标题
	MsgType    uint   `json:"msgType" gorm:"type:tinyint(1);NOT NULL;DEFAULT '0'"`     // 消息类型
	MsgStatus  uint   `json:"msgStatus" gorm:"type:tinyint(1);DEFAULT '0'"`            // 消息状态
	MsgContent string `json:"msgContent" gorm:"type:varchar(255);NOT NULL;DEFAULT ''"` // 消息内容
	MessageImg string `json:"messageImg" gorm:"type:varchar(255);DEFAULT ''"`          // 消息图片
	OrderID    uint   `json:"orderID" gorm:"type:mediumint(8);DEFAULT '0'"`            // 订单ID
}

// -------------------------------新建-------------------------------

// CreateFeedback 创建反馈
func CreateFeedback(userID uint, userName, userPhone, msgTitle string, msgType uint, msgContent, messageImg string, orderID uint) error {
	var feedback = Feedback{UserID: userID, UserName: userName, UserPhone: userPhone, MsgTitle: msgTitle, MsgType: msgType, MsgContent: msgContent, MessageImg: messageImg, OrderID: orderID}
	err := DB.Create(&feedback).Error
	return err
}

// -------------------------------删除-------------------------------

// -------------------------------更改-------------------------------

// -------------------------------查询单个-------------------------------

// QueryFeedbackByID 根据id查feedback
func QueryFeedbackByID(id uint) (feedback Feedback, err error) {
	err = DB.Where("id = ?", id).Find(&feedback).Error
	return feedback, err
}

// -------------------------------查询所有-------------------------------

// QueryFeedbacks 所有Feedback
func QueryFeedbacks(size, offset uint) (feedbacks []Feedback, err error) {
	err = DB.Limit(size).Offset(offset).Find(&feedbacks).Error
	return feedbacks, err
}
