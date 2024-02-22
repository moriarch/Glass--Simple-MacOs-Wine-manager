package app

import (
	"Glass/cmd/utils"
	"Glass/cmd/wine"
	"net/http"
)

func RunWine(w http.ResponseWriter, req *http.Request) {
	nameBottle := req.FormValue("name")
	appPath := utils.GetFile()
	wine.RunWine(nameBottle, appPath)
}
