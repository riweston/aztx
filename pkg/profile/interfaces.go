// Package profile provides interfaces and implementations for managing Azure profiles,
// including tenant and subscription management, configuration storage, and logging.
package profile

import (
	"github.com/google/uuid"
	"github.com/riweston/aztx/pkg/types"
)

// StorageAdapter defines the interface for configuration storage operations.
// Implementations of this interface handle reading and writing of Azure configuration data.
type StorageAdapter interface {
	// ReadConfig retrieves the current Azure configuration.
	// Returns a Configuration object and any error encountered during the read operation.
	ReadConfig() (*types.Configuration, error)

	// WriteConfig persists the provided Azure configuration.
	// Returns an error if the write operation fails.
	WriteConfig(*types.Configuration) error
}

// TenantService defines the interface for tenant-related operations.
// It provides functionality for managing Azure tenant information.
type TenantService interface {
	// GetTenants retrieves all available Azure tenants.
	// Returns a slice of Tenant objects and any error encountered.
	GetTenants() ([]types.Tenant, error)

	// SaveTenantName updates the display name for a tenant identified by its UUID.
	// Returns an error if the save operation fails.
	SaveTenantName(uuid.UUID, string) error
}

// LastContextAdapter defines the interface for managing the most recently used context.
// It provides functionality to persist and retrieve the last used Azure context.
type LastContextAdapter interface {
	// ReadLastContextId retrieves the ID of the last used context.
	ReadLastContextId() string

	// ReadLastContextDisplayName retrieves the display name of the last used context.
	ReadLastContextDisplayName() string

	// WriteLastContext persists both the ID and display name of the current context.
	WriteLastContext(string, string)
}

// Logger defines the interface for logging operations.
// It provides standard logging levels and formatting capabilities.
type Logger interface {
	// Info logs informational messages with optional formatting arguments.
	Info(msg string, args ...interface{})

	// Error logs error messages with optional formatting arguments.
	Error(msg string, args ...interface{})

	// Debug logs debug messages with optional formatting arguments.
	Debug(msg string, args ...interface{})

	// Warn logs warning messages with optional formatting arguments.
	Warn(msg string, args ...interface{})

	// Success logs success messages with optional formatting arguments.
	Success(msg string, args ...interface{})
}
