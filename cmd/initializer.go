package main

import (
	"fmt"
	"github.com/buzzxu/yuanmai-platform-go/logs"
	"github.com/buzzxu/yuanmai-platform-go/types"
	"go.uber.org/fx"
	"golang-echo-starter/api/routers"
	"golang-echo-starter/internal/lib"
)

type AppContextInitializer[T any] struct {
}

func (t *AppContextInitializer[T]) Initialize(env *types.Environment[T]) (fx.Option, error) {
	return fx.Options(fx.Provide(func(env *types.Environment[T]) (logs.Logger, error) {
		return logs.NewLogger(env)
	}), modules), nil
}

func (t *AppContextInitializer[T]) Run() interface{} {
	return func(lc fx.Lifecycle, requestHandler *lib.RequestHandler, env *types.Environment[T], routes routers.Routes) {
		routes.Setup()
		err := requestHandler.Echo.Start(fmt.Sprintf(":%d", env.Server.Port))
		if err != nil {
			return
		}
	}
}
