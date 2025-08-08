package command

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func extractRelativeFilePath(cwd string, fullPath string) string {

	cwdArr := strings.Split(cwd, "/")

	pathArr := []string{}

	for _, item := range strings.Split(fullPath, "/") {
		if !contains(cwdArr, item) {
			pathArr = append(pathArr, item)
		}
	}

	return strings.Join(pathArr, "/")

}

func UploadAll() {

	cwd, err := os.Getwd()

	if err != nil {
		fmt.Println(err)
		return
	}

	err = filepath.Walk(cwd, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		relativeFilePath := extractRelativeFilePath(cwd, path)

		Upload(relativeFilePath)

		return nil

	})

	fmt.Println(err)

}

func Upload(filePath string) {

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writter := multipart.NewWriter(body)

	part, err := writter.CreateFormFile("file", filePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = io.Copy(part, file)
	if err != nil {
		fmt.Println(err)
		return
	}

	projectName, _ := getProjectCredentials("root_folder")

	_ = writter.WriteField("path", projectName+"/"+filePath)

	err = writter.Close()
	if err != nil {
		fmt.Println("Errpr closing file")
		return
	}

	req, _ := http.NewRequest("POST", "http://localhost:2000/upload", body)
	req.Header.Set("Content-Type", writter.FormDataContentType())

	clientToken, err := getCredentials("client_token")
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Authorization", clientToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	fmt.Println(string(respBody))

}
