package controllers

import (
	"github.com/buzzxu/yuanmai-cloud-sdk-go/apis"
	"github.com/buzzxu/yuanmai-platform-go/logs"
	"github.com/buzzxu/yuanmai-platform-go/types"
	"github.com/labstack/echo/v4"
	"golang-echo-starter/internal/domains"
	"net/http"
)

type LoginController struct {
	logger      logs.Logger
	authService domains.AuthService
}

func NewLoginController(logger logs.Logger, authService domains.AuthService) *LoginController {
	return &LoginController{
		logger:      logger,
		authService: authService,
	}
}

// Captcha 图形验证码
func (t *LoginController) Captcha(c echo.Context) error {
	request := &apis.CaptchaRequest{}
	if err := c.Bind(request); err != nil {
		return err
	}
	response, err := t.authService.Captcha(request)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, types.R_(response))
}

// Login 登录
func (t *LoginController) Login(c echo.Context) error {
	request := &apis.LoginRequest{}
	if err := c.Bind(request); err != nil {
		return err
	}
	response, err := t.authService.Login(request)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, types.R_(response))
}
