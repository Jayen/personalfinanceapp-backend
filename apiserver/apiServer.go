package apiserver

import (
	"net/http"
	"personal-finance-app/backend/apiserver/accountdata"
	"personal-finance-app/backend/apiserver/institutionconnection"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Start() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	//TODO add middleware for validating the user / session

	r.Mount("/api/institutions/", institutionconnection.ApiRoutes())
	r.Mount("/api/account-data/", accountdata.ApiRoutes())

	http.ListenAndServe(":3000", r)
}
