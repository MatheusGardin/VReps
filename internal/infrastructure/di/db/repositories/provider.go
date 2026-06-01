package repositories

import (
	taskInterfaces "github.com/MatheusGardin/VReps/internal/domain/task/interfaces"
	uInterfaces "github.com/MatheusGardin/VReps/internal/domain/user/interfaces"
	"github.com/MatheusGardin/VReps/internal/infrastructure/db/repositories"

	"gorm.io/gorm"
)

func ProvideUserRepository(db *gorm.DB) uInterfaces.UserRepositoryInterface {
	return repositories.NewUserRepository(db)
}

func ProvideTaskRepository(db *gorm.DB) taskInterfaces.TaskRepositoryInterface {
	return repositories.NewTaskRepository(db)
}
