// Package types provides the core data structures and models used throughout the aztx application.
// It defines the configuration, tenant, and subscription types that represent Azure resources.
package types

import (
	"github.com/google/uuid"
	"github.com/riweston/aztx/pkg/errors"
)

// Configuration represents the root configuration structure for Azure profiles.
// It contains the installation ID, tenants, and subscriptions associated with the Azure account.
type Configuration struct {
	InstallationID uuid.UUID      `json:"installationId"`    // Unique identifier for the installation
	Tenants        []Tenant       `json:"tenants,omitempty"` // List of available Azure tenants
	Subscriptions  []Subscription `json:"subscriptions"`     // List of available Azure subscriptions
}

// Validate checks if the configuration has valid data.
// Returns an error if any validation check fails.
func (c *Configuration) Validate() error {
	if c == nil {
		return errors.ErrEmptyConfiguration
	}
	if c.InstallationID == uuid.Nil {
		return errors.ErrOperation("validating installation ID", errors.ErrEmptyConfiguration)
	}
	for _, tenant := range c.Tenants {
		if err := tenant.Validate(); err != nil {
			return errors.ErrOperation("validating tenant", err)
		}
	}
	for _, subscription := range c.Subscriptions {
		if err := subscription.Validate(); err != nil {
			return errors.ErrOperation("validating subscription", err)
		}
	}
	return nil
}

// Tenant represents an Azure tenant with its associated metadata.
// It includes both the system-assigned name and an optional custom name for better identification.
type Tenant struct {
	ID         uuid.UUID `json:"tenantId"`             // Unique identifier for the tenant
	Name       string    `json:"name"`                 // System-assigned tenant name
	CustomName string    `json:"customName,omitempty"` // User-defined custom name for the tenant
}

// GetID implements the IDGetter interface for Tenant
func (t Tenant) GetID() uuid.UUID {
	return t.ID
}

// Validate checks if the tenant has valid data.
// It ensures that required fields like ID and at least one name (either Name or CustomName) are set.
// Returns an error if any validation check fails.
func (t *Tenant) Validate() error {
	if t == nil {
		return errors.ErrEmptyConfiguration
	}
	if t.ID == uuid.Nil {
		return errors.ErrInvalidTenantID
	}
	if t.Name == "" && t.CustomName == "" {
		return errors.ErrEmptyTenantName
	}
	return nil
}

// Subscription represents an Azure subscription with its associated metadata and relationships.
// It contains detailed information about the subscription, including its state, user information,
// and tenant relationships.
type Subscription struct {
	ID    uuid.UUID `json:"id"`    // Unique identifier for the subscription
	Name  string    `json:"name"`  // Display name of the subscription
	State string    `json:"state"` // Current state of the subscription
	User  struct {
		Name string `json:"name"` // Name of the user associated with the subscription
		Type string `json:"type"` // Type of user account
	} `json:"user"`
	IsDefault        bool      `json:"isDefault"`       // Whether this is the default subscription
	TenantID         uuid.UUID `json:"tenantId"`        // ID of the tenant this subscription belongs to
	HomeTenantID     uuid.UUID `json:"homeTenantId"`    // ID of the home tenant for this subscription
	EnvironmentName  string    `json:"environmentName"` // Name of the Azure environment
	ManagedByTenants []struct {
		TenantID uuid.UUID `json:"tenantId"` // ID of the tenant managing this subscription
	} `json:"managedByTenants"`
}

// GetID implements the IDGetter interface for Subscription
func (s Subscription) GetID() uuid.UUID {
	return s.ID
}

// Validate checks if the subscription has valid data.
// It ensures that required fields like ID, Name, and TenantID are properly set.
// Returns an error if any validation check fails.
func (s *Subscription) Validate() error {
	if s == nil {
		return errors.ErrEmptyConfiguration
	}
	if s.ID == uuid.Nil {
		return errors.ErrInvalidSubscriptionID
	}
	if s.Name == "" {
		return errors.ErrOperation("validating subscription name", errors.ErrEmptyConfiguration)
	}
	if s.TenantID == uuid.Nil {
		return errors.ErrInvalidTenantID
	}
	if s.State == "" {
		return errors.ErrOperation("validating subscription state", errors.ErrEmptyConfiguration)
	}
	if s.User.Name == "" || s.User.Type == "" {
		return errors.ErrOperation("validating subscription user", errors.ErrEmptyConfiguration)
	}
	return nil
}

// User represents an Azure user
type User struct {
	Name string `json:"name"`
}
