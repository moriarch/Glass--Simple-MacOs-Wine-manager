package app

import (
	"Glass/cmd/utils"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/spf13/viper"
)

type ChangeFolderStruct struct {
	WINE_PATH string
}

func ChangeFolder(w http.ResponseWriter, req *http.Request) {
	path := utils.GetFolder()
	viper.Set("WINE_PATH", path)
	viper.WriteConfig()

	if err := os.Mkdir(path+"/wines", os.ModePerm); err != nil {
		log.Fatal(err)
	}
	if err := os.Mkdir(path+"/bottles", os.ModePerm); err != nil {
		log.Fatal(err)
	}

	m := ChangeFolderStruct{path}

	js, _ := json.Marshal(m)

	w.Write(js)
}
