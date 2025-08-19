package main

import (
	"errors"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:youda123@@tcp(rm-bp1y051g2ury93z6ufo.mysql.rds.aliyuncs.com:3306)/kt_ruoyi?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println(err)
	// one
	// testStudent(db)
	testAccount(db)
}

/*
	题目1：基本CRUD操作

假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
要求 ：
编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
*/
func testStudent(db *gorm.DB) {
	db.AutoMigrate(&Student{})

	student := Student{Name: "张三", Age: 20, Grade: "三年级"}
	db.Debug().Create(&student)
	fmt.Println(student)

	queryStudent := []Student{}
	db.Debug().Model(&Student{}).Where("Age > ?", 20).Find(&queryStudent)
	fmt.Println(queryStudent)

	db.Debug().Model(&Student{}).Where("Name = ?", "张三").Update("Grade", "四年级")

	db.Debug().Where("Age < ?", 15).Delete(&Student{})
}

/*
	题目2：事务语句

假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
要求 ：
编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。
*/
func testAccount(db *gorm.DB) {
	db.AutoMigrate(&Account{})
	db.AutoMigrate(&Transaction{})

	aAccount := Account{Balance: 200}
	bAccount := Account{Balance: 0}
	db.Create(&aAccount)
	db.Create(&bAccount)

	for i := 0; i < 3; i++ {
		errorInfo := transfrom(aAccount.ID, bAccount.ID, 100, db)
		if errorInfo != "" {
			fmt.Println("trans error", i, errorInfo)
		}
		fmt.Printf("trans times = %d次", i+1)
	}
}

func transfrom(fromAccountId uint, toAccountId uint, transAmount int, db *gorm.DB) string {

	errorInfo := db.Transaction(func(tx *gorm.DB) error {
		fromAccount := Account{ID: fromAccountId}
		err := tx.Debug().Find(&fromAccount)
		if err.Error != nil {
			return errors.New("账户不存在")
		}

		if fromAccount.Balance < transAmount {
			return errors.New("账户余额不足，无法转账")
		}

		tx.Debug().Model(&fromAccount).Update("Balance", gorm.Expr("Balance - ?", transAmount))
		toAccount := Account{ID: toAccountId}
		tx.Debug().Model(&toAccount).Update("Balance", gorm.Expr("Balance + ?", transAmount))

		transactionInfo := Transaction{FromAccountId: fromAccount.ID, ToAccountId: toAccount.ID, Amount: transAmount}
		tx.Debug().Create(&transactionInfo)

		return nil
	})

	if errorInfo != nil {
		return errorInfo.Error()
	}

	return ""
}
