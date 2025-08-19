package main

import "gorm.io/gorm"

// `User` 属于 `Company`，`CompanyID` 是外键
type User struct {
	gorm.Model
	Name        string
	CompanyID   int // 默认外键
	Company     Company
	IdCard      CreditCard   // 一对一
	MoneyCardss []MoneyCards // 一对多
	Languages   []Language   `gorm:"many2many:user_language"`
}

type Company struct {
	ID   int
	Name string
}

type CreditCard struct {
	gorm.Model
	CardName string
	UserID   int
}

type MoneyCards struct {
	gorm.Model
	CardName string
	UserID   int
}

type Language struct {
	gorm.Model
	LanguageName string
}
