package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"probable-system/main.go/server/handlers"
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
	addFileIORoutes(s3Client, mux)
	addGTFSRoutes(mux)

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
		handlers.CreateUser(client, w, r)
	}))
	mux.HandleFunc("/users/login", services.LoggerMiddleware(func(w http.ResponseWriter, r *http.Request) {
		handlers.AuthUser(client, w, r)
	}))
	mux.HandleFunc("/users/all", services.LoggerMiddleware(services.VerifyJWT(func(w http.ResponseWriter, r *http.Request) {
		handlers.GetAllUsers(client, w, r)
	})))
	mux.HandleFunc("/users/id/{id}", services.LoggerMiddleware(services.VerifyJWT(func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		handlers.GetUserByID(client, w, r, id)
	})))
	mux.HandleFunc("/users/update", services.LoggerMiddleware(services.VerifyJWT(func(w http.ResponseWriter, r *http.Request) {
		handlers.UpdateUser(client, w, r)
	})))
	mux.HandleFunc("/users/delete/{id}", services.LoggerMiddleware(services.VerifyJWT(func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		handlers.DeleteUser(client, w, r, id)
	})))
}

func addChatMessageRoutes(client *dynamodb.Client, mux *http.ServeMux) {
	mux.HandleFunc("/chats/new", services.LoggerMiddleware(func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateChat(client, w, r)
	}))
	mux.HandleFunc("/chats/chat/{id}/messages/new", services.LoggerMiddleware(services.VerifyJWT(func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		handlers.CreateChatMessage(client, w, r, id)
	})))
	mux.HandleFunc("/chats/all", services.LoggerMiddleware(services.VerifyJWT(func(w http.ResponseWriter, r *http.Request) {
		handlers.GetAllChats(client, w, r)
	})))
	mux.HandleFunc("/chats/chat/{id}", services.LoggerMiddleware(services.VerifyJWT(func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		handlers.GetChatById(client, w, r, id)
	})))
	mux.HandleFunc("/chats/chat/{chatId}/messages", services.LoggerMiddleware(services.VerifyJWT(func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("chatId")
		handlers.GetChatMessages(client, w, r, id)
	})))
	mux.HandleFunc("/chats/chat/update", services.LoggerMiddleware(services.VerifyJWT(func(w http.ResponseWriter, r *http.Request) {
		handlers.UpdateChat(client, w, r)
	})))
	mux.HandleFunc("/chats/chat/{id}/delete", services.LoggerMiddleware(services.VerifyJWT(func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		handlers.DeleteChat(client, w, r, id)
	})))
	mux.HandleFunc("/chats/chat/{chatId}/messages/message/{messageId}/delete", services.LoggerMiddleware(services.VerifyJWT(func(w http.ResponseWriter, r *http.Request) {
		chatId := r.PathValue("chatId")
		messageId := r.PathValue("messageId")
		handlers.DeleteChatMessage(client, w, r, chatId, messageId)
	})))
}

func addFileIORoutes(client *s3.Client, mux *http.ServeMux) {
	mux.HandleFunc("/upload", services.LoggerMiddleware(func(w http.ResponseWriter, r *http.Request) {
		handlers.HandleFileUpload(client, w, r)
	}))
	mux.HandleFunc("/download", services.LoggerMiddleware(func(w http.ResponseWriter, r *http.Request) {
		handlers.HandleFileDownload(client, w, r)
	}))
}

func addGTFSRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/gtfs/alerts", services.LoggerMiddleware(func(w http.ResponseWriter, r *http.Request) {
		handlers.HandleGTFSRT(w, r)
	}))
	mux.HandleFunc("/gtfs/tripupdates", services.LoggerMiddleware(func(w http.ResponseWriter, r *http.Request) {
		handlers.HandleGTFSRT(w, r)
	}))
	mux.HandleFunc("/gtfs/vehiclepositions", services.LoggerMiddleware(func(w http.ResponseWriter, r *http.Request) {
		handlers.HandleGTFSRT(w, r)
	}))
}
