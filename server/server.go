package server

import (
	"fmt"
	"log"
	"net/http"
	"probable-system/main.go/server/services"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

var dynamoClient *dynamodb.Client

func newUser(w http.ResponseWriter, r *http.Request) {
	newUser := map[string]types.AttributeValue{
		"id":    &types.AttributeValueMemberS{Value: "123"},
		"name":  &types.AttributeValueMemberS{Value: "John Doe"},
		"email": &types.AttributeValueMemberS{Value: "John.Doe@email.com"},
	}

	services.PutItem(dynamoClient, "users", newUser)

}

func StartServer() {
	http.HandleFunc("/test", newUser)
	cfg, err := services.StartAws()
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	dynamoClient = dynamodb.NewFromConfig(cfg)

	tables, err := services.GetTables(dynamoClient)
	if err != nil {
		log.Fatalf("unable to load dynamoDB tables, %v", err)
	}
	fmt.Printf("Loaded %s table(s)\n", tables)
	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}
