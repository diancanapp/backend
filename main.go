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
	// token, err := middleware.GenerateToken("beetle")
	// const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyTmFtZSI6ImJlZXRsZSIsImV4cCI6MTU5NzgyMDI2OCwiaXNzIjoid296YWl6aGFvIn0.X3rIns68va7dQpnxWIG7p6Oog08B-K9RcVqKV44b6iQ"
	// data, err := middleware.ParseToken(token)
	// fmt.Println(err)
	// fmt.Println(data)
	r := server.SetupRouter()
	r.Run(":8089")
}
