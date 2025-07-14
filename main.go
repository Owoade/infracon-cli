package main

import (
	"fmt"

	command "github.com/Owoade/infracon-cli/commands"
	"github.com/spf13/cobra"
)

func main() {

	var rootCommand = &cobra.Command{
		Use:   "infracon",
		Short: "Infracon CLI is a tool for managing infrastructure",
		Run: func(cmd *cobra.Command, args []string) {
			var action string
			if len(args) > 0 {
				action = args[0]
			} else {
				action = ""
			}

			if action == "init" {
				command.Init()
				fmt.Println("Initialized successfully")
			} else if action == "init-project" {
				command.InitializeProject()
			} else if action == "credentials"{
				command.Credentials()
			} else if action == "authenticate" {
				command.Authenticate()
			} 	else if action == "upload" {
				filePath := args[1]
				if filePath == "" {
					fmt.Println("Missing arg: filePath")
				}
				command.Upload(filePath)
			} else {
				fmt.Println("Welcome to Infracon go cli")
			}
			 

		},
	}

	rootCommand.Execute()
}
