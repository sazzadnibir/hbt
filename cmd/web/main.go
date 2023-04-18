package main

import (
	"fmt"
	"net/http"

	"github.com/sazzadnibir/go_modules/pkg/handlers"
)

const port = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", port))
	http.ListenAndServe(port, nil)
}
