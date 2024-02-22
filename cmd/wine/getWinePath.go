package wine

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

func GetWinePath() string {
	winePath := viper.GetString("WINE_PATH") + "/wines/" + viper.GetString("CURRENT") + "/"
	files, _ := os.ReadDir(winePath)

	for _, file := range files {
		if file.IsDir() && strings.Contains(file.Name(), "Wine") {
			winePath = winePath + file.Name()
		}
	}
	winePath = winePath + "/Contents/Resources/wine/bin/"
	return winePath
}
