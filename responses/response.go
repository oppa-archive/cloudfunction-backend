package responses

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
)

const (
	//StatusOK = "Ok"

	ValidationError     = "ValidationError"
	BadRequestError     = "BadRequest"
	InternalServerError = "InternalServerError"
)

var (
	EmptyResponse = events.APIGatewayProxyResponse{}
)

type Errors struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Field   string `json:"field,omitempty"`
}

type ErrorResponse struct {
	Code   int      `json:"code"` // Internal status code >=1000 | http.statusCode
	Fatal  bool     `json:"fatal,omitempty"`
	Errors []Errors `json:"errors"`
}

type SuccessResponse struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func GenerateSuccessResponse(data SuccessResponse) events.APIGatewayProxyResponse {
	rdata, _ := json.Marshal(data)

	return events.APIGatewayProxyResponse{
		IsBase64Encoded: true,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body:       string(rdata),
		StatusCode: 200,
	}
}

func GenerateErrorResponse(code int, e ErrorResponse) events.APIGatewayProxyResponse {
	rdata, _ := json.Marshal(e)

	return events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body:       string(rdata),
		StatusCode: code,
	}
}
