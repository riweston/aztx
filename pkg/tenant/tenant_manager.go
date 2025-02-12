package tenant

import (
	"fmt"

	"github.com/google/uuid"
	pkgerrors "github.com/riweston/aztx/pkg/errors"
	"github.com/riweston/aztx/pkg/finder"
	"github.com/riweston/aztx/pkg/types"
)

type Manager struct {
	types.BaseManager
}

// GetTenants retrieves a list of unique tenants from subscriptions.
func (tm *Manager) GetTenants() ([]types.Tenant, error) {
	uniqueTenants := make(map[string]types.Tenant)

	for _, sub := range tm.Configuration.Subscriptions {
		if sub.TenantID != uuid.Nil {
			tenant := types.Tenant{
				ID:   sub.TenantID,
				Name: sub.User.Name,
			}
			// Check if we have a custom name for this tenant
			for _, t := range tm.Configuration.Tenants {
				if t.ID == sub.TenantID && t.CustomName != "" {
					tenant.CustomName = t.CustomName
					break
				}
			}
			uniqueTenants[sub.TenantID.String()] = tenant
		}
	}

	if len(uniqueTenants) == 0 {
		return nil, pkgerrors.ErrTenantNotFound
	}

	tenants := make([]types.Tenant, 0, len(uniqueTenants))
	for _, tenant := range uniqueTenants {
		tenants = append(tenants, tenant)
	}
	return tenants, nil
}

// FindTenantIndex uses fuzzy finding to let user select a tenant
func (tm *Manager) FindTenantIndex() (*types.Tenant, error) {
	tenants, err := tm.GetTenants()
	if err != nil {
		return nil, fmt.Errorf("failed to get tenants: %w", err)
	}

	return finder.Fuzzy(tenants, func(t types.Tenant) string {
		if t.CustomName != "" {
			return fmt.Sprintf("%s (%s)", t.CustomName, t.ID)
		}
		return fmt.Sprintf("%s (%s)", t.Name, t.ID)
	})
}

// SaveTenantName saves or updates a tenant's custom name.
func (tm *Manager) SaveTenantName(id uuid.UUID, customName string) error {
	if id == uuid.Nil {
		return fmt.Errorf("invalid tenant ID")
	}
	if customName == "" {
		return fmt.Errorf("custom name cannot be empty")
	}

	// First verify the tenant exists
	tenants, err := tm.GetTenants()
	if err != nil {
		return fmt.Errorf("failed to verify tenant: %w", err)
	}

	_, err = finder.ByID(tenants, id)
	if err != nil {
		return pkgerrors.ErrTenantNotFound
	}

	// Update or add the custom name
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
