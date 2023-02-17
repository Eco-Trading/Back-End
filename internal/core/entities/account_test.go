package entities

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAccount_must_test_the_validation_of_all_required_fields_no_state(t *testing.T) {
	assert := assert.New(t)

	account := Account{
		AccountNumber: "098876",
		Email:         "diego.morais@teste.com",
		City:          "São Paulo",
		FirstName:     "Diego",
		Address:       "Rua teste, n 8887",
		LastName:      "Morais",
		Phone:         "(19) 9 9999-9999",
		Balance:       10000.09909,
	}

	_, err := account.ValidateData()

	assert.EqualError(err, "state is not filled")
}

func TestAccount_must_test_the_validation_of_all_required_fields_no_account_number(t *testing.T) {
	assert := assert.New(t)

	account := Account{
		State:     "São paulo",
		Email:     "diego.morais@teste.com",
		City:      "São Paulo",
		FirstName: "Diego",
		Address:   "Rua teste, n 8887",
		LastName:  "Morais",
		Phone:     "(19) 9 9999-9999",
		Balance:   10000.09909,
	}

	_, err := account.ValidateData()

	assert.EqualError(err, "account number is not filled")
}

func TestAccount_must_test_the_validation_of_all_required_fields_no_email(t *testing.T) {
	assert := assert.New(t)

	account := Account{
		AccountNumber: "098876",
		State:         "São paulo",
		City:          "São Paulo",
		FirstName:     "Diego",
		Address:       "Rua teste, n 8887",
		LastName:      "Morais",
		Phone:         "(19) 9 9999-9999",
		Balance:       10000.09909,
	}

	_, err := account.ValidateData()

	assert.EqualError(err, "email is not filled")
}

func TestAccount_must_test_the_validation_of_all_required_fields_no_city(t *testing.T) {
	assert := assert.New(t)

	account := Account{
		AccountNumber: "098876",
		State:         "São paulo",
		Email:         "diego.morais@gmail.com",
		FirstName:     "Diego",
		Address:       "Rua teste, n 8887",
		LastName:      "Morais",
		Phone:         "(19) 9 9999-9999",
		Balance:       10000.09909,
	}

	_, err := account.ValidateData()

	assert.EqualError(err, "city name is not filled")
}

func TestAccount_must_test_the_validation_of_all_required_fields_no_firstName(t *testing.T) {
	assert := assert.New(t)

	account := Account{
		AccountNumber: "098876",
		Email:         "diego.morais@teste.com",
		City:          "São Paulo",
		Address:       "Rua teste, n 8887",
		LastName:      "Morais",
		Phone:         "(19) 9 9999-9999",
		State:         "São Paulo",
		Balance:       10000.09909,
	}

	_, err := account.ValidateData()

	assert.EqualError(err, "first name is not filled")
}

func TestAccount_must_test_the_validation_of_all_required_fields_no_lastName(t *testing.T) {
	assert := assert.New(t)

	account := Account{
		AccountNumber: "098876",
		Email:         "diego.morais@teste.com",
		City:          "São Paulo",
		FirstName:     "Diego",
		Address:       "Rua teste, n 8887",
		Phone:         "(19) 9 9999-9999",
		State:         "São Paulo",
		Balance:       10000.09909,
	}

	_, err := account.ValidateData()

	assert.EqualError(err, "last name is not filled")
}

func TestAccount_must_test_the_validation_of_all_required_fields_no_phone(t *testing.T) {
	assert := assert.New(t)

	account := Account{
		AccountNumber: "098876",
		Email:         "diego.morais@teste.com",
		City:          "São Paulo",
		FirstName:     "Diego",
		Address:       "Rua teste, n 8887",
		LastName:      "Morais",
		State:         "São Paulo",
		Balance:       10000.09909,
	}

	_, err := account.ValidateData()

	assert.EqualError(err, "state is not filled")
}

func TestAccount_must_test_the_validation_of_all_required_fields_no_address(t *testing.T) {
	assert := assert.New(t)

	account := Account{
		AccountNumber: "098876",
		Email:         "diego.morais@teste.com",
		City:          "São Paulo",
		FirstName:     "Diego",
		LastName:      "Morais",
		Phone:         "(19) 9 9999-9999",
		State:         "São Paulo",
		Balance:       10000.09909,
	}

	_, err := account.ValidateData()

	assert.EqualError(err, "address is not filled")
}

func TestAccount_must_test_the_validation_of_all_required_fields(t *testing.T) {
	assert := assert.New(t)

	account := Account{
		AccountNumber: "098876",
		Email:         "diego.morais@teste.com",
		City:          "São Paulo",
		FirstName:     "Diego",
		Address:       "Rua teste, n 8887",
		LastName:      "Morais",
		Phone:         "(19) 9 9999-9999",
		State:         "São Paulo",
		Balance:       10000.09909,
	}

	isOk, err := account.ValidateData()

	assert.NoError(err)
	assert.True(isOk, true)
}

func TestAccount_must_test_the_generation_account_number(t *testing.T) {
	assert := assert.New(t)

	account := Account{
		AccountNumber: "098876",
		Email:         "diego.morais@teste.com",
		City:          "São Paulo",
		FirstName:     "Diego",
		Address:       "Rua teste, n 8887",
		LastName:      "Morais",
		Phone:         "(19) 9 9999-9999",
		State:         "São Paulo",
		Balance:       10000.09909,
	}

	account.GenerateAccountNumber()
	isOk, err := account.ValidateData()

	assert.NoError(err)
	assert.True(isOk, true)
	assert.NotNil(account)
}
