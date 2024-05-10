package domains

import "github.com/buzzxu/yuanmai-cloud-sdk-go/apis"

type (
	AuthService interface {
		Captcha(request *apis.CaptchaRequest) (*apis.CaptchaResponse, error)
		Login(request *apis.LoginRequest) (map[string]interface{}, error)
	}
)
