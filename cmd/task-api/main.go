package main

import "github.com/l4rma/todo-go/pkg/routes"

func main() {
	//lambda.Start(routes.HandleRequest)
	routes.HandleRequest()
}
