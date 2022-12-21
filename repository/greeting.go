package repository

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/oremi-co/api-template/config"
	"github.com/oremi-co/api-template/models"
)

type GreetingRepository struct {
	client    *dynamodb.Client
	tableName string
}

func NewGreetingRepository(client *dynamodb.Client) models.GreetingRepository {
	return &GreetingRepository{
		client:    client,
		tableName: config.DatabaseTable,
	}
}

func (g GreetingRepository) AddGreeting(greeting *models.Greeting) (*models.Greeting, error) {
	//TODO implement me
	panic("implement me")
}

func (g GreetingRepository) GetGreeter(name string) (*models.Greeting, error) {
	//TODO implement me
	panic("implement me")
}
