package command

import (
	"fmt"
	"os"
	"path/filepath"

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

func getProjectCredentials(key string)(value string, err error){

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

	viper.AddConfigPath(".")
	viper.SetConfigName("infracon")
	viper.SetConfigType("yaml")
	viper.Set(key, value)

	if _, err := os.Stat("infracon.yml"); os.IsNotExist(err) {
		viper.SafeWriteConfig()
		return
	}

	viper.WriteConfig()

}
