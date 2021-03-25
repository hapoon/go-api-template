package route

import (
	"hapoon/go-api-template/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Route(e *echo.Echo) {
	// /sample
	sample := service.NewSampleService()
	s := e.Group("/sample")
	{
		s.GET("/", func(c echo.Context) error {
			res, _ := sample.ListSamples()
			return c.JSON(http.StatusOK, res)
		})
		s.GET("/:id", func(c echo.Context) error {
			id := c.Param("id")
			res, _ := sample.GetSample(id)
			return c.JSON(http.StatusOK, res)
		})
		s.POST("/", func(c echo.Context) error {
			var body = struct {
				Id   string `json:"id"`
				Name string `json:"name"`
			}{}
			if err := c.Bind(&body); err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
			return c.JSON(http.StatusAccepted, body)
		})
	}

}
