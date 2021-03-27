package repository_test

import (
	"testing"

	"github.com/hapoon/go-api-template/entity"
	"github.com/hapoon/go-api-template/repository"

	"github.com/stretchr/testify/assert"
)

func TestSampleRepositoryList(t *testing.T) {
	// Test patterns
	tests := map[string]struct {
		HasError bool
		Expect   entity.SampleEntities
	}{
		"No error": {
			HasError: false,
			Expect: entity.SampleEntities{
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
		Expect   entity.SampleEntity
	}{
		"No error": {
			HasError: false,
			Id:       "c0001",
			Expect:   entity.SampleEntity{Id: "c0001", Name: "Charles"},
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
		Expect   *entity.SampleEntity
	}{
		"No error": {
			HasError: false,
			Data:     `{"id": "d0001","name": "Donald"}`,
			Expect: &entity.SampleEntity{
				Id:   "d0001",
				Name: "Donald",
			},
		},
		"Json decode error": {
			HasError: true,
			Data:     "aaaa",
			Expect:   &entity.SampleEntity{},
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
