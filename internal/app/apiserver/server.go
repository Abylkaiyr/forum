package apiserver

import (
	"fmt"
	"net/http"
)

func (c *APIServer) Server() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", c.authMiddleWare(c.Home))
	mux.HandleFunc("/register", c.Register)
	mux.HandleFunc("/login", c.Login)
	mux.HandleFunc("/logout", c.Logout)
	mux.HandleFunc("/post", c.Post)
	mux.HandleFunc("/createpost", c.authMiddleWare(c.CreatePost))
	addr := fmt.Sprintf("%s:%d", "localhost", c.config.BindAddr)
	mux.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./pkg/static/styles"))))
	if err := http.ListenAndServe(addr, mux); err != nil {
		fmt.Println(err)
	}
}
