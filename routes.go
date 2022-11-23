package eclosion

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func (ecl *Eclosion) routes() http.Handler {

	mux := chi.NewRouter()

	// mux.Use(mux.Middlewares()...)
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)

	if ecl.Debug {
		mux.Use(middleware.Logger)
	}
	return mux

}
