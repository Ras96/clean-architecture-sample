package infrastructure

import (
	"github.com/Ras96/clean-architecture-sample/2_interface/handler"
	"github.com/labstack/echo/v4"
)

type Context struct {
	echo.Context
}

func f(next func(handler.Context) error) echo.HandlerFunc {
	return func(c echo.Context) error {
		return next(Context{c})
	}
}
