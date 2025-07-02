package provider

import (
	"github.com/itsLeonB/together-list/internal/service"
)

type Services struct {
	List *service.ListService
}

func ProvideServices(repositories *Repositories) *Services {
	return &Services{
		List: service.NewListService(repositories.Notion),
	}
}
