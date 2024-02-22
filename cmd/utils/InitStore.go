package utils

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func existConfig() {
	_, err := os.Stat("./app.json")
	if err != nil {
		fmt.Printf("File does not exist\n")
		os.Create("./app.json")
	}
}

func InitStore() {
	existConfig()
	viper.SetConfigName("app")
	viper.AutomaticEnv()
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		viper.Set("WINE_PATH", false)
		viper.Set("WINE", false)
		viper.Set("CURRENT", false)
		viper.WriteConfig()
	}

}
