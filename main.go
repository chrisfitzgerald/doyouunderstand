package main

import (
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path == "/" {
            http.ServeFile(w, r, "src/index.html")
        } else {
            if strings.HasSuffix(r.URL.Path, ".css") {
                w.Header().Set("Content-Type", "text/css")
            }
            if strings.HasPrefix(r.URL.Path, "/tailwindcss/") {
                http.ServeFile(w, r, "node_modules" + r.URL.Path)
            } else {
                http.ServeFile(w, r, "src" + r.URL.Path)
            }
        }
    })

    http.ListenAndServe(":8080", nil)
}
