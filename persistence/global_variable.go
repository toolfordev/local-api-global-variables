package persistence

import (
	"github.com/toolfordev/local-api-global-variables/models"
)

type GlobalVariableEntity struct {
	ID    uint   `gorm:"primarykey"`
	Name  string `gorm:"unique"`
	Value string
}

type GlobalVariableConverter struct {
}

func (GlobalVariableConverter) ToEntity(model models.GlobalVariable) (entity GlobalVariableEntity) {
	entity.ID = model.ID
	entity.Name = model.Name
	entity.Value = model.Value
	return
}

func (GlobalVariableConverter) ToModel(entity GlobalVariableEntity) (model models.GlobalVariable) {
	model.ID = entity.ID
	model.Name = entity.Name
	model.Value = entity.Value
	return
}

func (converter GlobalVariableConverter) ToModels(entities []GlobalVariableEntity) (modelsResult []models.GlobalVariable) {
	modelsResult = make([]models.GlobalVariable, len(entities))
	for index, entity := range entities {
		model := converter.ToModel(entity)
		modelsResult[index] = model
	}
	return
}

type GlobalVariableRepository struct {
	database  *ToolForDevDatabase
	converter *GlobalVariableConverter
}

func (repository *GlobalVariableRepository) Init(database *ToolForDevDatabase) {
	repository.database = database
	repository.database.RegisterEntities(GlobalVariableEntity{})

	repository.converter = &GlobalVariableConverter{}
}

func (repository *GlobalVariableRepository) Create(model *models.GlobalVariable) (err error) {
	entity := repository.converter.ToEntity(*model)
	err = repository.database.Create(&entity)
	if err != nil {
		return
	}
	*model = repository.converter.ToModel(entity)
	return
}

func (repository *GlobalVariableRepository) Update(model *models.GlobalVariable) (err error) {
	entity := repository.converter.ToEntity(*model)
	err = repository.database.Update(&entity)
	if err != nil {
		return
	}
	*model = repository.converter.ToModel(entity)
	return
}

func (repository *GlobalVariableRepository) DeleteByID(id uint) (err error) {
	err = repository.database.DeleteByID(&GlobalVariableEntity{}, id)
	return
}

func (repository *GlobalVariableRepository) GetByID(id uint) (model models.GlobalVariable, err error) {
	entity := GlobalVariableEntity{}
	err = repository.database.SelectByID(&entity, id)
	if err != nil {
		return
	}
	model = repository.converter.ToModel(entity)
	return
}

func (repository *GlobalVariableRepository) GetAll() (models []models.GlobalVariable, err error) {
	entities := make([]GlobalVariableEntity, 0)
	err = repository.database.SelectAll(&entities)
	if err != nil {
		return
	}
	models = repository.converter.ToModels(entities)
	return
}
