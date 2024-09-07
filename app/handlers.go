package app

import (
	"boilerplate/view/layout"

	"github.com/labstack/echo/v4"
)

// GetIndexHandler ...
func GetIndexHandler(c echo.Context) error {
	return render(c, layout.Base())
}
