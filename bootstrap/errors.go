package bootstrap

import (
	"log"
	"net/http"
)

const (
	ErrorMessageInternalServerError = "internal server error"
	ErrorMessageNotAllowed          = "method not allowed"
)

// error format
type WasabiErrorResult struct {
	Errors []WasabiError `json:"errors"`
}
type WasabiError struct {
	Message string `json:"message"`
	Field   string `json:"field,omitempty"`
}

func errorHandler(errors []string) string {
	return ""
}

func recoverHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %+v", err)
				http.Error(w, ErrorMessageInternalServerError, http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
