package azure_cli

import "github.com/google/uuid"

type Configuration struct {
	InstallationID uuid.UUID      `json:"installationId"`
	Subscriptions  []Subscription `json:"subscriptions"`
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
	EnvironmentName  string    `json:"environmentName"`
	HomeTenantID     uuid.UUID `json:"homeTenantId"`
	ManagedByTenants []struct {
		TenantID uuid.UUID `json:"tenantId"`
	} `json:"managedByTenants"`
}
