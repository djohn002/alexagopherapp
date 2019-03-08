package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

// Response struct - return messsage
type Response struct {
	Message string
}

//Handlerfunc is the lambda handler
func Handlerfunc() (Response, error) {
	return Response{Message: "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software"}, nil
}

func main() {
	lambda.Start(Handlerfunc)
}
