package main

import (
	"wozaizhao.com/diancan/common"
	"wozaizhao.com/diancan/models"
	"wozaizhao.com/diancan/server"
)

func main() {
	common.Loginit()
	common.GetEnv()
	models.DBinit()
	common.Log("main", "Server Start")
	r := server.SetupRouter()
	r.Run(":8089")
}
