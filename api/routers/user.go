package routers

import (
	"github.com/buzzxu/yuanmai-platform-go/logs"
	"golang-echo-starter/api/controllers"
	"golang-echo-starter/api/middlewares"
	"golang-echo-starter/internal/lib"
)

type UserRoutes struct {
	logger         logs.Logger
	userController *controllers.UserController
	hander         *lib.RequestHandler
	jwtMiddleware  middlewares.JWTAuthMiddleware
}

func (t UserRoutes) Setup() {
	group := t.hander.Echo.Group("/user", t.jwtMiddleware.Handler())
	{
		group.GET("/info", t.userController.Info)
		group.GET("/menus", t.userController.Menus)
	}
}
func NewUserRoutes(handler *lib.RequestHandler, logger logs.Logger, userController *controllers.UserController, jwtMiddleware middlewares.JWTAuthMiddleware) UserRoutes {
	return UserRoutes{
		logger:         logger,
		userController: userController,
		hander:         handler,
		jwtMiddleware:  jwtMiddleware,
	}
}
