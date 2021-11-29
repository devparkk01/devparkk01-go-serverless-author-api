package basic

import (
	"fmt"
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

const tableName string = "authors-table"

type Item struct {
	Author_id string `json:"author_id"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
}

func Post(body string) (Item, error) {
	// Create the dynamo client object
	mySession := session.Must(session.NewSession())
	svc := dynamodb.New(mySession) 

	// unMarshall the request body
	var thisItem Item
	json.Unmarshal([]byte(body), &thisItem)

	// Marshall the Item into a Map DynamoDB can deal with
	av, err := dynamodbattribute.MarshalMap(thisItem)
	if err != nil {
		fmt.Println("Got error marshalling map:")
		fmt.Println(err.Error())
		return thisItem, err
	}

	// Create Item in table and return
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}
	_, err = svc.PutItem(input)
	return thisItem, err

}


func Delete(author_id string)( error) {
	mySession := session.Must(session.NewSession())
	svc := dynamodb.New(mySession)
	
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"author_id" :{
				S: aws.String(author_id),
			},
		},
		TableName : aws.String(tableName), 
	}

	_ , err := svc.DeleteItem(input)

	if err!= nil {
		return err 
	}
	return nil


}

func Update(body string ) (Item , error ) {
	mySession := session.Must(session.NewSession())
	svc := dynamodb.New(mySession)

	var thisItem Item 
	json.Unmarshal([]byte(body) , &thisItem)

	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(tableName),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":f": {
				S: aws.String(thisItem.FirstName),
			},
			":l": {
				S: aws.String(thisItem.LastName),
			},
		},

		Key: map[string]*dynamodb.AttributeValue{
			"author_id": {
				S: aws.String(thisItem.Author_id),
			},
		},
		UpdateExpression: aws.String("set firstName = :f , lastName = :l "),
		ReturnValues: aws.String("UPDATED_NEW"),

	}

	_ , err := svc.UpdateItem(input)
	if err != nil {
		fmt.Println(err.Error())
	}
	return thisItem, nil 
}

func Get(author_id string )(Item , error ) {
	mySession := session.Must(session.NewSession())
	svc := dynamodb.New(mySession)
	

	output , err := svc.GetItem(&dynamodb.GetItemInput{
		TableName : aws.String(tableName), 
		Key: map[string]*dynamodb.AttributeValue{
			"author_id": {
				S: aws.String(author_id),
			},
		},
	})

	if err != nil {
		return Item{} , nil 
	}

	var thisItem Item 

	dynamodbattribute.UnmarshalMap(output.Item , &thisItem)

	return thisItem , nil 

}
