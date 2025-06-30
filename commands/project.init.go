package command

import (
	"os"
	"strings"
)

func InitializeProject() {

	currentWorkingDirectoryPath, _ := os.Getwd()
	currentWorkingDirectoryComponents := strings.Split(currentWorkingDirectoryPath, "/")
	currentWorkingDirectoryName := currentWorkingDirectoryComponents[len(currentWorkingDirectoryComponents)-1]

	setProjectCredentials("root_folder", currentWorkingDirectoryName)

}
