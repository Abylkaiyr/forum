package server

import (
	"fmt"
	"net/http"

	configAddr "github.com/Abylkaiyr/forum/pkg/config"
	"github.com/Abylkaiyr/forum/pkg/handlers"
)

func Server() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/register", handlers.Register)
	// mux.HandleFunc("/login", handlers.Login)
	c := configAddr.SetConfig()
	addr := fmt.Sprintf("%s:%d", c.Addr, c.Port)
	mux.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./pkg/static/styles"))))
	fmt.Println("Starting server on localhost:8080")
	http.ListenAndServe(addr, mux)
}
