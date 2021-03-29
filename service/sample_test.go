package service_test

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
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

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Test patterns
	tests := map[string]struct {
		HasError bool
		Mock     func()
		Expect   string
	}{
		"No error": {
			HasError: false,
			Mock: func() {
				m.EXPECT().List().Return(entity.SampleEntities{
					{Id: "a0001", Name: "Alice"},
				}, nil)
			},
			Expect: `[{"id":"a0001","name":"Alice"}]`,
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
		err := s.ListSamples(c)
		if test.HasError {
			assert.Error(t, err, scenario)
		} else {
			assert.NoError(t, err, scenario)
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, test.Expect, trimLineFeed(rec.Body.String()))
		}
	}
}

func TestSampleServiceGetSample(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock_repository.NewMockSampleRepository(ctrl)

	e := echo.New()

	// Test patterns
	tests := map[string]struct {
		HasError bool
		Mock     func()
		Id       string
		Expect   string
	}{
		"No error": {
			HasError: false,
			Mock: func() {
				m.EXPECT().GetById("d0001").Return(entity.SampleEntity{Id: "d0001", Name: "Donald"}, nil)
			},
			Id:     "d0001",
			Expect: `{"id":"d0001","name":"Donald"}`,
		},
		"Error": {
			HasError: true,
			Mock: func() {
				m.EXPECT().GetById("d0002").Return(entity.SampleEntity{}, errors.New("d0002 is error"))
			},
			Id: "d0002",
		},
	}

	for scenario, test := range tests {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/sample/:id")
		c.SetParamNames("id")
		c.SetParamValues(test.Id)

		s := service.NewSampleService(m)
		if test.Mock != nil {
			test.Mock()
		}
		err := s.GetSample(c)
		if test.HasError {
			assert.Error(t, err, scenario)
		} else {
			assert.NoError(t, err, scenario)
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, test.Expect, trimLineFeed(rec.Body.String()))
		}
	}
}

func TestSampleServiceCreateSample(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock_repository.NewMockSampleRepository(ctrl)

	e := echo.New()

	// Test patterns
	tests := map[string]struct {
		HasError bool
		Mock     func()
		Body     string
		Expect   string
	}{
		"No error": {
			HasError: false,
			Body:     `{"id": "e0001","name": "Edward"}`,
			Expect:   `{"id":"e0001","name":"Edward"}`,
		},
		"Error": {
			HasError: true,
			Body:     `aaa`,
		},
	}

	for scenario, test := range tests {
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(test.Body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		s := service.NewSampleService(m)
		if test.Mock != nil {
			test.Mock()
		}
		err := s.CreateSample(c)
		if test.HasError {
			assert.Error(t, err, scenario)
		} else {
			assert.NoError(t, err, scenario)
			assert.Equal(t, http.StatusCreated, rec.Code)
			assert.Equal(t, test.Expect, trimLineFeed(rec.Body.String()))
		}
	}
}

func trimLineFeed(s string) string {
	s = strings.TrimRight(s, "\n")
	if strings.HasSuffix(s, "\r") {
		s = strings.TrimRight(s, "\r")
	}
	return s
}
