package middlewares

import (
	"github.com/buzzxu/yuanmai-platform-go/db"
	"github.com/buzzxu/yuanmai-platform-go/logs"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type TrxContent struct {
	echo.Context
	Tx *gorm.DB
}

type TrxMiddleware struct {
	logger logs.Logger
	db     *db.Database
}

// statusInList function checks if context writer status is in provided list
func statusInList(status int, statusList []int) bool {
	for _, i := range statusList {
		if i == status {
			return true
		}
	}
	return false
}

// NewDatabaseTrx creates new database transactions middleware
func NewDatabaseTrx(
	logger logs.Logger,
	db *db.Database,
) TrxMiddleware {
	return TrxMiddleware{
		logger: logger,
		db:     db,
	}
}

func (m TrxMiddleware) Setup() {

}

func (m TrxMiddleware) Handler() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			m.logger.INFO("setting up database transaction middleware")
			txHandle := m.db.DB.Begin()
			m.logger.INFO("beginning database transaction")
			cc := &TrxContent{c, txHandle}
			return next(cc)
		}
	}
}
