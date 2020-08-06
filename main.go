package main

import (
	"wozaizhao.com/diancan/common"
	"wozaizhao.com/diancan/server"
)

func main() {
	common.Loginit()
	common.GetEnv()
	common.Log("main", "Server Start")
	r := server.SetupRouter()

	r.Run(":8089")
}
