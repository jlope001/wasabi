package app

import (
	"fmt"
	"net/http"
)

type ParseEndpoints struct {
}

func (p *ParseEndpoints) Post(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `{"hello":"world"}`)
	w.WriteHeader(http.StatusOK)
}
