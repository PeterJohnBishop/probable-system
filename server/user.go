package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"probable-system/main.go/server/services"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt"
)

func CreateUser(client *dynamodb.Client, w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user services.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, `{"error": "Invalid request body"}`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	id, err := uuid.NewV1()
	if err != nil {
		http.Error(w, `{"error": "Error generating user id"}`, http.StatusInternalServerError)
		return
	}

	email := strings.ToLower(user.Email)

	userId := fmt.Sprintf("u_%s", id)

	hashedPassword, err := services.HashedPassword(user.Password)
	if err != nil {
		return
	}

	newUser := map[string]types.AttributeValue{
		"id":       &types.AttributeValueMemberS{Value: userId},
		"name":     &types.AttributeValueMemberS{Value: user.Name},
		"email":    &types.AttributeValueMemberS{Value: email},
		"password": &types.AttributeValueMemberS{Value: hashedPassword},
	}

	services.PutItem(client, "users", newUser)

	message := fmt.Sprintf(`{"message": "User created successfully", "user_id": %s}`, userId)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(message))

}

func AuthUser(client *dynamodb.Client, w http.ResponseWriter, r *http.Request) {

	type LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var req LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error": "Invalid request body"}`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	user, err := services.GetUserByEmail(client, "users", req.Email)
	if err != nil {
		http.Error(w, `{"error": "No user found with that email."}`, http.StatusInternalServerError)
	}

	pass := services.CheckPasswordHash(req.Password, user.Password)
	if !pass {
		http.Error(w, `{"error": "Password Verfication Failed"}`, http.StatusInternalServerError)
		return
	}

	userClaims := services.UserClaims{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
	}

	token, err := services.NewAccessToken(userClaims)
	if err != nil {
		http.Error(w, `{"error": "Password Verfication Failed"}`, http.StatusInternalServerError)
		return
	}

	refreshToken, err := services.NewRefreshToken(userClaims.StandardClaims)
	if err != nil {
		http.Error(w, `{"error": "Password Verfication Failed"}`, http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"message":       "Login Success",
		"token":         token,
		"refresh_token": refreshToken,
		"user":          user,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, `{"error": "Failed to encode response"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

}
