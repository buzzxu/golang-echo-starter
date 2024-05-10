package main

import (
	"github.com/buzzxu/yuanmai-platform-go"
	"github.com/buzzxu/yuanmai-platform-go/types"
	_types "golang-echo-starter/pkg/types"
	"time"
)

func main() {
	initializer := AppContextInitializer[_types.AppConf]{}
	yuanmai.New("env.toml", func(env *types.Environment[_types.AppConf]) error {
		return nil
	}, &initializer).RUN(5 * time.Second)
}
