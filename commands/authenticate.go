package command

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Authenticate() {

	var host, clientAccessKey string

	fmt.Print("Enter host url: ")
	fmt.Scanln(&host)

	fmt.Print("Enter access key: ")
	fmt.Scanln(&clientAccessKey)

	requestBody := map[string]string{
		"access_key": clientAccessKey,
	}

	jsonData, _ := json.Marshal(requestBody)

	req, err := http.NewRequest("POST", host+"/auth", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("Error response: %s\n", string(body))
		return
	}

	var response struct {
		Token string `json:"token"`
	}

	json.NewDecoder(resp.Body).Decode(&response)

	setCredential("client_token", response.Token)

	setCredential("host_url", host)

}
