package route

import (
	"hapoon/go-api-template/repository"
	"hapoon/go-api-template/service"

	"github.com/labstack/echo/v4"
)

func ApiV1(e *echo.Echo) {
	v1 := e.Group("/v1")
	{
		sample := service.NewSampleService(repository.NewSampleRepository())
		sampleApi := v1.Group("/sample")
		{
			sampleApi.GET("/", sample.ListSamples)
			sampleApi.GET("/:id", sample.GetSample)
			sampleApi.POST("/", sample.CreateSample)
		}
	}
}
