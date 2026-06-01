package interfaces

import (
	"context"

	uEntities "github.com/MatheusGardin/VReps/internal/domain/user/entities"
)

type JwtServiceInterface interface {
	GenerateIdentityToken(ctx context.Context, user *uEntities.User) (string, error)
}
