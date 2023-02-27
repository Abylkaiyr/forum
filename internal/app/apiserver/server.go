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
	mux.HandleFunc("/post-info/", c.authMiddleWare(c.PostInfo))
	mux.HandleFunc("/post-like/", c.authMiddleWare(c.postLike))
	mux.HandleFunc("/post-dislike/", c.authMiddleWare(c.postDisLike))
	addr := fmt.Sprintf("%s:%d", "localhost", c.config.BindAddr)
	mux.Handle("/pkg/static/styles/", http.StripPrefix("/pkg/static/styles", http.FileServer(http.Dir("./pkg/static/styles"))))
	mux.Handle("/uploads/", http.StripPrefix("/uploads", http.FileServer(http.Dir("./uploads"))))

	mux.Handle("/pkg/static/files/", http.StripPrefix("/pkg/static/files", http.FileServer(http.Dir("./pkg/static/files"))))
	if err := http.ListenAndServe(addr, mux); err != nil {
		fmt.Println(err)
	}
}
