package middlewares

import (
	"github.com/buzzxu/yuanmai-cloud-sdk-go/apis"
	"github.com/buzzxu/yuanmai-platform-go/logs"
	"github.com/buzzxu/yuanmai-platform-go/types"
	jsoniter "github.com/json-iterator/go"
	"github.com/labstack/echo/v4"
	_types "golang-echo-starter/pkg/types"
)

type JWTAuthMiddleware struct {
	service *apis.OpenSecurityApi
	logger  logs.Logger
}

// NewJWTAuthMiddleware creates new jwt auth middleware
func NewJWTAuthMiddleware(
	logger logs.Logger,
	openSecurityApi *apis.OpenSecurityApi,
) JWTAuthMiddleware {
	return JWTAuthMiddleware{
		service: openSecurityApi,
		logger:  logger,
	}
}

func (m JWTAuthMiddleware) Setup() {}

func (t *JWTAuthMiddleware) Handler() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token, err := GetToken(c, echo.HeaderAuthorization)
			if err != nil {
				return err
			}
			result, err := t.service.Authorize(&apis.AuthorizeTokenRequest{
				Token: token,
			})
			if err != nil {
				return echo.NewHTTPError(403, types.R__[string](403, err.Error()))
			}
			json, err := jsoniter.Marshal(result.User)
			if err != nil {
				return echo.NewHTTPError(500, types.R__[string](500, "网络不稳定,请稍后重试"))
			}
			user := &_types.UserInfo{}

			err = jsoniter.Unmarshal(json, user)
			if err != nil {
				return echo.NewHTTPError(500, types.R__[string](500, err.Error()))
			}
			c.Set("user", user)
			return next(c)
		}
	}
}

func GetToken(c echo.Context, name string) (string, error) {
	token := c.Request().Header.Get(name)
	if token == "" {
		token = c.QueryParam(name)
		if token == "" {
			cookie, err := c.Cookie(name)
			if err != nil {
				return "", nil
			}
			token = cookie.Value
		}
	}
	return token, nil
}
