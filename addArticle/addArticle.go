package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

const TABLENAME = "articles-table"
type Response events.APIGatewayProxyResponse

type Article struct {
	Article_id string `json:"article_id"`
	Shops []string `json:"shops"`
	ReleaseDate string `json:"releaseDate"`
}

func addArticle(body string)(Article , error) {
	mySession := session.Must(session.NewSession())
	svc := dynamodb.New(mySession) 

	var thisArticle Article 
	json.Unmarshal([]byte(body) , &thisArticle)

	av, err := dynamodbattribute.MarshalMap(thisArticle)
	if err != nil {
		fmt.Println("Got error marshalling map : ")
		return thisArticle , err 
	}

	input := &dynamodb.PutItemInput{
		Item : av , 
		TableName : aws.String(TABLENAME),
	}

	_ , err = svc.PutItem(input) 

	return thisArticle , err 
}


func handler(request events.APIGatewayProxyRequest) (Response , error){
	newArticle , err := addArticle(request.Body)

	if err!= nil {
		fmt.Println("Got error in post")
		fmt.Println(err.Error())
		return Response{Body: err.Error() , StatusCode: 500} , nil 
	}
	
	marshalledAritcle, _ := json.Marshal(newArticle)


	resp := Response{
		Body: string(marshalledAritcle) , 
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type" : "application/json",
		},
		
	} 
	
	return resp ,nil 
	

}

func main() {
	lambda.Start(handler)
}