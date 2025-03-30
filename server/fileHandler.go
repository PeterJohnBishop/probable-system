package server

import (
	"fmt"
	"net/http"

	"probable-system/main.go/server/services/fileIO"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func HandleFileUpload(client *s3.Client, w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Failed to retrieve file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	fileURL, err := fileIO.UploadFile(client, header.Filename, file)
	if err != nil {
		http.Error(w, "Failed to upload file", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("File uploaded successfully: %s", fileURL)))
}
