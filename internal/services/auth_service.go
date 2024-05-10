package services

import (
	"github.com/buzzxu/yuanmai-cloud-sdk-go/apis"
	"github.com/buzzxu/yuanmai-platform-go/logs"
	"golang-echo-starter/internal/domains"
)

type AuthService struct {
	logger          logs.Logger
	openSecurityApi *apis.OpenSecurityApi
}

func NewAuthService(logger logs.Logger, openSecurityApi *apis.OpenSecurityApi) domains.AuthService {
	return &AuthService{
		logger:          logger,
		openSecurityApi: openSecurityApi,
	}
}

// Captcha 图形验证码
func (t *AuthService) Captcha(request *apis.CaptchaRequest) (*apis.CaptchaResponse, error) {
	return t.openSecurityApi.Captcha(request)
}

// Login 登录
func (t *AuthService) Login(request *apis.LoginRequest) (map[string]interface{}, error) {
	return t.openSecurityApi.Login(request)
}
