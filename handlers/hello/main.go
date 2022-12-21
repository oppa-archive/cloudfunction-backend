package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	application "github.com/oremi-co/api-template/app"
	"github.com/oremi-co/api-template/config"
	"github.com/oremi-co/api-template/config/db"
	"github.com/oremi-co/api-template/repository"
	response "github.com/oremi-co/api-template/responses"
	"log"
)

const (
	Name = "name"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	b, err := json.Marshal(request)
	if err != nil {
		return response.EmptyResponse, err
	}
	log.Print(string(b))

	// establish database connection
	client, err := db.GetDynamodbClient(ctx, config.Region)
	if err != nil {
		panic(err)
	}

	// NewGreetingRepository set up the account repository
	repo := repository.NewGreetingRepository(client)

	// application logic
	greeting, err := application.Hello(request.PathParameters[Name], repo)
	if err != nil {
		return response.EmptyResponse, err
	}

	return response.GenerateSuccessResponse(
		response.SuccessResponse{
			Message: request.PathParameters[Name],
			Data:    greeting,
		},
	), nil

}

func main() {
	lambda.Start(handler)
}
