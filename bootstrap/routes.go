package bootstrap

import (
	"net/http"

	"github.com/gorilla/context"
	"github.com/jlope001/wasabi/app"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

// We could also put *httprouter.Router in a field to not get access to the original methods (GET, POST, etc. in uppercase)
type router struct {
	*httprouter.Router
}

func wrapHandler(h http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		context.Set(r, "params", ps)
		h.ServeHTTP(w, r)
	}
}

func NewRouter() *router {
	commonHandlers := alice.New(context.ClearHandler, loggingHandler, recoverHandler, headersHandler)

	parseEndpoints := app.ParseEndpoints{}

	router := &router{httprouter.New()}
	router.POST("/parse", wrapHandler(commonHandlers.ThenFunc(parseEndpoints.Post)))

	return router
}
