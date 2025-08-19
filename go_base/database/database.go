package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	createDb()
}

func createDb() *gorm.DB {
	dsn := "root:youda123@@tcp(rm-bp1y051g2ury93z6ufo.mysql.rds.aliyuncs.com:3306)/kt_ruoyi?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println(err)
	// fmt.Println(err)
	db.AutoMigrate(&Company{})    // 多对一
	db.AutoMigrate(&User{})       //
	db.AutoMigrate(&CreditCard{}) // 一对一
	db.AutoMigrate(&MoneyCards{}) // 一对多
	db.AutoMigrate(&Language{})   // 多对多

	db.Create(&Company{ID: 1, Name: "yonganxing"})
	db.Create(&User{Name: "king", CompanyID: 1})
	db.Create(&CreditCard{CardName: "id_card", UserID: 1})
	db.Create(&MoneyCards{CardName: "ICBC_CARD", UserID: 1})
	db.Create(&MoneyCards{CardName: "BBC_CARD", UserID: 1})
	db.Create(&Language{LanguageName: "ENGLISH"})
	db.Create(&Language{LanguageName: "CHINESE"})

	user := User{}
	db.Debug().Preload("Company").Preload("IdCard").Preload("MoneyCardss").First(&user, 1)

	fmt.Println(user)
	return db
}
