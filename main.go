package main

import (
	"github.com/tmdgo/dependencies"
	"github.com/toolfordev/local-api-global-variables/api"
	"github.com/toolfordev/local-api-global-variables/application"
	"github.com/toolfordev/local-api-global-variables/persistence"
)

func main() {
	manager := &dependencies.Manager{}
	manager.Init()
	manager.Add(&persistence.ToolForDevDatabase{})
	manager.Add(&persistence.GlobalVariableRepository{})
	manager.Add(&application.GlobalVariableService{})

	manager.CallFunc(func(database *persistence.ToolForDevDatabase) {

	})

	api := api.GlobalVariableApi{}
	api.Init(manager)
}
