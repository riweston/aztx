package profile

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserProfileFileAdapter_Read(t *testing.T) {
	type fields struct {
		sampleConfigFilePath string
	}
	tests := []struct {
		name      string
		fields    fields
		assertion assert.ComparisonAssertionFunc
		expected  *Configuration
		wantErr   assert.ErrorAssertionFunc
	}{
		{
			name: "Reads the configuration file",
			fields: fields{
				sampleConfigFilePath: "azureProfile.json",
			},
			assertion: assert.EqualValues,
			expected: &Configuration{
				Subscriptions: []Subscription{
					{
						ID:   uuid.MustParse("9e7969ef-4cb8-4a2d-959f-bfdaae452a3d"),
						Name: "A reeeeeally long subscription name that might be truncated.",
					},
					{
						ID:   uuid.MustParse("9bb28eee-ebaa-442a-83ba-5511810fb151"),
						Name: "Test Subscription 2",
					},
					{
						ID:   uuid.MustParse("8fff24dd-2842-4dbb-8a66-1410c7bc231f"),
						Name: "Short",
					},
				},
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserProfileFileAdapter{
				path: tt.fields.sampleConfigFilePath,
			}
			got, err := u.Read()
			for i := range got.Subscriptions {
				tt.assertion(t, tt.expected.Subscriptions[i].ID, got.Subscriptions[i].ID)
				tt.assertion(t, tt.expected.Subscriptions[i].Name, got.Subscriptions[i].Name)
			}
			tt.wantErr(t, err)
		})
	}
}
