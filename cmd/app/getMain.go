package app

import (
	"Glass/cmd/utils"
	"encoding/json"
	"net/http"

	"github.com/spf13/viper"
)

// func enableCors(w *http.ResponseWriter) {

// 	(*w).Header().Set("Content-Type", "text/html; charset=utf-8")
// 	(*w).Header().Set("Access-Control-Allow-Origin", "*")
// }
// func GetMain(w http.ResponseWriter, req *http.Request) {
// 	enableCors(&w)

// 	io.WriteString(w, "This is my website!\n")
// }

type Message struct {
	WINE_PATH string
	WINE      string
	CURRENT   string
	WineFiles []utils.WineFile   `json:"DOWNLOADED"`
	Bottles   []utils.BottleFile `json:"BOTTLES"`
}

func GetMain(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	wineFiles, _ := utils.ScanWines()
	bottles, _ := utils.ScanBottles()

	m := Message{
		viper.GetString("WINE_PATH"),
		viper.GetString("WINE"),
		viper.GetString("CURRENT"),
		wineFiles,
		bottles,
	}
	js, err := json.Marshal(m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(js)
}
