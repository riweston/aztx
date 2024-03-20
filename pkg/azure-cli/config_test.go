package azure_cli

//
//import (
//	azurestate "github.com/riweston/aztx/pkg/azure-state"
//	"github.com/stretchr/testify/assert"
//	"reflect"
//	"testing"
//)
//
////var mockConfiguration = &Configuration{
////	InstallationID: uuid.UUID{},
////	Subscriptions: []Subscription{
////		{
////			ID:    uuid.MustParse("a65f3ea9-af70-4fe7-8bfa-e81a36578ddd"),
////			Name:  "foo",
////			State: "Enabled",
////			User: struct {
////				Name string `json:"name"`
////				Type string `json:"type"`
////			}{
////				Name: "Test User",
////				Type: "User",
////			},
////			IsDefault:       false,
////			TenantID:        uuid.MustParse("cca3d611-fe0c-4156-9b11-d4a58510c508"),
////			EnvironmentName: "Test Environment",
////			HomeTenantID:    uuid.MustParse("cca3d611-fe0c-4156-9b11-d4a58510c508"),
////			ManagedByTenants: []struct {
////				TenantID uuid.UUID `json:"tenantId"`
////			}{
////				{
////					TenantID: uuid.MustParse("cca3d611-fe0c-4156-9b11-d4a58510c508"),
////				},
////			},
////		},
////		{
////			ID:    uuid.MustParse("7085cdbe-2b20-4b3f-847e-88d9519e0f4c"),
////			Name:  "bar",
////			State: "Enabled",
////			User: struct {
////				Name string `json:"name"`
////				Type string `json:"type"`
////			}{
////				Name: "Test User",
////				Type: "User",
////			},
////			IsDefault:       true,
////			TenantID:        uuid.MustParse("cca3d611-fe0c-4156-9b11-d4a58510c508"),
////			EnvironmentName: "Test Environment",
////			HomeTenantID:    uuid.MustParse("cca3d611-fe0c-4156-9b11-d4a58510c508"),
////			ManagedByTenants: []struct {
////				TenantID uuid.UUID `json:"tenantId"`
////			}{
////				{
////					TenantID: uuid.MustParse("cca3d611-fe0c-4156-9b11-d4a58510c508"),
////				},
////			},
////		},
////		{
////			ID:    uuid.MustParse("a44bfa7f-6819-4fe2-92d2-0d2b730f5a5f"),
////			Name:  "baz",
////			State: "Enabled",
////			User: struct {
////				Name string `json:"name"`
////				Type string `json:"type"`
////			}{
////				Name: "Test User",
////				Type: "User",
////			},
////			IsDefault:       false,
////			TenantID:        uuid.MustParse("cca3d611-fe0c-4156-9b11-d4a58510c508"),
////			EnvironmentName: "Test Environment",
////			HomeTenantID:    uuid.MustParse("cca3d611-fe0c-4156-9b11-d4a58510c508"),
////			ManagedByTenants: []struct {
////				TenantID uuid.UUID `json:"tenantId"`
////			}{
////				{
////					TenantID: uuid.MustParse("cca3d611-fe0c-4156-9b11-d4a58510c508"),
////				},
////			},
////		},
////	},
////}
////
////type mockUserProfile struct {
////	config *Configuration
////}
////
////func (m *mockUserProfile) Fetch() error {
////	return nil
////}
////
////func (m *mockUserProfile) Read() (*Configuration, error) {
////	return mockConfiguration, nil
////}
////
////func (m *mockUserProfile) Write(cfg *Configuration) error {
////	m.config = cfg
////	return nil
////}
////
////func (m *mockUserProfile) Find(cfg *Configuration) (int, error) {
////	return 1, nil
////}
//
//func TestConfigurationAdapter_SelectWithFinder(t *testing.T) {
//	type fields struct {
//		userProfile userProfileReadWriter
//	}
//	tests := []struct {
//		name      string
//		fields    fields
//		assertion assert.ComparisonAssertionFunc
//		expected  *Subscription
//		wantErr   assert.ErrorAssertionFunc
//	}{
//		{
//			name: "SelectWithFinder",
//			fields: fields{
//				userProfile: &mockUserProfile{},
//			},
//			assertion: assert.EqualValues,
//			expected:  &mockConfiguration.Subscriptions[1],
//			wantErr:   assert.NoError,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			c := &ConfigurationAdapter{
//				userProfile: tt.fields.userProfile,
//			}
//			got, err := c.SelectWithFinder()
//			tt.assertion(t, tt.expected, got)
//			tt.wantErr(t, err)
//		})
//	}
//}
//
//func TestConfigurationAdapter_SetContext(t *testing.T) {
//	type fields struct {
//		userProfile userProfileReadWriter
//	}
//	type args struct {
//		lastContext     *azurestate.LastContext
//		selectedContext *Subscription
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//		ass
//		wantErr assert.ErrorAssertionFunc
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			c := &ConfigurationAdapter{
//				userProfile: tt.fields.userProfile,
//			}
//			if err := c.SetContext(tt.args.lastContext, tt.args.selectedContext); (err != nil) != tt.wantErr {
//				t.Errorf("SetContext() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
//
//func TestConfigurationAdapter_SetPreviousContext(t *testing.T) {
//	type fields struct {
//		userProfile userProfileReadWriter
//	}
//	type args struct {
//		lastContext *azurestate.LastContext
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			c := &ConfigurationAdapter{
//				userProfile: tt.fields.userProfile,
//			}
//			if err := c.SetPreviousContext(tt.args.lastContext); (err != nil) != tt.wantErr {
//				t.Errorf("SetPreviousContext() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
//
//func TestNewConfigurationAdapter(t *testing.T) {
//	type args struct {
//		userProfile userProfileReadWriter
//	}
//	tests := []struct {
//		name string
//		args args
//		want *ConfigurationAdapter
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := NewConfigurationAdapter(tt.args.userProfile); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("NewConfigurationAdapter() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
