package main

import (
	"Glass/cmd/app"
	"Glass/cmd/utils"
	"embed"
	"io/fs"
	"net/http"

	"github.com/rs/cors"
	webview "github.com/webview/webview_go"
)

var port = "8181"

//go:embed app/dist
var staticFiles embed.FS

func init() {
	mux := http.NewServeMux()
	staticFs, _ := fs.Sub(staticFiles, "app/dist")
	utils.InitStore()

	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.FS(staticFs))))
	mux.HandleFunc("/api/main", app.GetMain)
	mux.HandleFunc("/api/main/changeFolder", app.ChangeFolder)
	mux.HandleFunc("/api/wine/get", app.DownloadWine)
	mux.HandleFunc("/api/downloaded", app.Downloaded)
	mux.HandleFunc("/api/wine/setcurrent", app.SetCurrent)
	mux.HandleFunc("/api/wine/createBottle", app.CreateBottle)
	mux.HandleFunc("/api/wine/run", app.RunWine)

	handler := cors.Default().Handler(mux)
	go http.ListenAndServe(":"+port, handler)
}

func main() {
	w := webview.New(false)
	defer w.Destroy()
	w.SetTitle("Glass")
	w.SetSize(780, 520, webview.HintNone)
	// w.Navigate("http://localhost:" + port)
	w.Navigate("http://localhost:" + "5173")
	w.Run()
}
