package command

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"syscall"

	"github.com/spf13/viper"
)

func Credentials() {
	accessKey, err := getCredentials("access_key")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		message := fmt.Sprintf("Acess Key: %s", accessKey)
		fmt.Println(message)
	}
}

func getCredentials(key string) (value string, err error) {
	home, _ := os.UserHomeDir()

	viper.SetConfigFile(filepath.Join(home, "config.yaml"))

	if err := viper.ReadInConfig(); err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}

	if !viper.IsSet(key) {
		return "", fmt.Errorf("value of %s is not set", key)
	}

	return viper.GetString(key), nil
}

func getProjectCredentials(key string) (value string, err error) {

	viper.AddConfigPath(".")
	viper.SetConfigName("infracon")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return "", err
	}

	if !viper.IsSet(key) {
		return "", fmt.Errorf("value of %s is not set", key)
	}

	return viper.GetString(key), nil

}

func setCredential(key, value string) error {
	home, _ := os.UserHomeDir()

	viper.SetConfigFile(filepath.Join(home, "config.yaml"))

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	viper.Set(key, value)

	if err := viper.WriteConfig(); err != nil {
		return err
	}

	return nil
}

func setProjectCredentials(key, value string) {

	fmt.Println("key", key, "value", value)

	home, _ := os.UserHomeDir()
	dirInfo, _ := os.Stat(".")
	stat := dirInfo.Sys().(*syscall.Stat_t)

	// string(unint64) won't work cos it returns the first digit to rune
	inoNoToString := strconv.FormatUint(stat.Ino, 10)
	directoryPath := filepath.Join(home, ".infracon-app-configs")
	configFilePath := filepath.Join(directoryPath, inoNoToString+".yaml")

	if _, err := os.Stat(directoryPath); os.IsNotExist(err) {
		os.MkdirAll(directoryPath, 0755)
		os.Create(configFilePath)
	}

	viper.SetConfigFile(configFilePath)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading config file: infracon.yaml", err)
	}

	viper.Set(key, value)

	if err := viper.WriteConfig(); err != nil {
		log.Fatal(err)
	}

}
