package types

import (
	"testing"

	"github.com/google/uuid"
	"github.com/riweston/aztx/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestConfiguration_Validate(t *testing.T) {
	validID := uuid.MustParse("a1a2a3a4-b1b2-c1c2-d1d2-d3d4d5d6d7d8")
	validTenant := Tenant{
		ID:   validID,
		Name: "Test Tenant",
	}
	validSubscription := Subscription{
		ID:    validID,
		Name:  "Test Subscription",
		State: "Enabled",
		User: struct {
			Name string `json:"name"`
			Type string `json:"type"`
		}{
			Name: "Test User",
			Type: "User",
		},
		TenantID:        validID,
		HomeTenantID:    validID,
		EnvironmentName: "AzureCloud",
	}

	tests := []struct {
		name    string
		config  *Configuration
		wantErr error
	}{
		{
			name:    "nil configuration returns error",
			config:  nil,
			wantErr: errors.ErrEmptyConfiguration,
		},
		{
			name: "empty installation ID returns error",
			config: &Configuration{
				InstallationID: uuid.Nil,
				Tenants:        []Tenant{},
				Subscriptions:  []Subscription{},
			},
			wantErr: errors.ErrOperation("validating installation ID", errors.ErrEmptyConfiguration),
		},
		{
			name: "valid configuration returns no error",
			config: &Configuration{
				InstallationID: validID,
				Tenants:        []Tenant{validTenant},
				Subscriptions:  []Subscription{validSubscription},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.Validate()
			if tt.wantErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.wantErr.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestTenant_Validate(t *testing.T) {
	validID := uuid.MustParse("a1a2a3a4-b1b2-c1c2-d1d2-d3d4d5d6d7d8")

	tests := []struct {
		name    string
		tenant  *Tenant
		wantErr bool
	}{
		{
			name:    "nil tenant returns error",
			tenant:  nil,
			wantErr: true,
		},
		{
			name: "empty ID returns error",
			tenant: &Tenant{
				ID:   uuid.Nil,
				Name: "Test Tenant",
			},
			wantErr: true,
		},
		{
			name: "empty name returns error",
			tenant: &Tenant{
				ID:   validID,
				Name: "",
			},
			wantErr: true,
		},
		{
			name: "valid tenant returns no error",
			tenant: &Tenant{
				ID:   validID,
				Name: "Test Tenant",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.tenant.Validate()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestSubscription_Validate(t *testing.T) {
	validID := uuid.MustParse("a1a2a3a4-b1b2-c1c2-d1d2-d3d4d5d6d7d8")

	tests := []struct {
		name         string
		subscription *Subscription
		wantErr      bool
	}{
		{
			name:         "nil subscription returns error",
			subscription: nil,
			wantErr:      true,
		},
		{
			name: "empty ID returns error",
			subscription: &Subscription{
				ID:    uuid.Nil,
				Name:  "Test Subscription",
				State: "Enabled",
			},
			wantErr: true,
		},
		{
			name: "empty name returns error",
			subscription: &Subscription{
				ID:    validID,
				Name:  "",
				State: "Enabled",
			},
			wantErr: true,
		},
		{
			name: "valid subscription returns no error",
			subscription: &Subscription{
				ID:    validID,
				Name:  "Test Subscription",
				State: "Enabled",
				User: struct {
					Name string `json:"name"`
					Type string `json:"type"`
				}{
					Name: "Test User",
					Type: "User",
				},
				TenantID:        validID,
				HomeTenantID:    validID,
				EnvironmentName: "AzureCloud",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.subscription.Validate()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestIDGetter_Implementation(t *testing.T) {
	// Test that both Tenant and Subscription implement IDGetter
	validID := uuid.MustParse("a1a2a3a4-b1b2-c1c2-d1d2-d3d4d5d6d7d8")

	tenant := Tenant{ID: validID}
	subscription := Subscription{ID: validID}

	assert.Equal(t, validID, tenant.GetID())
	assert.Equal(t, validID, subscription.GetID())
}
