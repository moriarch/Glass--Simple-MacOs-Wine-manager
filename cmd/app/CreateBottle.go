package app

import (
	"Glass/cmd/utils"
	"Glass/cmd/wine"
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"

	"github.com/spf13/viper"
)

func CreateBottle(w http.ResponseWriter, req *http.Request) {
	name := req.FormValue("name")
	winePath := wine.GetWinePath() + "winecfg"
	bottlePath := viper.GetString("WINE_PATH") + "/bottles/" + name

	println(winePath)
	println(bottlePath)

	cmd := exec.Command("bash", "-c", fmt.Sprintf(`WINEPREFIX="%s" "%s"`, bottlePath, winePath))
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Print the output
	fmt.Println(string(stdout))

	bottles, _ := utils.ScanBottles()
	js, _ := json.Marshal(bottles)
	w.Write(js)
}
