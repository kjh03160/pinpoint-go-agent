// Package ppgorilla instruments the gorilla/mux package (https://github.com/gorilla/mux).
//
// This package instruments inbound requests handled by a gorilla mux.Router.
// Register the Middleware as the middleware of the router to trace all handlers:
//
//	r := mux.NewRouter()
//	r.Use(ppgorilla.Middleware())
//
// Use WrapHandler or WrapHandlerFunc to select the handlers you want to track:
//
//	r.HandleFunc("/outgoing", ppgorilla.WrapHandlerFunc(outGoing))
package ppgorilla

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pinpoint-apm/pinpoint-go-agent"
	"github.com/pinpoint-apm/pinpoint-go-agent/plugin/http"
)

const serverName = "Gorilla/Mux HTTP Server"

// Middleware returns a mux middleware that creates a pinpoint.Tracer that instruments the http handler function.
func Middleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !pinpoint.GetAgent().Enable() {
				next.ServeHTTP(w, r)
				return
			}

			tracer := pphttp.NewHttpServerTracer(r, serverName)
			defer tracer.EndSpan()
			if !tracer.IsSampled() {
				next.ServeHTTP(w, r)
				return
			}
			defer func() {
				if e := recover(); e != nil {
					status := http.StatusInternalServerError
					pphttp.RecordHttpServerResponse(tracer, status, w.Header())
					panic(e)
				}
			}()

			tracer.NewSpanEvent("gorilla/mux.HandlerFunc()")
			defer tracer.EndSpanEvent()

			status := http.StatusOK
			w = pphttp.WrapResponseWriter(w, &status)
			r = pinpoint.RequestWithTracerContext(r, tracer)

			next.ServeHTTP(w, r)
			pphttp.RecordHttpServerResponse(tracer, status, w.Header())
		})
	}
}

// WrapHandler wraps the given http handler.
func WrapHandler(handler http.Handler) http.Handler {
	return pphttp.WrapHandler(handler, serverName)
}

// WrapHandlerFunc wraps the given http handler function.
func WrapHandlerFunc(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return pphttp.WrapHandlerFunc(f, serverName)
}
