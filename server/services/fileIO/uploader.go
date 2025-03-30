package fileIO

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/joho/godotenv"
)

func UploadFile(client *s3.Client, filename string, fileContent multipart.File) (string, error) {

	err := godotenv.Load("server/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	bucketName := os.Getenv("AWS_BUCKET_NAME")

	_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(filename),
		Body:   fileContent, // Directly passing io.Reader
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload file: %w", err)
	}

	fileURL := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucketName, filename)
	return fileURL, nil
}
