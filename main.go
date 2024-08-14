package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	// "github.com/labstack/echo/middleware/rate"
	"fanolabs.com/fke/pkg/template"
)

func main() {
	e := echo.New()

	// Little bit of middlewares for housekeeping
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	// e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(
	// 	rate.Limit(20),
	// )))

	template.NewTemplateRenderer(e, "public/*.html")
	e.GET("/", func(e echo.Context) error {
		return e.Render(http.StatusOK, "index", nil)
	})
	e.GET("/choice", func(c echo.Context) error {
		choices := []string{"Standalone", "Clustered"}
		return c.Render(http.StatusOK, "step1", map[string]interface{}{
			"choices": choices,
		})
	})
	e.POST("/choice", func(c echo.Context) error {
		choice := new(struct {
			Choice string `json:"choice"`
		})
		if err := c.Bind(choice); err != nil {
			return err
		}

		// selectedChoice := strings.ToLower(choice.Choice)
		return c.JSON(http.StatusOK, map[string]interface{}{
			"success": true,
		})
	})

	e.Logger.Fatal(e.Start(":4040"))
}
