package app

import (
	"encoding/json"
	"net/http"

	"github.com/spf13/viper"
)

func SetCurrent(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	name := req.FormValue("name")
	viper.Set("CURRENT", name)
	viper.WriteConfig()

	type Message struct {
		CURRENT string
	}
	m := Message{

		viper.GetString("CURRENT"),
	}

	js, err := json.Marshal(m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(js)
}
