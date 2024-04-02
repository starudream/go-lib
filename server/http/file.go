package http

import (
	"io/fs"
	"net/http"
	"strings"
)

func FileServer(r Router, path string, fs fs.FS) {
	if path != "/" && path[len(path)-1] != '/' {
		r.Handle(path, http.RedirectHandler(path+"/", 301))
		path += "/"
	}
	path += "*"
	r.HandleFunc("GET "+path, func(w http.ResponseWriter, r *http.Request) {
		prefix := strings.TrimSuffix(NewContext(w, r).GetRoutePattern(), "/*")
		StripPrefix(prefix, http.FileServerFS(fs)).ServeHTTP(w, r)
	})
}
