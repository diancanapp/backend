package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"wozaizhao.com/diancan/common"
	"wozaizhao.com/diancan/controllers/base"
	"wozaizhao.com/diancan/middleware"
	"wozaizhao.com/diancan/models"
	"wozaizhao.com/diancan/util"
	"wozaizhao.com/diancan/wechat"
)

type wxLoginRequest struct {
	Code string `json:"code" form:"code" binding:"required"`
}

type wxUserInfo struct {
	EncryptedData string `json:"encryptedData"  binding:"required"`
	Iv            string `json:"iv" binding:"required"`
}

type loginRequest struct {
	UserName string `json:"userName"  binding:"required"`
	Password string `json:"password" binding:"required"`
}

// WxLogin 微信登录
func WxLogin(c *gin.Context) {
	var wxlreq wxLoginRequest
	if err := c.Bind(&wxlreq); err != nil {
		base.FailUnprocessableEntity(c, err)
		return
	}
	res, errCode2Session := wechat.Code2Session(wxlreq.Code)

	if errCode2Session != nil {
		common.Log("code2Session Error:", errCode2Session)
		base.Fail(c, "Login fail")
		return
	}
	token := util.Md5(res.SessionKey)
	if models.UserNotExistByOpenID(res.OpenID) {
		errCreateUser := models.CreateUser(res.OpenID, res.SessionKey, token)
		if errCreateUser != nil {
			base.Fail(c, "Login fail")
			return
		}
	} else {
		updateErr := models.UpdateSessionKey(res.OpenID, res.SessionKey, token)
		if updateErr != nil {
			base.Fail(c, "Login fail")
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": token})
}

// SaveUserInfo 保存用户信息
func SaveUserInfo(c *gin.Context) {
	var wxUser wxUserInfo
	if err := c.Bind(&wxUser); err != nil {
		base.FailUnprocessableEntity(c, err)
		return
	}
	token := c.Request.Header.Get("token")

	if token == "" {
		base.Fail(c, "Token is missing")
		return
	}
	sessionKey, errGetKey := models.GetSessionKeyByToken(token)
	if errGetKey != nil {
		base.Fail(c, "Save userinfo fail")
		return
	}

	user, errDecrypt := wechat.Decrypt(sessionKey, wxUser.EncryptedData, wxUser.Iv)
	if errDecrypt != nil {
		base.Fail(c, "Save userinfo fail")
		return
	}
	errUpdate := models.UpdateUserInfo(user.OpenID, user.UnionID, user.NickName, user.Gender, user.City, user.Province, user.Country, user.AvatarURL, user.Language, user.PhoneNumber, user.PurePhoneNumber, user.CountryCode)
	if errUpdate != nil {
		base.Fail(c, "Save userinfo fail")
		return
	}
	base.Ok(c, "Save userinfo successfully")
}

// GetCurrentUser 获取当前小程序用户
func GetCurrentUser(c *gin.Context) {
}

// Login 登录管理后台
func Login(c *gin.Context) {
	var loginReq loginRequest
	if err := c.Bind(&loginReq); err != nil {
		base.FailUnprocessableEntity(c, err)
		return
	}
	admin, errGetAdmin := models.GetAdminByUserName(loginReq.UserName)
	if errGetAdmin != nil {
		base.Fail(c, "Admin not found")
		return
	}

	passwordMd5 := util.Md5(loginReq.Password)

	if passwordMd5 != admin.Password {
		base.FailUnauthorized(c, "Username or password not match!")
		return
	}

	token, errGenerateToken := middleware.GenerateToken(admin.UserName)
	if errGenerateToken != nil {
		base.Fail(c, "Login fail")
		return
	}

	base.Ok(c, token)

}

// getCurrentAdmin 获取当前后台用户
func GetCurrentAdmin(c *gin.Context) {
	token := c.Request.Header.Get("token")

	if token == "" {
		base.FailUnauthorized(c, "Token is missing!")
	}
	data, err := middleware.ParseToken(token)

	if err != nil {
		base.FailUnauthorized(c, "Token is invalid!")
		return
	}

	admin, errGetAdmin := models.GetAdminByUserName(data.Username)

	if errGetAdmin != nil {
		base.FailUnauthorized(c, "get admin error")
		return
	}

	base.Ok(c, admin)

}
