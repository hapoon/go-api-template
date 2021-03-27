package route

import (
	"github.com/hapoon/go-api-template/repository"
	"github.com/hapoon/go-api-template/service"

	"github.com/labstack/echo/v4"
)

func ApiV2(e *echo.Echo) {
	v2 := e.Group("/v2")
	{
		sample := service.NewSampleService(repository.NewSampleRepository())
		sampleApi := v2.Group("/sample")
		{
			sampleApi.GET("/", sample.ListSamples)
			sampleApi.GET("/:id", sample.GetSample)
			sampleApi.POST("/", sample.CreateSample)
		}
	}
}
