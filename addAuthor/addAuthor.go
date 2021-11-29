package main

import (
	"fmt"
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

func handler(request events.APIGatewayProxyRequest) (Response , error){
	fmt.Println("Received body " , request.Body)
	newItem , err := basic.Post(request.Body)

	if err!= nil {
		fmt.Println("Got error in post")
		fmt.Println(err.Error())
		return Response{Body:"Error" , StatusCode: 500} , nil 
	}
	fmt.Println("Successfully added new item" , newItem)
	return Response{Body:"Success" , StatusCode: 200} , nil 
	

}

func main() {
	lambda.Start(handler)
}