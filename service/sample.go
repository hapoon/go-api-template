package service

import (
	"hapoon/go-api-template/repository"
	"net/http"

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
	sample := repository.NewSampleRepository()
	res, _ := sample.List()
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
