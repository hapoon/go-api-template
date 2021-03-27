package service

import (
	"fmt"
	"net/http"

	"github.com/hapoon/go-api-template/repository"

	"github.com/labstack/echo/v4"
)

type SampleService interface {
	ListSamples(echo.Context) error
	GetSample(echo.Context) error
	CreateSample(echo.Context) error
}

type sampleService struct {
	sample repository.SampleRepository
}

func NewSampleService(sample repository.SampleRepository) SampleService {
	return &sampleService{
		sample: sample,
	}
}

func (s sampleService) ListSamples(c echo.Context) error {
	res, err := s.sample.List()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	fmt.Printf("res: %+v\n", res)
	if len(res) == 0 {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

func (s sampleService) GetSample(c echo.Context) error {
	id := c.Param("id")
	res, _ := s.sample.GetById(id)
	return c.JSON(http.StatusOK, res)
}

func (s sampleService) CreateSample(c echo.Context) error {
	var body = struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	}{}
	if err := c.Bind(&body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, body)
}
