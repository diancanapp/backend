package main

import (
	"fmt"

	"wozaizhao.com/diancan/common"
	"wozaizhao.com/diancan/models"
	"wozaizhao.com/diancan/server"
)

func main() {
	common.Loginit()
	common.GetEnv()
	models.DBinit()
	common.Log("main", "Server Start")
	res := models.UpdateBanner(2, "3333333333333333333", "555555555555", 300)
	fmt.Println(res)
	r := server.SetupRouter()
	r.Run(":8089")
}
