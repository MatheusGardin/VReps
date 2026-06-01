package api

import "github.com/MatheusGardin/VReps/internal/presentation/api/routers"

type API struct {
	Router *routers.Router
}

func NewAPI(router *routers.Router) *API {
	return &API{
		Router: router,
	}
}
