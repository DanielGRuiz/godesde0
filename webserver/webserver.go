package webserver

import (
	"fmt"
	"net/http"
)

func MiWebServer() {
	fmt.Println("Servidor en ejecucción...")
	http.HandleFunc("/", home)
	http.ListenAndServe(":3000", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	var pagina string = "./webserver/index.html"

	http.ServeFile(w, r, pagina)
}
