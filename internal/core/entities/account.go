package entities

/*
Author: Diego Morais
Data: 10/20/2023
Description: Class responsible for the business rules of the account entity
*/
import (
	"errors"
	generatorNumeric "github.com/Eco-Trading/util"
)

type Account struct {
	AccountNumber string
	FirstName     string
	LastName      string
	Address       string
	Email         string
	City          string
	State         string
	Phone         string
	Balance       float64
}

func (a *Account) GenerateAccountNumber() {
	a.AccountNumber = generatorNumeric.NewGeneratorNumeric().GenerateToString(1000000000)
}

func (a *Account) ValidateData() (bool, error) {
	if a.City == "" {
		return false, errors.New("city name is not filled")
	}
	if a.FirstName == "" {
		return false, errors.New("first name is not filled")
	}
	if a.LastName == "" {
		return false, errors.New("last name is not filled")
	}
	if a.Email == "" {
		return false, errors.New("email is not filled")
	}
	if a.Address == "" {
		return false, errors.New("address is not filled")
	}
	if a.State == "" {
		return false, errors.New("state is not filled")
	}
	if a.Phone == "" {
		return false, errors.New("state is not filled")
	}
	if a.AccountNumber == "" {
		return false, errors.New("account number is not filled")
	}
	return true, nil
}
