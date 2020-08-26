package models

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // mysql
	"wozaizhao.com/diancan/common"
)

// DB 数据库
var DB *gorm.DB

// Models 数据库实体
var models = []interface{}{
	&Banner{}, &Bonus{}, &Category{}, &Config{}, &Feedback{}, &Goods{}, &OrderGoods{}, &Order{}, &Qiniu{}, &UserBonus{}, &User{}, &Admin{},
}

// DBinit 数据库初始化
func DBinit() {
	if db, err := gorm.Open("mysql", common.GetDbSource()); err != nil {
		common.Log("DBinit Error:", err)
		os.Exit(0)
	} else {
		common.Log("DBinit Success:", "DBinit Success")
		DB = db
		DB.DB().SetMaxIdleConns(10)
		DB.DB().SetMaxOpenConns(20)
		DB.LogMode(true)
		if err = db.AutoMigrate(models...).Error; nil != err {
			common.Log("auto migrate tables failed:", err.Error())
		}
	}

}
