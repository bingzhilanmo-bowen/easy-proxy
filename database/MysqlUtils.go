package database


import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)



func GormDb() *gorm.DB{
	fmt.Println("init mysql database !!!!!")
	db, err := gorm.Open("mysql","root:hcb_inf_test_mysql@tcp(10.6.1.35:3306)/local_test?parseTime=true")
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
