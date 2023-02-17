package models

/*
Author: Diego Morais
Data: 10/20/2023
Description: Class Message Request
*/

type MessageRequest struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}
