package storage

import (
	"encoding/json"
	"io"
	"os"

	"github.com/riweston/aztx/pkg/errors"
	"github.com/riweston/aztx/pkg/types"
)

var (
	ErrFileDoesNotExist = errors.ErrFileDoesNotExist
	ErrFetchingHomePath = errors.ErrFetchingHomePath
	ErrPathNotProvided  = errors.ErrPathIsEmpty
)

// FileAdapter handles file read and write operations.
type FileAdapter struct {
	Path string
}

// FetchDefaultPath sets the path to the default file location.
func (fa *FileAdapter) FetchDefaultPath(defaultFilename string) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return errors.ErrFetchingHomePath
	}
	fa.Path = home + defaultFilename
	return nil
}

// Read reads the content of the file.
func (fa *FileAdapter) Read() ([]byte, error) {
	if fa.Path == "" {
		return nil, errors.ErrPathIsEmpty
	}
	if _, err := os.Stat(fa.Path); os.IsNotExist(err) {
		return nil, errors.ErrFileDoesNotExist
	}

	file, err := os.Open(fa.Path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return io.ReadAll(file)
}

func (fa *FileAdapter) ReadConfig() (*types.Configuration, error) {
	data, err := fa.Read()
	if err != nil {
		return nil, err
	}

	var config types.Configuration
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

// Write writes data to the file at the specified path.
func (fa *FileAdapter) Write(data []byte) error {
	if fa.Path == "" {
		return errors.ErrPathIsEmpty
	}
	return os.WriteFile(fa.Path, data, 0644)
}

func (fa *FileAdapter) WriteConfig(config *types.Configuration) error {
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}
	return fa.Write(data)
}
