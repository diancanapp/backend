package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"wozaizhao.com/diancan/common"
	"wozaizhao.com/diancan/models"
	"wozaizhao.com/diancan/util"
	"wozaizhao.com/diancan/wechat"
)

type wxLoginRequest struct {
	Code string `json:"code" form:"code" binding:"required"`
}

type wxUserInfo struct {
	EncryptedData string `json:"encrypted_data" form:"encrypted_data" binding:"required"`
	Iv            string `json:"iv" form:"iv" binding:"required"`
}

// WxLogin 微信登录
func WxLogin(c *gin.Context) {
	var wxlreq wxLoginRequest
	if err := c.Bind(&wxlreq); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"success": false, "message": err.Error()})
		return
	}
	res, err := wechat.Code2Session(wxlreq.Code)

	if err != nil {
		common.Log("code2Session Error:", err)
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "code2Session fail"})
		return
	}
	token := util.Md5(res.SessionKey)
	if models.UserNotExistByOpenID(res.OpenID) {
		createErr := models.CreateUser(res.OpenID, res.SessionKey, token)
		if createErr != nil {
			common.Log("CreateUser Error:", err)
			c.JSON(http.StatusOK, gin.H{"success": false, "message": "CreateUser fail"})
		}
	} else {
		updateErr := models.UpdateSessionKey(res.OpenID, res.SessionKey, token)
		if updateErr != nil {
			common.Log("UpdateSessionKey Error:", err)
			c.JSON(http.StatusOK, gin.H{"success": false, "message": "UpdateSessionKey fail"})
		}
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": "login successfully"})
}

// SaveUserInfo 保存用户信息
func SaveUserInfo(c *gin.Context) {
	var wxUser wxUserInfo
	if err := c.Bind(&wxUser); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"success": false, "message": err.Error()})
		return
	}

	token := c.Request.Header["Token"]

	if token == nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "token missing"})
		return
	}
	sessionKey, _ := models.GetSessionKeyByToken(token[0])
	userInfo, _ := wechat.Decrypt(sessionKey, wxUser.EncryptedData, wxUser.Iv)
	common.Log("userinfo", userInfo)
}
