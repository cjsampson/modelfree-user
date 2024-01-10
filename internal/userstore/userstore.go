package userstore

import "github.com/aws/aws-sdk-go-v2/service/dynamodb"

type UserDB struct {
	Client *dynamodb.Client
}

func NewUserDB(client *dynamodb.Client) *UserDB {
	return &UserDB{client}
}

func (u *UserDB) CreateItem() {

}

func (u *UserDB) GetItem() {

}
