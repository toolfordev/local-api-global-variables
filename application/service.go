package application

import (
	"github.com/toolfordev/local-api-global-variables/models"
	"github.com/toolfordev/local-api-global-variables/persistence"
)

type GlobalVariableService struct {
	repository *persistence.GlobalVariableRepository
}

func (service *GlobalVariableService) Init(repository *persistence.GlobalVariableRepository) {
	service.repository = repository
}

func (service *GlobalVariableService) Create(model *models.GlobalVariable) (err error) {
	err = service.repository.Create(model)
	return
}

func (service *GlobalVariableService) Update(model *models.GlobalVariable) (err error) {
	err = service.repository.Update(model)
	return
}

func (service *GlobalVariableService) GetByID(id uint) (model models.GlobalVariable, err error) {
	model, err = service.repository.GetByID(id)
	return
}

func (service *GlobalVariableService) GetAll() (models []models.GlobalVariable, err error) {
	models, err = service.repository.GetAll()
	return
}

func (service *GlobalVariableService) DeleteByID(id uint) (err error) {
	err = service.repository.DeleteByID(id)
	return
}
