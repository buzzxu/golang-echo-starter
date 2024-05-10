package main

import (
	"go.uber.org/fx"
	"golang-echo-starter/api/controllers"
	"golang-echo-starter/api/middlewares"
	"golang-echo-starter/api/routers"
	"golang-echo-starter/internal/lib"
	"golang-echo-starter/internal/sdk/yuanmai"
	"golang-echo-starter/internal/services"
)

var modules = fx.Options(
	fx.Provide(lib.NewRequestHandler),
	middlewares.Module,
	routers.Module,
	controllers.Module,
	services.Module,
	yuanmai.Module,
)
