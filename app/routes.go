package app

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (app *App) SetupROutes() {
	app.Handler.Use(middleware.Logger())

	app.Handler.GET("/", GetIndexHandler)
}

// render templ component
func render(ctx echo.Context, component templ.Component) error {
	return component.Render(ctx.Request().Context(), ctx.Response())
}
