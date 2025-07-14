package command

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

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
