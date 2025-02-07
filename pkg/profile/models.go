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
	IsDefault bool `json:"isDefault"`

	// These fields are commented out in case they are needed in the future.
	// They cause some issues with unmarshalling if they're not present in the JSON, as they're not used we'll just ignore them

	//TenantID         uuid.UUID `json:"tenantId"`
	//EnvironmentName  string    `json:"environmentName"`
	//HomeTenantID     uuid.UUID `json:"homeTenantId"`
	//ManagedByTenants []*struct {
	//	TenantID uuid.UUID `json:"tenantId"`
	//} `json:"managedByTenants"`
}
