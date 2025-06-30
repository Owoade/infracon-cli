package command

import (
	"crypto/rand"
	"encoding/base64"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func Init() {
	home, _ := os.UserHomeDir()
	configPath := filepath.Join(home, "config.yaml")

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		accesskey, _ := generateAccessKey(16)
		viper.Set("access_key", accesskey)
		viper.SafeWriteConfigAs(configPath)
	}

}

func generateAccessKey(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(bytes), nil
}
