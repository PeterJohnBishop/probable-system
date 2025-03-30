package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"probable-system/main.go/server/services"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func StartServer() {

	mux := http.NewServeMux()

	cfg, err := services.StartAws()
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	dynamoClient := dynamodb.NewFromConfig(cfg)

	_, err = GetTables(dynamoClient)
	if err != nil {
		log.Fatalf("unable to load dynamoDB tables, %v", err)
	}
	fmt.Printf("Connected to DynamoDB\n")

	s3Client := s3.NewFromConfig(cfg)
	_, err = s3Client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {
		log.Fatalf("unable to load S3 buckets, %v", err)
	}
	fmt.Printf("Connected to S3\n")

	services.InitAuth()
	addUserRoutes(dynamoClient, mux)
	addChatMessageRoutes(dynamoClient, mux)
	addFileUploadRoutes(s3Client, mux)

	fmt.Println("Server started on port 8080")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("unable to load dynamoDB tables, %v", err)
	}
}

func GetTables(client *dynamodb.Client) ([]string, error) {

	result, err := client.ListTables(context.TODO(), &dynamodb.ListTablesInput{})
	if err != nil {
		return nil, err
	}
	return result.TableNames, nil
}

func addUserRoutes(client *dynamodb.Client, mux *http.ServeMux) {
	mux.HandleFunc("/users/new", services.LoggerMiddleware(func(w http.ResponseWriter, r *http.Request) {
		CreateUser(client, w, r)
	}))
	mux.HandleFunc("/users/login", services.LoggerMiddleware(func(w http.ResponseWriter, r *http.Request) {
		AuthUser(client, w, r)
	}))
	mux.HandleFunc("/users/all", services.LoggerMiddleware(services.VerifyJWT(func(w http.ResponseWriter, r *http.Request) {
		GetAllUsers(client, w, r)
	})))
	mux.HandleFunc("/users/id/{id}", services.LoggerMiddleware(services.VerifyJWT(func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		GetUserByID(client, w, r, id)
	})))
	mux.HandleFunc("/users/update", services.LoggerMiddleware(services.VerifyJWT(func(w http.ResponseWriter, r *http.Request) {
		UpdateUser(client, w, r)
	})))
	mux.HandleFunc("/users/delete/{id}", services.LoggerMiddleware(services.VerifyJWT(func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		DeleteUser(client, w, r, id)
	})))
}

func addChatMessageRoutes(client *dynamodb.Client, mux *http.ServeMux) {
	mux.HandleFunc("/chats/new", services.LoggerMiddleware(func(w http.ResponseWriter, r *http.Request) {
		CreateChat(client, w, r)
	}))
	mux.HandleFunc("/chats/chat/{id}/messages/new", services.LoggerMiddleware(services.VerifyJWT(func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		CreateChatMessage(client, w, r, id)
	})))
	mux.HandleFunc("/chats/all", services.LoggerMiddleware(services.VerifyJWT(func(w http.ResponseWriter, r *http.Request) {
		GetAllChats(client, w, r)
	})))
	mux.HandleFunc("/chats/chat/{id}", services.LoggerMiddleware(services.VerifyJWT(func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		GetChatById(client, w, r, id)
	})))
	mux.HandleFunc("/chats/chat/{chatId}/messages", services.LoggerMiddleware(services.VerifyJWT(func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("chatId")
		GetChatMessages(client, w, r, id)
	})))
	mux.HandleFunc("/chats/chat/update", services.LoggerMiddleware(services.VerifyJWT(func(w http.ResponseWriter, r *http.Request) {
		UpdateChat(client, w, r)
	})))
	mux.HandleFunc("/chats/chat/{id}/delete", services.LoggerMiddleware(services.VerifyJWT(func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		DeleteChat(client, w, r, id)
	})))
	mux.HandleFunc("/chats/chat/{chatId}/messages/message/{messageId}/delete", services.LoggerMiddleware(services.VerifyJWT(func(w http.ResponseWriter, r *http.Request) {
		chatId := r.PathValue("chatId")
		messageId := r.PathValue("messageId")
		DeleteChatMessage(client, w, r, chatId, messageId)
	})))
}

func addFileUploadRoutes(client *s3.Client, mux *http.ServeMux) {
	mux.HandleFunc("/upload", services.LoggerMiddleware(func(w http.ResponseWriter, r *http.Request) {
		HandleFileUpload(client, w, r)
	}))
}
