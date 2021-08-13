package main

import (
	"fmt"
	"net/http"

	"github.com/jbrainz/go-web/pkg/handlers"
)

const port = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("starting application o port %s", port))
	_ = http.ListenAndServe(port, nil)
}
