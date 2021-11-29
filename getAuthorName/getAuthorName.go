package main

import (
	"fmt"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"serverless-second/basic"
)

type Response events.APIGatewayProxyResponse

func handler(request events.APIGatewayProxyRequest)(Response , error) {
	author_id := request.PathParameters["author_id"]

	thisItem , err := basic.Get(author_id)

	if err != nil {
		return Response{Body : "Author not found" , StatusCode: 400},nil
	}

	fullName := fmt.Sprintf("%s %s" , thisItem.FirstName , thisItem.LastName)

	body , _ := json.Marshal(map[string]interface{}{
		"name" : fullName , 
	})

	return Response{Body : string(body) , StatusCode: 200} , nil 

}


func main() {
	lambda.Start(handler)

}
