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
	c := configAddr.SetConfig()
	addr := fmt.Sprintf("%s:%d", c.Addr, c.Port)
	fmt.Println("Starting server on localhost:8080")
	http.ListenAndServe(addr, mux)
}
