package profile

import "github.com/google/uuid"

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
