package utils

import (
	"os"

	"github.com/spf13/viper"
)

type BottleFile struct {
	Name  string `json:"name"`
	IsDir bool   `json:"is_dir"`
}

func ScanBottles() ([]BottleFile, error) {
	var bottles []BottleFile

	files, err := os.ReadDir(viper.GetString("WINE_PATH") + "/bottles/")
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			bottles = append(bottles, BottleFile{
				Name:  file.Name(),
				IsDir: file.IsDir(),
			})
		}
	}

	// return json.Marshal(bottles)
	return bottles, nil
}
