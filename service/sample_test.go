package service_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/hapoon/go-api-template/entity"
	mock_repository "github.com/hapoon/go-api-template/repository/mock"
	"github.com/hapoon/go-api-template/service"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestSampleServiceListSamples(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock_repository.NewMockSampleRepository(ctrl)

	// Test patterns
	tests := map[string]struct {
		HasError bool
		Input    echo.Context
		Mock     func()
	}{
		"No error": {
			HasError: false,
			Input:    echo.New().AcquireContext(),
			Mock: func() {
				m.EXPECT().List().Return(entity.SampleEntities{
					{Id: "a0001", Name: "Alice"},
				}, nil)
			},
		},
		"Error": {
			HasError: true,
			Mock:     func() { m.EXPECT().List().Return(entity.SampleEntities{}, fmt.Errorf("aaa")) },
		},
	}

	for scenario, test := range tests {
		s := service.NewSampleService(m)
		if test.Mock != nil {
			test.Mock()
		}
		fmt.Printf("Input: %+v\n", test.Input)
		err := s.ListSamples(test.Input)
		if test.HasError {
			assert.Error(t, err, scenario)
		} else {
			assert.NoError(t, err, scenario)
		}
	}
}
