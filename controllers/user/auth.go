package user

import (
	"strings"

	"github.com/gin-gonic/gin"
	"wozaizhao.com/diancan/middleware"
)

const (
	bearer string = "bearer"
)

// 验证小程序token

// TokenValid 验证jwt token
func TokenValid(c *gin.Context) bool {
	token, ok := extractTokenFromAuthHeader(c.Request.Header.Get("Authorization"))
	if !ok {
		return false
	}
	data, err := middleware.ParseToken(token)
	if err != nil {
		return false
	}
	c.Set("username", data.Username)
	return true
}

func extractTokenFromAuthHeader(val string) (token string, ok bool) {
	authHeaderParts := strings.Split(val, " ")
	if len(authHeaderParts) != 2 || !strings.EqualFold(authHeaderParts[0], bearer) {
		return "", false
	}

	return authHeaderParts[1], true
}
