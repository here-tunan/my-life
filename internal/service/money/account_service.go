package service

import (
	"go-my-life/internal/domain/repository/moneydb"
)

func PutTransactionAccount(account *moneydb.TransactionAccount) error {
	var err error
	if account.Id == 0 {
		err = account.Insert()
	} else {
		err = account.Update()
	}
	return err
}

func AllAccounts() ([]moneydb.TransactionAccount, error) {
	return moneydb.AllAccounts()
}
