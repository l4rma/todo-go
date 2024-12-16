package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/l4rma/todo-go/pkg/routes"
)

func main() {
	lambda.Start(routes.HandleRequest)
}
