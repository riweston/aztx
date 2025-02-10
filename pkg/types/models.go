package types

import (
	"github.com/google/uuid"
	"github.com/riweston/aztx/pkg/errors"
)

type Configuration struct {
	InstallationID uuid.UUID      `json:"installationId"`
	Tenants        []Tenant       `json:"tenants,omitempty"`
	Subscriptions  []Subscription `json:"subscriptions"`
}

type Tenant struct {
	ID         uuid.UUID `json:"tenantId"`
	Name       string    `json:"name"`
	CustomName string    `json:"customName,omitempty"`
}

type Subscription struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	State string    `json:"state"`
	User  struct {
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"user"`
	IsDefault        bool      `json:"isDefault"`
	TenantID         uuid.UUID `json:"tenantId"`
	HomeTenantID     uuid.UUID `json:"homeTenantId"`
	EnvironmentName  string    `json:"environmentName"`
	ManagedByTenants []struct {
		TenantID uuid.UUID `json:"tenantId"`
	} `json:"managedByTenants"`
}

// Validate checks if the subscription has valid data
func (s *Subscription) Validate() error {
	if s.ID == uuid.Nil {
		return errors.ErrInvalidSubscriptionID
	}
	if s.Name == "" {
		return errors.ErrEmptyConfiguration
	}
	if s.TenantID == uuid.Nil {
		return errors.ErrInvalidTenantID
	}
	return nil
}

// Validate checks if the tenant has valid data
func (t *Tenant) Validate() error {
	if t.ID == uuid.Nil {
		return errors.ErrInvalidTenantID
	}
	if t.Name == "" && t.CustomName == "" {
		return errors.ErrEmptyTenantName
	}
	return nil
}
