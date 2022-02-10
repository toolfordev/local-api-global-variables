package api

import (
	"github.com/tmdgo/dependencies"
	"github.com/tmdgo/router"
)

type GlobalVariableApi struct {
	router router.Router
}

func (api *GlobalVariableApi) Init(manager *dependencies.Manager) {
	api.router.Init(manager)
	api.router.AddController(GlobalVariableController{})
	api.router.ListenAndServe()
}
