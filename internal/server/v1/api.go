package v1

import (
	"net/http"

	"github.com/go-chi/chi"

	"brsrker.com/diamond/proxyserver/internal/data"
)

func New() http.Handler {
	r := chi.NewRouter()

	cliAppRouter := &ClientAppRouter{
		Repository: &data.ClientAppRepository{
			Data: data.New(),
		},
	}

	r.Mount("/clientApp", cliAppRouter.Routes())

	return r
}
