package utils

import (
	"os"

	"github.com/spf13/viper"
)

type WineFile struct {
	Name  string `json:"name"`
	IsDir bool   `json:"is_dir"`
}

func ScanWines() ([]WineFile, error) {
	var wineFiles []WineFile

	files, err := os.ReadDir(viper.GetString("WINE_PATH") + "/wines/")
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			wineFiles = append(wineFiles, WineFile{
				Name:  file.Name(),
				IsDir: file.IsDir(),
			})
		}
	}

	// return json.Marshal(wineFiles)
	return wineFiles, nil
}
