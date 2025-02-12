package storage

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/google/uuid"
	"github.com/riweston/aztx/pkg/errors"
	"github.com/riweston/aztx/pkg/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFileAdapter_FetchDefaultPath(t *testing.T) {
	tests := []struct {
		name           string
		defaultFile    string
		wantErrType    error
		wantPathSuffix string
	}{
		{
			name:           "valid default file",
			defaultFile:    "/test.json",
			wantErrType:    nil,
			wantPathSuffix: "/test.json",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fa := &FileAdapter{}
			err := fa.FetchDefaultPath(tt.defaultFile)
			if tt.wantErrType != nil {
				assert.ErrorIs(t, err, tt.wantErrType)
			} else {
				assert.NoError(t, err)
				home, _ := os.UserHomeDir()
				assert.Equal(t, home+tt.wantPathSuffix, fa.Path)
			}
		})
	}
}

func TestFileAdapter_Read(t *testing.T) {
	// Create a temporary directory for test files
	tmpDir, err := os.MkdirTemp("", "file_adapter_test")
	require.NoError(t, err)
	defer os.RemoveAll(tmpDir)

	// Create a test file
	testContent := []byte(`{"test": "data"}`)
	testFile := filepath.Join(tmpDir, "test.json")
	err = os.WriteFile(testFile, testContent, 0644)
	require.NoError(t, err)

	tests := []struct {
		name    string
		path    string
		want    []byte
		wantErr error
		setup   func() error
		cleanup func()
	}{
		{
			name:    "empty path returns error",
			path:    "",
			want:    nil,
			wantErr: errors.ErrPathIsEmpty,
		},
		{
			name:    "non-existent file returns error",
			path:    filepath.Join(tmpDir, "nonexistent.json"),
			want:    nil,
			wantErr: errors.ErrFileDoesNotExist,
		},
		{
			name:    "existing file returns content",
			path:    testFile,
			want:    testContent,
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				err := tt.setup()
				require.NoError(t, err)
			}

			fa := &FileAdapter{Path: tt.path}
			got, err := fa.Read()

			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
				assert.Nil(t, got)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}

			if tt.cleanup != nil {
				tt.cleanup()
			}
		})
	}
}

func TestFileAdapter_ReadConfig(t *testing.T) {
	// Create a temporary directory for test files
	tmpDir, err := os.MkdirTemp("", "file_adapter_test")
	require.NoError(t, err)
	defer os.RemoveAll(tmpDir)

	testID := uuid.MustParse("a1a2a3a4-b1b2-c1c2-d1d2-d3d4d5d6d7d8")
	validConfig := &types.Configuration{
		InstallationID: testID,
		Tenants:        []types.Tenant{},
		Subscriptions:  []types.Subscription{},
	}
	validConfigJSON := `{"installationId":"a1a2a3a4-b1b2-c1c2-d1d2-d3d4d5d6d7d8","subscriptions":[]}`
	invalidConfigJSON := `{"installationId": invalid_json`

	// Create test files
	validFile := filepath.Join(tmpDir, "valid.json")
	err = os.WriteFile(validFile, []byte(validConfigJSON), 0644)
	require.NoError(t, err)

	invalidFile := filepath.Join(tmpDir, "invalid.json")
	err = os.WriteFile(invalidFile, []byte(invalidConfigJSON), 0644)
	require.NoError(t, err)

	tests := []struct {
		name    string
		path    string
		want    *types.Configuration
		wantErr error
	}{
		{
			name:    "empty path returns error",
			path:    "",
			want:    nil,
			wantErr: errors.ErrPathIsEmpty,
		},
		{
			name:    "non-existent file returns error",
			path:    filepath.Join(tmpDir, "nonexistent.json"),
			want:    nil,
			wantErr: errors.ErrFileDoesNotExist,
		},
		{
			name:    "invalid JSON returns error",
			path:    invalidFile,
			want:    nil,
			wantErr: errors.ErrUnmarshallingJSON(nil),
		},
		{
			name:    "valid config file returns configuration",
			path:    validFile,
			want:    validConfig,
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fa := &FileAdapter{Path: tt.path}
			got, err := fa.ReadConfig()

			if tt.wantErr != nil {
				if tt.wantErr == errors.ErrUnmarshallingJSON(nil) {
					// Special case for JSON unmarshalling errors
					assert.Error(t, err)
				} else {
					assert.ErrorIs(t, err, tt.wantErr)
				}
				assert.Nil(t, got)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
