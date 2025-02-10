package profile

import (
	"github.com/google/uuid"
	"github.com/riweston/aztx/pkg/types"
)

type Reader interface {
	ReadConfig() (*types.Configuration, error)
}

type Writer interface {
	WriteConfig(*types.Configuration) error
}

type TenantService interface {
	GetTenants() ([]types.Tenant, error)
	SaveTenantName(uuid.UUID, string) error
}

type SubscriptionSelector interface {
	FindSubscription(*types.Configuration) (int, error)
}
