package util

import (
	"crypto/md5"
	"fmt"
	"io"
)

// Md5 md5
func Md5(src string) string {
	hashMd5 := md5.New()
	io.WriteString(hashMd5, src)
	return fmt.Sprintf("%x", hashMd5.Sum(nil))
}
