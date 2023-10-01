package configs

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

var basePath string

func init() {
	basePath = getBasePath()
}

func getBasePath() string {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current working directory: %v", err)
	}
	for {
		if _, err := os.Stat(filepath.Join(cwd, "go.mod")); err == nil {
			return cwd
		}
		parent := filepath.Dir(cwd)
		if parent == cwd {
			log.Fatalf("Could not find base directory")
		}
		cwd = parent
	}
}

func ViperEnvVariable(key string) string {
	envPath := filepath.Join(basePath, ".env")
	viper.SetConfigFile(envPath)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("Viper : Invalid type assertion for %s\n", key)
	}
	return value
}
