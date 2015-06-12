package bootstrap

import "net/http"

func headersHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	}

	return http.HandlerFunc(fn)
}
