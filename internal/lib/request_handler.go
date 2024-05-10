package lib

import (
	"github.com/buzzxu/yuanmai-platform-go/logs"
	"github.com/labstack/echo/v4"
)

type RequestHandler struct {
	Echo *echo.Echo
}

func NewRequestHandler(logger logs.Logger) *RequestHandler {
	e := echo.New()
	return &RequestHandler{
		Echo: e,
	}
}
