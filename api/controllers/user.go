package controllers

import (
	"github.com/buzzxu/yuanmai-cloud-sdk-go/apis"
	"github.com/buzzxu/yuanmai-platform-go/logs"
	"github.com/buzzxu/yuanmai-platform-go/types"
	"github.com/labstack/echo/v4"
	"golang-echo-starter/internal/domains"
	_types "golang-echo-starter/pkg/types"
	"net/http"
)

type UserController struct {
	logger      logs.Logger
	userService domains.UserService
}

func NewUserController(logger logs.Logger, userService domains.UserService) *UserController {
	return &UserController{
		logger:      logger,
		userService: userService,
	}
}

func (t *UserController) Info(c echo.Context) error {
	user := c.Get("user").(*_types.UserInfo)
	request := &apis.GetUserInfoRequest{
		UserId: int64(user.ID),
	}
	if err := c.Bind(request); err != nil {
		return err
	}
	response, err := t.userService.Info(request)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, types.R_(response))
}

func (t *UserController) Menus(c echo.Context) error {
	user := c.Get("user").(*_types.UserInfo)
	request := &apis.FindRoleMenuRequest{
		RoleId: int(user.RoleId),
		Region: "formflow",
	}
	response, err := t.userService.FindRoleMenu(request)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, types.R_(response))
}
