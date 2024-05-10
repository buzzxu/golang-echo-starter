package services

import (
	"fmt"
	"github.com/buzzxu/yuanmai-platform-go/logs"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

// Module exports services present
var Module = fx.Options(
	fx.Provide(NewAuthService),
	fx.Provide(NewUserService),
)

type BaseService struct {
	Logger logs.Logger
	DB     *gorm.DB
}

func (s *BaseService) WithTrx(fn func(tx *gorm.DB) error) error {
	s.Logger.INFO("beginning database transaction")
	tx := s.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := fn(tx); err != nil {
		s.Logger.INFO(fmt.Sprintf("rolling back transaction due to error: %s", err.Error()))
		tx.Rollback()
		return err
	}
	s.Logger.INFO("committing transactions")
	if err := tx.Commit().Error; err != nil {
		s.Logger.ERROR("trx commit error: ", err)
		return err
	}
	return nil
}
