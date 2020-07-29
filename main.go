package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	ID    float64 `json:"id"`
	Value string  `json:"value"`
}

type Response struct {
	Message string `json:"message"`
	Ok      bool   `json: "ok"`
}

func Handler(request Request) (Response, error) { //takes in request and outputs response
	return Response{
		Message: fmt.Sprintf("Process Request ID %f", request.ID), //format string without printing
		Ok:      true,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
