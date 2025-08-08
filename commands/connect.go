package command

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func Connect() {

	fmt.Println("1. New application")
	fmt.Println("2. Existing application")
	fmt.Println("Choose 1 oo 2")

	var _option string
	fmt.Scan(&_option)

	_option = strings.TrimSpace(_option)
	option, err := strconv.Atoi(_option)

	if err != nil {
		fmt.Println("Your input is not an integer")
		return
	}

	if option > 2 || option <= 0 {
		fmt.Println("Your input is out of range")
		return
	}

	if option == 1 {
		connectNewProject()
		return
	}

	connectExistingProject()

}

func connectNewProject() {

	clientToken, err := getCredentials("client_token")
	if err != nil {
		fmt.Println(err)
		return
	}

	clientPath, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		return
	}

	requestBody := map[string]string{
		"type":        "new",
		"client_path": clientPath,
	}

	jsonData, _ := json.Marshal(requestBody)

	req, err := http.NewRequest("POST", "http://localhost:2000/connect", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", clientToken)

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
		ApplicationID string `json:"application_id"`
	}

	json.NewDecoder(resp.Body).Decode(&response)

	setProjectCredentials("application_id", response.ApplicationID)

}

func connectExistingProject() {

	clientToken, err := getCredentials("client_token")
	if err != nil {
		fmt.Println(err)
		return
	}

	req, err := http.NewRequest("GET", "http://localhost:2000/apps", nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	req.Header.Set("Authorization", clientToken)

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

	var applications []struct {
		ID         string `json:"id"`
		Name       string `json:"name"`
		ClientPath string `json:"client_path"`
		CreatedAt  string `json:"created_at"`
	}

	json.NewDecoder(resp.Body).Decode(&applications)

	fmt.Println("S/N\tNAME\tPATH\tAPPLICATION ID\tCREATED AT")
	for i, app := range applications {
		line := fmt.Sprintf("%d\t%s\t%s\t%s\t%s", i+1, app.Name, app.ClientPath, app.ID, app.CreatedAt)
		fmt.Println(line)
	}

	if len(applications) == 0 {
		fmt.Println("There are no apps to choose from, create a new one!")
		return
	}

	fmt.Printf("Choose from application 1-%d: ", len(applications))
	var _applicationSN string
	fmt.Scan(&_applicationSN)

	_applicationSN = strings.TrimSpace(_applicationSN)
	applicationSN, err := strconv.Atoi(_applicationSN)

	if err != nil {
		fmt.Println("Your input is not an integer")
		return
	}

	if applicationSN > len(applications) {
		fmt.Printf("Your selection is out of range!")
		return
	}

	selectedApplication := applications[applicationSN-1]
	setProjectCredentials("application_id", selectedApplication.ID)

	fmt.Printf("Application %s connected!", selectedApplication.Name)

}
