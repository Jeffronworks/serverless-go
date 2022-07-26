package main

import (
	"os"
	"github.com/aws/aws-skd-go/service/dynamodb"
	"github.com/aws/aws-skd-go/service/dynamodb/dynamodbiface"
	"github.com/aws-skd-go/aws"
	"github.com/aws/aws-lamda-go/lamda"
	"github.com/aws/aws-lamda-go/events"
	"github.com/aws/aws-skd-go/aws/session"
)

func main() {
	region := os.Getenv("AWS_REGION")

	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String(region)
	})

	if err != nil {
		return
	}

	dynaClient = dynamodb.New(awsSession)
	lambda.Start(handler)
}

const tableName = "LamdaInGoUser"

func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResoponse, error){
	switch req.HTTPMethod{
	case "GET":
		return handlers.GetUser(req, tableName, dynaClient)
	case "POST":
		return handlers.CreateUser(req, tableName, dynaClient)
	case "PUT":
		return handlers.UpdateUser(req, tableName, dynaClient)
	case "DELETE":
		return handlers.DeleteUser(req, tableName, dynaClient)
	}
default:
		return  handlers.unhandledMethod(req, tableName, dynaClient)
}