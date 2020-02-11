package AppInit

import (
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

var dbConn *gorm.DB
var err error

//init引入包时就执行
func init() {
	dbConn, err = gorm.Open("mysql", "root:105013sdta@/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	//开启调试模式
	dbConn.LogMode(true)
	//不开启表明复数模式
	dbConn.SingularTable(true)
	//todo 以下这几个值设置多少合适呢?
	dbConn.DB().SetMaxIdleConns(10)           //最大空闲
	dbConn.DB().SetMaxOpenConns(100)          //最大连接数
	dbConn.DB().SetConnMaxLifetime(time.Hour) //空闲时间
}

//返回db连接
func GetDB() *gorm.DB {
	return dbConn
}
