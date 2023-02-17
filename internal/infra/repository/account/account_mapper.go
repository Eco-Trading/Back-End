package repository_account

/*
Author: Diego Morais
Data: 10/20/2023
Description: Class mapper database
*/

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type AccountMapper struct {
	ID            primitive.ObjectID `bson:"_id"`
	AccountNumber string             `bson:"account_number"`
	FirstName     string             `bson:"first_name"`
	LastName      string             `bson:"last_name"`
	Email         string             `bson:"email"`
	Address       string             `bson:"address"`
	City          string             `bson:"city"`
	State         string             `bson:"state"`
	Phone         string             `bson:"phone"`
	Balance       float64            `bson:"balance"`
	CreateAt      time.Time          `bson:"create_at"`
}
