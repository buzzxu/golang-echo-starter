package routers

import (
	"github.com/buzzxu/yuanmai-platform-go/logs"
	"golang-echo-starter/api/controllers"
	"golang-echo-starter/internal/lib"
)

type AuthRoutes struct {
	logger        logs.Logger
	logController *controllers.LoginController
	hander        *lib.RequestHandler
}

// Setup user routes
func (s AuthRoutes) Setup() {
	s.hander.Echo.GET("/captcha", s.logController.Captcha)
	s.hander.Echo.POST("/login", s.logController.Login)
}

func NewAuthRoutes(handler *lib.RequestHandler, logger logs.Logger, logController *controllers.LoginController) AuthRoutes {
	return AuthRoutes{
		logger:        logger,
		logController: logController,
		hander:        handler,
	}
}
