package src

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

var DBConn *gorm.DB
var err error

func init() {
	DBConn, err = gorm.Open("mysql", "root:105013sdta@/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	DBConn.LogMode(true)
	//todo 以下这几个值设置多少合适呢?
	DBConn.DB().SetMaxIdleConns(10)           //最大空闲
	DBConn.DB().SetMaxOpenConns(100)          //最大连接数
	DBConn.DB().SetConnMaxLifetime(time.Hour) //空闲时间
}
