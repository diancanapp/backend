package base

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Fail 失败返回
func Fail(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"success": false, "data": data})
}

// FailUnprocessableEntity 参数错误
func FailUnprocessableEntity(c *gin.Context, data interface{}) {
	c.JSON(http.StatusUnprocessableEntity, gin.H{"success": false, "data": data})
}

// FailUnauthorized 没有权限
func FailUnauthorized(c *gin.Context, data interface{}) {
	c.JSON(http.StatusUnauthorized, gin.H{"success": false, "data": data})
}

// Ok 成功返回
func Ok(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"success": false, "data": data})
}
