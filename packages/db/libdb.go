package libdb

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var conn *gorm.DB

func Connect() {
	c, err := gorm.Open("mysql", "<username>:<password>@(127.0.0.1:3306)/library?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	conn = c
	fmt.Println(conn)
}

func GetConn() *gorm.DB {
	return conn
}
