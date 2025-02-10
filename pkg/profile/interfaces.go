package profile

import (
	"github.com/google/uuid"
	"github.com/riweston/aztx/pkg/types"
)

// StorageAdapter handles configuration storage operations
type StorageAdapter interface {
	ReadConfig() (*types.Configuration, error)
	WriteConfig(*types.Configuration) error
}

// TenantService handles tenant-related operations
type TenantService interface {
	GetTenants() ([]types.Tenant, error)
	SaveTenantName(uuid.UUID, string) error
}

// LastContextAdapter handles context persistence
type LastContextAdapter interface {
	ReadLastContextId() string
	ReadLastContextDisplayName() string
	WriteLastContext(string, string)
}

// Logger defines the logging interface
type Logger interface {
	Info(msg string, args ...interface{})
	Error(msg string, args ...interface{})
	Debug(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
}
