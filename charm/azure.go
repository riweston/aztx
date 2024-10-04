package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/ktr0731/go-fuzzyfinder"
	azurestate "github.com/riweston/aztx/pkg/state"
)

// ProfileManager is the main interface for managing Azure profiles
type ProfileManager interface {
	SelectSubscription() (*Subscription, error)
	SetContext(lastContext *azurestate.LastContext, selectedContext *Subscription) error
	SetPreviousContext(lastContext *azurestate.LastContext) error
}

// AzureProfileManager implements ProfileManager
type AzureProfileManager struct {
	configPath string
	config     *Configuration
}

// NewAzureProfileManager creates a new AzureProfileManager
func NewAzureProfileManager() (*AzureProfileManager, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("error getting home directory: %w", err)
	}

	configPath := filepath.Join(homedir, ".azure", "azureProfile.json")
	return &AzureProfileManager{configPath: configPath}, nil
}

// loadConfig loads the configuration from file
func (a *AzureProfileManager) loadConfig() error {
	if a.config != nil {
		return nil
	}

	data, err := os.ReadFile(a.configPath)
	if err != nil {
		return fmt.Errorf("error reading config file: %w", err)
	}

	var config Configuration
	if err := json.Unmarshal(data, &config); err != nil {
		return fmt.Errorf("error unmarshalling config: %w", err)
	}

	a.config = &config
	return nil
}

// saveConfig saves the configuration to file
func (a *AzureProfileManager) saveConfig() error {
	data, err := json.MarshalIndent(a.config, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshalling config: %w", err)
	}

	if err := os.WriteFile(a.configPath, data, 0644); err != nil {
		return fmt.Errorf("error writing config file: %w", err)
	}

	return nil
}

// SelectSubscription allows the user to select a subscription
func (a *AzureProfileManager) SelectSubscription() (*Subscription, error) {
	if err := a.loadConfig(); err != nil {
		return nil, err
	}

	idx, err := a.findSubscription()
	if err != nil {
		return nil, fmt.Errorf("error selecting subscription: %w", err)
	}

	return &a.config.Subscriptions[idx], nil
}

// findSubscription uses fuzzy finder to let the user select a subscription
func (a *AzureProfileManager) findSubscription() (int, error) {
	maxWidth := a.longestSubscriptionNameWidth()
	format := fmt.Sprintf("%%-%ds %%s", maxWidth)

	idx, err := fuzzyfinder.Find(
		a.config.Subscriptions,
		func(i int) string {
			sub := a.config.Subscriptions[i]
			prefix := "  "
			if sub.IsDefault {
				prefix = "* "
			}
			return prefix + fmt.Sprintf(format, sub.Name, sub.ID)
		},
	)

	if err != nil {
		return 0, err
	}

	return idx, nil
}

// longestSubscriptionNameWidth returns the length of the longest subscription name
func (a *AzureProfileManager) longestSubscriptionNameWidth() int {
	maxWidth := 0
	for _, sub := range a.config.Subscriptions {
		if len(sub.Name) > maxWidth {
			maxWidth = len(sub.Name)
		}
	}
	return maxWidth + 2
}

// SetContext sets the current context
func (a *AzureProfileManager) SetContext(lastContext *azurestate.LastContext, selectedContext *Subscription) error {
	if err := a.loadConfig(); err != nil {
		return err
	}

	found := false
	for i := range a.config.Subscriptions {
		if a.config.Subscriptions[i].IsDefault {
			lastContext.WriteLastContext(a.config.Subscriptions[i].ID.String(), a.config.Subscriptions[i].Name)
		}
		a.config.Subscriptions[i].IsDefault = (a.config.Subscriptions[i].ID == selectedContext.ID)
		if a.config.Subscriptions[i].IsDefault {
			found = true
		}
	}

	if !found {
		return ErrSubscriptionNotFound
	}

	if err := a.saveConfig(); err != nil {
		return err
	}

	PrintInfo(fmt.Sprintf("Switched to \"%s\" (%s)", selectedContext.Name, selectedContext.ID))
	return nil
}

// SetPreviousContext sets the context to the previous one
func (a *AzureProfileManager) SetPreviousContext(lastContext *azurestate.LastContext) error {
	lastContextId := lastContext.ReadLastContextId()
	lastContextDisplayName := lastContext.ReadLastContextDisplayName()

	if lastContextId == "" || lastContextDisplayName == "" {
		return ErrNoPreviousContext
	}

	selectedContext := &Subscription{ID: uuid.MustParse(lastContextId), Name: lastContextDisplayName}
	return a.SetContext(lastContext, selectedContext)
}

// The rest of the code (Configuration, Subscription structs, and error definitions) remains the same
