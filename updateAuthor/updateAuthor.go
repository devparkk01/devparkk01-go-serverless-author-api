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


func handler(request events.APIGatewayProxyRequest)(Response , error) {
	thisItem , err := basic.Update(request.Body) 
	if err != nil {
		return Response{
			Body: "Error",
			StatusCode: 400,
		},nil
	}

	message := fmt.Sprintf("updated author_id : %s" , thisItem.Author_id)
	return Response{
		Body: message,
		StatusCode: 200,
	},nil


}


func main() {
	lambda.Start(handler)
}