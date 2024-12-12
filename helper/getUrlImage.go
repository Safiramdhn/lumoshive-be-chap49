package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"
)

// ApiResponse struct tetap dipertahankan
type ApiResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    struct {
		FileId      string `json:"fileId"`
		Name        string `json:"name"`
		Size        int    `json:"size"`
		VersionInfo struct {
			Id   string `json:"id"`
			Name string `json:"name"`
		} `json:"versionInfo"`
		FilePath     string      `json:"filePath"`
		Url          string      `json:"url"`
		FileType     string      `json:"fileType"`
		Height       int         `json:"height"`
		Width        int         `json:"width"`
		ThumbnailUrl string      `json:"thumbnailUrl"`
		AITags       interface{} `json:"AITags"`
	} `json:"data"`
}

// Helper function untuk upload gambar dan mendapatkan respons lengkap
func UploadImage(file *multipart.FileHeader) (string, error) {
	var apiResponse ApiResponse

	f, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open file: %v", err)
	}
	defer f.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("image", filepath.Base(file.Filename))
	if err != nil {
		return "", fmt.Errorf("failed to create form file: %v", err)
	}
	_, err = io.Copy(part, f)
	if err != nil {
		return "", fmt.Errorf("failed to copy file content: %v", err)
	}

	err = writer.Close()
	if err != nil {
		return "", fmt.Errorf("failed to close writer: %v", err)
	}

	request, err := http.NewRequest("POST", "https://cdn-lumoshive-academy.vercel.app/api/v1/upload", body)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}
	request.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %v", err)
	}
	defer response.Body.Close()

	resBody, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %v", err)
	}

	err = json.Unmarshal(resBody, &apiResponse)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal response: %v", err)
	}

	if !apiResponse.Success {
		return "", fmt.Errorf("upload failed: %s", apiResponse.Message)
	}

	return apiResponse.Data.Url, nil
}
