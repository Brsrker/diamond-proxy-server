package v1

import (
	"net/http"

	"github.com/go-chi/chi"

	"brsrker.com/diamond/proxyserver/pkg/clientapp"
	"brsrker.com/diamond/proxyserver/pkg/response"
)

type ClientAppRouter struct {
	Repository clientapp.Repository
}

func (cliAppRouter *ClientAppRouter) Routes() http.Handler {
	r := chi.NewRouter()
	r.Get("/", cliAppRouter.GetByClientCodeOrigin)

	return r
}

func (cliAppRouter *ClientAppRouter) GetByClientCodeOrigin(w http.ResponseWriter, r *http.Request) {

	clientCode := r.URL.Query().Get("clientCode")
	if clientCode == "" {
		response.HTTPError(w, r, http.StatusBadRequest, "'clientCode' cannot be null")
		return
	}

	appCode := r.URL.Query().Get("appCode")
	if appCode == "" {
		response.HTTPError(w, r, http.StatusBadRequest, "'appCode' cannot be null")
		return
	}

	ctx := r.Context()
	clientApp, err := cliAppRouter.Repository.GetByClientCodeOrigin(ctx, clientCode, appCode)
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, response.Map{"clientApp": clientApp})
}
