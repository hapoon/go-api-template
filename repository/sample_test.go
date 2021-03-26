package repository_test

import (
	"hapoon/go-api-template/model"
	"hapoon/go-api-template/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSampleRepositoryList(t *testing.T) {
	// Test patterns
	tests := map[string]struct {
		HasError bool
		Expect   model.SampleModels
	}{
		"No error": {
			HasError: false,
			Expect: model.SampleModels{
				{Id: "a0001", Name: "Alice"},
				{Id: "a0002", Name: "Bob"},
			},
		},
	}
	for scenario, test := range tests {
		actual, err := repository.NewSampleRepository().List()
		if test.HasError {
			assert.Error(t, err, scenario)
		} else {
			assert.NoError(t, err, scenario)
			assert.Equal(t, test.Expect, actual, scenario)
		}
	}
}

func TestSampleRepositoryGetById(t *testing.T) {
	// Test patterns
	tests := map[string]struct {
		HasError bool
		Id       string
		Expect   model.SampleModel
	}{
		"No error": {
			HasError: false,
			Id:       "c0001",
			Expect:   model.SampleModel{Id: "c0001", Name: "Charles"},
		},
	}

	for scenario, test := range tests {
		actual, err := repository.NewSampleRepository().GetById(test.Id)
		if test.HasError {
			assert.Error(t, err, scenario)
		} else {
			assert.NoError(t, err, scenario)
			assert.Equal(t, test.Expect, actual, scenario)
		}
	}
}

func TestSampleRepositoryCreate(t *testing.T) {
	// Test patterns
	tests := map[string]struct {
		HasError bool
		Data     string
		Expect   *model.SampleModel
	}{
		"No error": {
			HasError: false,
			Data:     `{"id": "d0001","name": "Donald"}`,
			Expect: &model.SampleModel{
				Id:   "d0001",
				Name: "Donald",
			},
		},
		"Json decode error": {
			HasError: true,
			Data:     "aaaa",
			Expect:   &model.SampleModel{},
		},
	}

	for scenario, test := range tests {
		actual, err := repository.NewSampleRepository().Create(test.Data)
		if test.HasError {
			assert.Error(t, err, scenario)
		} else {
			assert.NoError(t, err, scenario)
			assert.Equal(t, test.Expect, actual, scenario)
		}
	}
}
