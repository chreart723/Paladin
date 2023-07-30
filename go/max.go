// 1. init go mod with Max's module
// go mod init github.com/maximilienandile/serverless-rest-api
// 2. create main.go file, in this case "max.go"
// 3. github.com/gin-gonic/gin
// go get -u github.com/gin-gonic/gin
// 4. type in gin basic code, go run it and curl localhost:8080/ping to trigger anonymous
// function(func1) and get "message": "pong"
// 5. install nodejs from github.com/nodesource/distributions, if "locked" check command ran again
// finish and test installation with curl -fsSL https://deb.nodesource.com/test | bash -
// 6. use npm install -g serverless to install serverless module via NPM(node package manager)
// xx. go get github.com/aws/aws-lambda-go
// xx. go get github.com/awslabs/aws-lambda-go-api-proxy
package main

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambda

func init() {
	r := gin.Default()
	r.GET("/ping", func (c *gin.Context)  {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello Paladin's AWS World!",
		})
	})
	ginLambda = ginadapter.New(r)
}

func HandleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// return events.APIGatewayProxyResponse{
	// 	StatusCode: 200,
	// 	Body: "Hello World",
	// }, nil
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(HandleRequest)
}
