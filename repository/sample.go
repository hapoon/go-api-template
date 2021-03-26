package repository

import (
	"encoding/json"
	"hapoon/go-api-template/model"
)

type sampleRepository struct {
}

type SampleRepository interface {
	List() (model.SampleModels, error)
	GetById(id string) (model.SampleModel, error)
	Create(data string) (*model.SampleModel, error)
}

func NewSampleRepository() SampleRepository {
	return &sampleRepository{}
}

func (s sampleRepository) List() (model.SampleModels, error) {
	sms := model.SampleModels{
		{
			Id:   "a0001",
			Name: "Alice",
		},
		{
			Id:   "a0002",
			Name: "Bob",
		},
	}
	return sms, nil
}

func (s sampleRepository) GetById(id string) (model.SampleModel, error) {
	sm := model.SampleModel{
		Id:   id,
		Name: "Charles",
	}
	return sm, nil
}

func (s sampleRepository) Create(data string) (*model.SampleModel, error) {
	sm := new(model.SampleModel)
	if err := json.Unmarshal([]byte(data), sm); err != nil {
		return nil, err
	}
	return sm, nil
}
