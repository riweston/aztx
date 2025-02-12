package profile

import (
	"bytes"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/riweston/aztx/pkg/tenant"

	"github.com/riweston/aztx/pkg/storage"
	"github.com/riweston/aztx/pkg/types"
)

type UserProfileFileAdapter struct {
	fileAdapter   *storage.FileAdapter
	configuration *types.Configuration
}

// NewUserProfileFileAdapter creates a new instance with a file adapter.
func NewUserProfileFileAdapter(path string) *UserProfileFileAdapter {
	return &UserProfileFileAdapter{
		fileAdapter: &storage.FileAdapter{Path: path},
	}
}

// Read reads the configuration from the file.
func (u *UserProfileFileAdapter) Read() (*types.Configuration, error) {
	if u.configuration != nil {
		return u.configuration, nil
	}

	data, err := u.fileAdapter.Read()
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON while handling BOM characters.
	var config types.Configuration
	err = u.unmarshalConfig(data, &config)
	if err != nil {
		return nil, err
	}

	u.configuration = &config
	return u.configuration, nil
}

// Write writes the configuration back to the file.
func (u *UserProfileFileAdapter) Write(cfg *types.Configuration) error {
	jsonData, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}

	return u.fileAdapter.Write(jsonData)
}

func (u *UserProfileFileAdapter) unmarshalConfig(data []byte, cfg *types.Configuration) error {
	// Handle BOM (Byte Order Mark)
	data = bytes.Replace(data, []byte("\uFEFF"), []byte(""), -1)
	return json.Unmarshal(data, cfg)
}

func (u *UserProfileFileAdapter) GetTenants() ([]types.Tenant, error) {
	if u.configuration == nil {
		if _, err := u.Read(); err != nil {
			return nil, err
		}
	}

	tenantManager := tenant.Manager{BaseManager: types.BaseManager{Configuration: u.configuration}}
	return tenantManager.GetTenants()
}

func (u *UserProfileFileAdapter) SaveTenantName(id uuid.UUID, customName string) error {
	if u.configuration == nil {
		if _, err := u.Read(); err != nil {
			return err
		}
	}

	tenantManager := tenant.Manager{BaseManager: types.BaseManager{Configuration: u.configuration}}
	return tenantManager.SaveTenantName(id, customName)
}
