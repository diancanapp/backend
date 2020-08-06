package common

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	ds string
	ak string
	sk string
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
	port := os.Getenv("PORT")
	database := os.Getenv("DATABASE")

	ds = usrname + ":" + password + "@tcp(" + host + ":" + port + ")/" + database

	return
}

// GetQiniuKey 获取七牛key
func GetQiniuKey() (key map[string]string) {
	key = make(map[string]string)
	key["ak"] = ak
	key["sk"] = sk
	return
}
