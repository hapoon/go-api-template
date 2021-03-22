package service

import (
	"encoding/json"
	"hapoon/go-api-template/model"

	"github.com/labstack/echo/v4"
)

type SampleService interface {
	ListSamples() (model.SampleModels, error)
	GetSample(id string) (model.SampleModel, error)
	CreateSample(data string) (*model.SampleModel, error)
}

type sampleService struct {
}

func NewSampleService(e *echo.Echo) SampleService {
	return &sampleService{}
}

func (s sampleService) ListSamples() (model.SampleModels, error) {
	sms := model.SampleModels{
		model.SampleModel{
			Id:   "a0001",
			Name: "Alice",
		},
		model.SampleModel{
			Id:   "a0002",
			Name: "Bob",
		},
	}
	return sms, nil
}

func (s sampleService) GetSample(id string) (model.SampleModel, error) {
	sm := model.SampleModel{
		Id:   id,
		Name: "Charles",
	}
	return sm, nil
}

func (s sampleService) CreateSample(data string) (*model.SampleModel, error) {
	sm := new(model.SampleModel)
	if err := json.Unmarshal([]byte(data), sm); err != nil {
		return nil, err
	}
	return sm, nil
}
