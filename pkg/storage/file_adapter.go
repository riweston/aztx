package storage

import (
	"bytes"
	"encoding/json"
	"io"
	"os"

	pkgerrors "github.com/riweston/aztx/pkg/errors"
	"github.com/riweston/aztx/pkg/types"
)

// FileAdapter handles file read and write operations.
type FileAdapter struct {
	Path string
}

// FetchDefaultPath sets the path to the default file location.
func (fa *FileAdapter) FetchDefaultPath(defaultFilename string) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return pkgerrors.ErrFetchingHomePath
	}
	fa.Path = home + defaultFilename
	return nil
}

// Read reads the content of the file.
func (fa *FileAdapter) Read() ([]byte, error) {
	if fa.Path == "" {
		return nil, pkgerrors.ErrPathIsEmpty
	}
	if _, err := os.Stat(fa.Path); os.IsNotExist(err) {
		return nil, pkgerrors.ErrFileDoesNotExist
	}

	file, err := os.Open(fa.Path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return io.ReadAll(file)
}

// ReadConfig reads and unmarshals configuration from file
func (fa *FileAdapter) ReadConfig() (*types.Configuration, error) {
	if fa.Path == "" {
		return nil, pkgerrors.ErrPathIsEmpty
	}

	data, err := os.ReadFile(fa.Path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, pkgerrors.ErrFileDoesNotExist
		}
		return nil, pkgerrors.ErrFileOperation("reading", err)
	}

	// Handle BOM (Byte Order Mark)
	data = bytes.TrimPrefix(data, []byte("\xef\xbb\xbf"))

	var config types.Configuration
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, pkgerrors.ErrFileOperation("unmarshaling", err)
	}
	return &config, nil
}

// Write writes data to the file at the specified path.
func (fa *FileAdapter) Write(data []byte) error {
	if fa.Path == "" {
		return pkgerrors.ErrPathIsEmpty
	}
	return os.WriteFile(fa.Path, data, 0644)
}

// WriteConfig marshals and writes configuration to file
func (fa *FileAdapter) WriteConfig(config *types.Configuration) error {
	if fa.Path == "" {
		return pkgerrors.ErrPathIsEmpty
	}

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return pkgerrors.ErrFileOperation("marshaling", err)
	}

	return os.WriteFile(fa.Path, data, 0644)
}
