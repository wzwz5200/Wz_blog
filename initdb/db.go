package initdb

import (

	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func Initdb() {

	dsn := "host=localhost user=wz password=18326307873qq dbname=wz_test13 port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{

			SingularTable: true, // 使用单数表名

		},
	})

	if err != nil {
		fmt.Println("数据库连接错误", err)
	}

	DB = db

}
