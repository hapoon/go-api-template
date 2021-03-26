package repository

import (
	"encoding/json"
	"hapoon/go-api-template/entity"
)

type sampleRepository struct {
}

type SampleRepository interface {
	List() (entity.SampleEntities, error)
	GetById(id string) (entity.SampleEntity, error)
	Create(data string) (*entity.SampleEntity, error)
}

func NewSampleRepository() SampleRepository {
	return &sampleRepository{}
}

func (s sampleRepository) List() (entity.SampleEntities, error) {
	sms := entity.SampleEntities{
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

func (s sampleRepository) GetById(id string) (entity.SampleEntity, error) {
	sm := entity.SampleEntity{
		Id:   id,
		Name: "Charles",
	}
	return sm, nil
}

func (s sampleRepository) Create(data string) (*entity.SampleEntity, error) {
	sm := new(entity.SampleEntity)
	if err := json.Unmarshal([]byte(data), sm); err != nil {
		return nil, err
	}
	return sm, nil
}
