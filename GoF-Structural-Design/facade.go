package main

import (
	"fmt"
)

type Account struct {
	id string
	accountType string
}

// Account class method create - creates account given AccountType
func (account *Account) create(accountType string) *Account {
	fmt.Println("account creation with type")
	account.accountType = accountType
	return account
}

func (account *Account) getById(id string) *Account {
	fmt.Println("getting account by Id")
	return account
}

func (account *Account) deleteById(id string) () {
	fmt.Println("delete ccount by id")
}

type Customer struct {
	name string
	id int
}

// Customer class method craete - create Customer given name
func (customer *Customer) create(name string) *Customer {
	fmt.Println("creating customer")
	customer.name = name
	return customer
}

type Transaction struct {
	id string
	amount float32
	srcAccountId string
	destAccountId string
}

//Transaction class method create Transaction
func (transaction *Transaction) create(srcAccountId, destAccountId string, amount float32) *Transaction {

	fmt.Println("creating transaction")
	transaction.srcAccountId = srcAccountId
	transaction.destAccountId = destAccountId
	transaction.amount = amount
	return transaction
}

type BranchManagerFacade struct {
	account *Account
	customer *Customer
	transaction *Transaction
}

func NewBranchManagerFacade() *BranchManagerFacade {
	return &BranchManagerFacade{ &Account{}, &Customer{}, &Transaction{}}
}

// BranchManagerFacade class method createCustomerAccount
func (facade *BranchManagerFacade) createCustomerAccount(customerName, accountType string) (*Customer, *Account) {

	var (
		customer = facade.customer.create(customerName)
		account = facade.account.create(accountType)
	)

	return customer, account
}

// BranchManagerFacade class method createTransaction
func (facade *BranchManagerFacade) createTransaction(srcAccountId, destAccountId string, amount float32) *Transaction {

	var transaction = facade.transaction.create(srcAccountId, destAccountId, amount)

	return transaction
}

func main() {
	var (
		facade = NewBranchManagerFacade()
		customer *Customer
		account *Account
		transaction = facade.createTransaction("21456", "87345", 1000)
	)

	customer, account = facade.createCustomerAccount("Thomas Smith", "Savings")
	fmt.Println(customer.name)
	fmt.Println(account.accountType)
	fmt.Println(transaction.amount)
}