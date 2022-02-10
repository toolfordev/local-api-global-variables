package api

import (
	"net/http"
	"strconv"

	"github.com/tmdgo/router"
	"github.com/toolfordev/local-api-global-variables/application"
	"github.com/toolfordev/local-api-global-variables/models"
)

type GlobalVariableController struct {
}

func (GlobalVariableController) GetRoutes() []router.Route {
	return []router.Route{
		{Path: "/global-variable", Method: http.MethodPost, UseRequestModel: true, RequestModel: func() interface{} { return &models.GlobalVariable{} }, HandleFunc: func(service *application.GlobalVariableService, model *models.GlobalVariable) (result router.Result, err router.Error) {
			errFunc := service.Create(model)
			if errFunc != nil {
				err.Err = errFunc
				return
			}
			result.Model = model
			result.StatusCode = http.StatusCreated
			return
		}},
		{Path: "/global-variable", Method: http.MethodPut, UseRequestModel: true, RequestModel: func() interface{} { return &models.GlobalVariable{} }, HandleFunc: func(service *application.GlobalVariableService, model *models.GlobalVariable) (result router.Result, err router.Error) {
			errFunc := service.Update(model)
			if errFunc != nil {
				err.Err = errFunc
				return
			}
			result.Model = model
			result.StatusCode = http.StatusOK
			return
		}},
		{Path: "/global-variable", Method: http.MethodGet, HandleFunc: func(service *application.GlobalVariableService) (result router.Result, err router.Error) {
			models, errFunc := service.GetAll()
			if errFunc != nil {
				err.Err = errFunc
				return
			}

			result.StatusCode = http.StatusOK
			result.Model = models
			return
		}},
		{Path: "/global-variable/{id:[0-9]+}", Method: http.MethodGet, UseVars: true, HandleFunc: func(service *application.GlobalVariableService, vars *router.Vars) (result router.Result, err router.Error) {
			id, errFunc := strconv.ParseUint(vars.Value["id"], 10, 32)
			if errFunc != nil {
				err.Err = errFunc
				return
			}

			model, errFunc := service.GetByID(uint(id))
			if errFunc != nil {
				err.Err = errFunc
				return
			}

			result.StatusCode = http.StatusOK
			result.Model = model
			return
		}},
		{Path: "/global-variable/{id:[0-9]+}", Method: http.MethodDelete, UseVars: true, HandleFunc: func(service *application.GlobalVariableService, vars *router.Vars) (result router.Result, err router.Error) {
			id, errFunc := strconv.ParseUint(vars.Value["id"], 10, 32)
			if errFunc != nil {
				err.Err = errFunc
				return
			}

			errFunc = service.DeleteByID(uint(id))
			if errFunc != nil {
				err.Err = errFunc
				return
			}

			result.StatusCode = http.StatusOK
			return
		}},
	}
}
