package business

import (
	"errors"

	"github.com/gin-gonic/gin"
	"wozaizhao.com/diancan/controllers/base"
	"wozaizhao.com/diancan/models"
)

var (
	errTokenMissing error = errors.New("Token is missing")
)

func parseWxToken(token string) (uint, error) {
	if token == "" {
		return 0, errTokenMissing
	}
	return models.GetUserIDByToken(token)
}

// GetMyBonus 获取我的红包
func GetMyBonus(c *gin.Context) {
	token := c.Request.Header.Get("token")
	userID, errParseWxToken := parseWxToken(token)
	if errParseWxToken != nil {
		base.Fail(c, errParseWxToken)
	}
	bonus, err := models.QueryUserBonusByUserID(userID)
	if err != nil {
		base.Fail(c, "Bonus not found")
	}
	base.Ok(c, bonus)
}

// MakeOrder
func MakeOrder(c *gin.Context) {
}

// PayOrder
func PayOrder(c *gin.Context) {
}

// GetMyOrders
func GetMyOrders(c *gin.Context) {
}

// GetMyOrder
func GetMyOrder(c *gin.Context) {
}

// SubmitFeedback
func SubmitFeedback(c *gin.Context) {
}
