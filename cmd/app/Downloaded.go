package app

import (
	"Glass/cmd/utils"
	"encoding/json"
	"net/http"
)

func Downloaded(w http.ResponseWriter, req *http.Request) {
	files, _ := utils.ScanWines()
	js, _ := json.Marshal(files)
	w.Write(js)
}
