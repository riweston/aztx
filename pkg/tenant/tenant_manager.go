package tenant

import (
	"github.com/google/uuid"
	"github.com/riweston/aztx/pkg/types"
)

type Manager struct {
	Configuration *types.Configuration
}

// GetTenants retrieves a list of unique tenants from subscriptions.
func (tm *Manager) GetTenants() ([]types.Tenant, error) {
	uniqueTenants := make(map[string]types.Tenant)

	for _, sub := range tm.Configuration.Subscriptions {
		if sub.TenantID != uuid.Nil {
			uniqueTenants[sub.TenantID.String()] = types.Tenant{
				ID:   sub.TenantID,
				Name: sub.User.Name,
			}
		}
	}

	tenants := make([]types.Tenant, 0, len(uniqueTenants))
	for _, tenant := range uniqueTenants {
		tenants = append(tenants, tenant)
	}
	return tenants, nil
}

// SaveTenantName saves or updates a tenant's custom name.
func (tm *Manager) SaveTenantName(id uuid.UUID, customName string) error {
	found := false
	for i, tenant := range tm.Configuration.Tenants {
		if tenant.ID == id {
			tm.Configuration.Tenants[i].CustomName = customName
			found = true
			break
		}
	}
	if !found {
		tm.Configuration.Tenants = append(tm.Configuration.Tenants, types.Tenant{
			ID:         id,
			CustomName: customName,
		})
	}
	return nil
}
