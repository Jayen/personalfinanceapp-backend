package institutionconnection

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func ApiRoutes() *chi.Mux {
	apiRouter := chi.NewRouter()

	apiRouter.Get("/country/{countryCode}", func(w http.ResponseWriter, r *http.Request) {
		//TODO Implement
	})

	apiRouter.Route("/connection", func(r chi.Router) {

		r.Post("/request",
			func(w http.ResponseWriter, r *http.Request) {
				//TODO Implement
			})

		r.Get("/{institutionConnectionUid:^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$}",
			func(w http.ResponseWriter, r *http.Request) {
				//TODO Implement
			})

		r.Put("/{institutionConnectionUid:^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$}/established",
			func(w http.ResponseWriter, r *http.Request) {
				//TODO Implement
			})

		r.Post("/{institutionConnectionUid:^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$}/renew",
			func(w http.ResponseWriter, r *http.Request) {
				//TODO Implement
			})

	})

	return apiRouter
}
