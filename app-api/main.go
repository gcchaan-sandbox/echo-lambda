package main

import (
	"context"
	"log"

	"koans/app-api/app"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	echolamda "github.com/awslabs/aws-lambda-go-api-proxy/echo"
	"github.com/labstack/echo/v4"
)

var echoLambda *echolamda.EchoLambda
var e = echo.New()

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if echoLambda == nil {
		log.Printf("Echo cold start")
		echoLambda = echolamda.New(e)
	}
	app.Router(e, req)
	return echoLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
