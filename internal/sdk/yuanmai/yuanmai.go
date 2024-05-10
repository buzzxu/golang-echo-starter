package yuanmai

import (
	"github.com/buzzxu/yuanmai-cloud-sdk-go/apis"
	"github.com/buzzxu/yuanmai-cloud-sdk-go/sdk"
	"github.com/buzzxu/yuanmai-platform-go/types"
	"go.uber.org/fx"
	_types "golang-echo-starter/pkg/types"
)

var Module = fx.Options(fx.Provide(NewYuanmaiSdkClient, apis.NewOpenSecurityApi, apis.NewOpenUserApi))

func NewYuanmaiSdkClient(env *types.Environment[_types.AppConf]) *sdk.Client {
	var credential = sdk.NewCredential(env.App.YuanmiSdk.SecretId, env.App.YuanmiSdk.SecretKey)
	return sdk.NewDefaultClient(env.App.YuanmiSdk.Domain).WithCredential(credential)
}
