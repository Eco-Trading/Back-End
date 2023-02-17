package models

/*
Author: Diego Morais
Data: 10/20/2023
Description: Class that represents an account request
*/

type AccountRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	City      string `json:"city"`
	State     string `json:"state"`
	Phone     string `json:"phone"`
}
