package middleware

import "net/http"

func Static(urlpath, filepath string) http.Handler {
	return http.StripPrefix(urlpath, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")

		http.FileServer(http.Dir(filepath)).ServeHTTP(w, r)
	}))
}
