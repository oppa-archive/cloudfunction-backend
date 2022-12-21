package db

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"log"
)

// GetDynamodbClient gets the database client connection
func GetDynamodbClient(ctx context.Context, region string) (*dynamodb.Client, error) {
	log.Printf("connecting dynamodb client in %s", region)
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	if err != nil {
		return nil, err
	}

	client := dynamodb.NewFromConfig(cfg)
	return client, nil
}
