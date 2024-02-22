package wine

import (
	"Glass/cmd/utils"
	"log"
	"os"

	"github.com/spf13/viper"
)

func GetWine(uri string, name string) {

	arc, _ := utils.DownloadFromUrl(uri)
	println("Extract")
	println(arc)
	to := viper.GetString("WINE_PATH") + "/wines/" + name + "/"
	if err := os.Mkdir(to, os.ModePerm); err != nil {
		log.Fatal(err)
	}
	println(to)
	utils.Extract(arc, to)

	e := os.Remove(arc)
	if e != nil {
		log.Fatal(e)
	}
}
