package common

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

// Loginit 初始化日志
func Loginit() {
	if file, err := os.OpenFile("diancan.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666); err == nil {
		logrus.SetOutput(file)
	} else {
		logrus.Error(err)
	}
}

// Log 记录日志
func Log(ns string, log interface{}) {
	logrus.Error(fmt.Sprintf("[%s],%s", ns, log))
}
