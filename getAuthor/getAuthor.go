package main

import (
	"encoding/json"
	
	"serverless-second/basic"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response events.APIGatewayProxyResponse

func handler(request events.APIGatewayProxyRequest)(Response , error) {
	author_id := request.PathParameters["author_id"]
	thisItem , err := basic.Get(author_id )

	if err != nil {
		return Response{Body : "Author not found" , StatusCode: 400},nil
	}

	marshalledItem, _ := json.Marshal(thisItem)
	return Response{
		Body : string(marshalledItem) ,
		StatusCode: 200 , 
		Headers : map[string]string{
			"content-Type": "application/json" ,
		},
	} , nil 

}


func main() {
	lambda.Start(handler)
}
