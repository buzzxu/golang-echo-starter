package services

import (
	"github.com/buzzxu/yuanmai-cloud-sdk-go/apis"
	"github.com/buzzxu/yuanmai-platform-go/logs"
	"golang-echo-starter/internal/domains"
)

type UserService struct {
	logger      logs.Logger
	openuserApi *apis.OpenUserApi
}

func NewUserService(logger logs.Logger, openuserApi *apis.OpenUserApi) domains.UserService {
	return &UserService{
		logger:      logger,
		openuserApi: openuserApi,
	}
}

// Info 获取用户信息
func (u *UserService) Info(request *apis.GetUserInfoRequest) (map[string]interface{}, error) {
	return u.openuserApi.Info(request)
}

// FindRoleMenu 根据角色id查询菜单
func (u *UserService) FindRoleMenu(request *apis.FindRoleMenuRequest) ([]map[string]interface{}, error) {
	return u.openuserApi.FindRoleMenu(request)
}
