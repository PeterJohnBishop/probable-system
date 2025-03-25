package server

import (
	"fmt"
	"log"
	"net/http"
	"probable-system/main.go/server/services"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func StartServer() {

	mux := http.NewServeMux()

	cfg, err := services.StartAws()
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	dynamoClient := dynamodb.NewFromConfig(cfg)

	_, err = services.GetTables(dynamoClient)
	if err != nil {
		log.Fatalf("unable to load dynamoDB tables, %v", err)
	}
	fmt.Printf("Connected to DynamoDB\n")

	addUserRoutes(dynamoClient, mux)

	fmt.Println("Server started on port 8080")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("unable to load dynamoDB tables, %v", err)
	}
}

func addUserRoutes(client *dynamodb.Client, mux *http.ServeMux) {
	mux.HandleFunc("/users/new", func(w http.ResponseWriter, r *http.Request) {
		CreateUser(client, w, r)
	})
	mux.HandleFunc("/users/login", func(w http.ResponseWriter, r *http.Request) {
		AuthUser(client, w, r)
	})
}
