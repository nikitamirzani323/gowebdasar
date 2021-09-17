package router

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Init() *chi.Mux {
	app := chi.NewRouter()
	app.Use(middleware.RequestID)
	app.Use(middleware.RealIP)
	app.Use(middleware.Logger)
	app.Use(middleware.Recoverer)
	app.Use(middleware.Timeout(60 * time.Second))

	app.Get("/", func(rw http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		reqID := ctx.Value(middleware.RequestIDKey).(string)

		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(reqID))
	})
	return app
}
