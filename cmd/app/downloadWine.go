package app

import (
	"Glass/cmd/utils"
	"Glass/cmd/wine"
	"encoding/json"
	"net/http"

	"github.com/spf13/viper"
)

func DownloadWine(w http.ResponseWriter, req *http.Request) {

	uri := req.FormValue("uri")
	name := req.FormValue("name")
	wine.GetWine(uri, name)
	files, _ := utils.ScanWines()
	viper.Set("CURRENT", name)
	viper.WriteConfig()
	js, _ := json.Marshal(files)
	w.Write(js)
}
