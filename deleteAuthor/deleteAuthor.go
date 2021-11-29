package main

import (
	"bytes"
	"encoding/json"
	// "fmt"
	basic "serverless-second/basic"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response events.APIGatewayProxyResponse

type Item struct {
	Author_id string `json:"author_id"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
}


func handler(request events.APIGatewayProxyRequest)(Response , error) {

	var buf bytes.Buffer

	author_id := request.PathParameters["author_id"]
	err := basic.Delete(author_id)
	
	var message string 
	if err != nil {
		message = "Item not found "
	} else {
		message = "Item Deleted"
	}
	
	body , _ := json.Marshal(map[string]interface{}{
		"message" : message,
	})
	
	json.HTMLEscape(&buf , body)

	if err!= nil {
		return Response{Body: buf.String() , StatusCode: 400}, err 
	}else {
		return Response{Body : buf.String(), StatusCode:  200}, nil
	}

	
}


func main() {
	lambda.Start(handler)
}