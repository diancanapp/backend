package common

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	ds        string
	ak        string
	sk        string
	appid     string
	appsecret string
	jwtSecret string
)

// GetEnv 获取env
func GetEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ak = os.Getenv("QINIU_AK")
	sk = os.Getenv("QINIU_SK")

	usrname := os.Getenv("USRNAME")
	password := os.Getenv("PASSWORD")
	host := os.Getenv("HOST")
	database := os.Getenv("DATABASE")

	jwtSecret = os.Getenv("JWTSECRET")

	appid = os.Getenv("APPID")
	appsecret = os.Getenv("APPSECRET")
	ds = usrname + ":" + password + "@(" + host + ")/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"
	return
}

// GetDbSource 获取数据库源
func GetDbSource() (dbsource string) {
	dbsource = ds
	return
}

// GetWeappKey 获取七牛key
func GetWeappKey() (key map[string]string) {
	key = make(map[string]string)
	key["appid"] = appid
	key["appsecret"] = appsecret
	return
}

// GetQiniuKey 获取七牛key
func GetQiniuKey() (key map[string]string) {
	key = make(map[string]string)
	key["ak"] = ak
	key["sk"] = sk
	return
}

// getJwtSecret 获取jwtSecret
func GetJwtSecret() string {
	return jwtSecret
}
