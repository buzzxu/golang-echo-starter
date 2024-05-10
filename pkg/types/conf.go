package types

type (
	AppConf struct {
		YuanmiSdk *YuanmaiSdk `toml:"yuanmai-sdk"`
	}

	YuanmaiSdk struct {
		Domain    string `toml:"domain"`
		SecretId  string `toml:"secret-id"`
		SecretKey string `toml:"secret-key"`
	}
)
