package accountdata

import (
	"net/http"

	"github.com/go-chi/chi"
)

func ApiRoutes() *chi.Mux {
	apiRouter := chi.NewRouter()

	apiRouter.Get("/list-accounts",
		func(w http.ResponseWriter, r *http.Request) {
			//TODO Implement
		})

	apiRouter.Route("/{accountUid:^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$}", func(r chi.Router) {
		r.Get("/details",
			func(w http.ResponseWriter, r *http.Request) {
				//TODO Implement
			})

		r.Post("/details",
			func(w http.ResponseWriter, r *http.Request) {
				//TODO Implement
			})

		r.Get("/balance-history",
			func(w http.ResponseWriter, r *http.Request) {
				//TODO Implement
			})

		r.Get("/transaction-history",
			func(w http.ResponseWriter, r *http.Request) {
				//TODO Implement
			})

		r.Post("/sync",
			func(w http.ResponseWriter, r *http.Request) {
				//TODO Implement
			})

		r.Post("/sync/{institutionConnectionUid:^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$}",
			func(w http.ResponseWriter, r *http.Request) {
				//TODO Implement
			})

		r.Get("/sync/status/{institutionConnectionUid:^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$}",
			func(w http.ResponseWriter, r *http.Request) {
				//TODO Implement
			})
	})

	return apiRouter
}
