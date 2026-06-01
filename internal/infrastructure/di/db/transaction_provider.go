package db

import (
	commonInterfaces "github.com/MatheusGardin/VReps/internal/domain/common/interfaces"
	"github.com/MatheusGardin/VReps/internal/infrastructure/db"

	"gorm.io/gorm"
)

func ProvideTransactionManager(gormDB *gorm.DB) commonInterfaces.TransactionManagerInterface {
	return db.NewTransactionManager(gormDB)
}
