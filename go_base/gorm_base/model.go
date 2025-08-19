package main

type Student struct {
	ID    uint
	Name  string
	Age   int
	Grade string
}

// accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
type Account struct {
	ID      uint
	Balance int
}

type Transaction struct {
	ID            uint
	FromAccountId uint
	ToAccountId   uint
	Amount        int
}
