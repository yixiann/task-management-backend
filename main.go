package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"

	_ "github.com/go-sql-driver/mysql"
	"github.com/heroku/go-getting-started/mappings"
)

var ginLambda *ginadapter.GinLambdaV2

func init() {
    mappings.CreateUrlMappings()
    ginLambda = ginadapter.NewV2(mappings.Router) 
}

func handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
    return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
    lambda.Start(handler)
}

// For local development, you can run the following command to start the server:
// func main() {
// 	mappings.CreateUrlMappings()
// 	mappings.Router.Run()
// }