package domains

import "github.com/buzzxu/yuanmai-cloud-sdk-go/apis"

type UserService interface {
	// Info 用户信息
	Info(request *apis.GetUserInfoRequest) (map[string]interface{}, error)
	// FindRoleMenu 查询角色菜单
	FindRoleMenu(request *apis.FindRoleMenuRequest) ([]map[string]interface{}, error)
}
