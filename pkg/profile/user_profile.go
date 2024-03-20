package profile

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ktr0731/go-fuzzyfinder"
	"io"
	"os"
)

type userProfileReadWriter interface {
	Fetch() error
	Read() (*Configuration, error)
	Write(*Configuration) error
	Find(*Configuration) (int, error)
}

type UserProfileFileAdapter struct {
	path          string
	configuration *Configuration
}

func (u *UserProfileFileAdapter) Fetch() error {
	if u.path != "" {
		return nil
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return ErrGettingHomeDirectory(err)
	}
	defaultPath := home + "/.azure/azureProfile.json"
	if _, err := os.Stat(defaultPath); os.IsNotExist(err) {
		return ErrFileDoesNotExist(err)
	}
	u.path = defaultPath
	return nil
}

func (u *UserProfileFileAdapter) Write(cfg *Configuration) error {
	if u.path == "" {
		return ErrPathIsEmpty
	}
	jsonData, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return ErrMarshallingJSON(err)
	}
	err = os.WriteFile(u.path, jsonData, 0644)
	if err != nil {
		return ErrWritingFile(err)
	}
	return nil
}

func (u *UserProfileFileAdapter) Read() (*Configuration, error) {
	if u.configuration != nil {
		return u.configuration, nil
	}
	if err := u.Fetch(); err != nil {
		return nil, ErrFetchingUserProfile(err)
	}
	file, err := u.openConfigFile()
	if err != nil {
		return nil, ErrReadingFile(err)
	}
	var d Configuration
	err = u.unmarshalConfig(file, &d)
	if err != nil {
		return nil, ErrUnmarshallingJSON(err)
	}
	u.configuration = &d
	return u.configuration, nil
}

func (u *UserProfileFileAdapter) openConfigFile() ([]byte, error) {
	if u.path == "" {
		return nil, ErrPathIsEmpty
	}
	file, err := os.Open(u.path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	b, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (u *UserProfileFileAdapter) unmarshalConfig(data []byte, d *Configuration) error {
	// handle zero width space character
	data = bytes.Replace(data, []byte("\uFEFF"), []byte(""), -1)
	err := json.Unmarshal(data, &d)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserProfileFileAdapter) Find(cfg *Configuration) (int, error) {
	columnWidth := u.longestDisplayNameCharacterWidth()
	format := fmt.Sprintf("%%-%ds %%s", columnWidth)
	idx, err := fuzzyfinder.Find(
		cfg.Subscriptions,
		func(i int) string {
			if cfg.Subscriptions[i].IsDefault {
				currentContext := fmt.Sprintf(format, cfg.Subscriptions[i].Name, cfg.Subscriptions[i].ID)
				return currentContext
			}
			return fmt.Sprintf(format, cfg.Subscriptions[i].Name, cfg.Subscriptions[i].ID)
		},
	)
	if err != nil {
		return 0, err
	}
	return idx, nil
}

func (u *UserProfileFileAdapter) longestDisplayNameCharacterWidth() int {
	var max int
	for _, sub := range u.configuration.Subscriptions {
		if len(sub.Name) > max {
			max = len(sub.Name)
		}
	}
	return max + 2
}
