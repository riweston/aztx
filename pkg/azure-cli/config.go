package azure_cli

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/ktr0731/go-fuzzyfinder"
	azurestate "github.com/riweston/aztx/pkg/azure-state"
)

type ConfigurationAdapter struct {
	userProfile userProfileReadWriter
}

func NewConfigurationAdapter(userProfile userProfileReadWriter) *ConfigurationAdapter {
	return &ConfigurationAdapter{
		userProfile: userProfile,
	}
}

func (c *ConfigurationAdapter) SelectWithFinder() (*Subscription, error) {
	cfg, err := c.userProfile.Read()
	if err != nil {
		return nil, ErrReadingConfiguration(err)
	}
	idx, err := c.userProfile.Find(cfg)
	if errors.Is(err, fuzzyfinder.ErrAbort) {
		PrintNotice("Operation aborted")
		return nil, err
	}
	if err != nil {
		return nil, ErrSelectingSubscription(err)
	}
	return &cfg.Subscriptions[idx], nil
}

func (c *ConfigurationAdapter) SetPreviousContext(lastContext *azurestate.LastContext) error {
	lastContextId := lastContext.ReadLastContextId()
	lastContextDisplayName := lastContext.ReadLastContextDisplayName()
	if lastContextId == "" || lastContextDisplayName == "" {
		return ErrNoPreviousContext
	}
	if err := c.SetContext(lastContext, &Subscription{ID: uuid.MustParse(lastContextId), Name: lastContextDisplayName}); err != nil {
		return ErrSettingPreviousContext(err)
	}
	return nil
}

func (c *ConfigurationAdapter) SetContext(lastContext *azurestate.LastContext, selectedContext *Subscription) error {
	cfg, err := c.userProfile.Read()
	var errNotFound bool
	if err != nil {
		return ErrReadingConfiguration(err)
	}
	for i, sub := range cfg.Subscriptions {
		if sub.IsDefault {
			lastContext.WriteLastContext(sub.ID.String(), sub.Name)
		}
		cfg.Subscriptions[i].IsDefault = false
		if sub.ID == selectedContext.ID {
			errNotFound = false
			cfg.Subscriptions[i].IsDefault = true
		}
	}
	if errNotFound {
		return ErrSubscriptionNotFound
	}
	err = c.userProfile.Write(cfg)
	if err != nil {
		return ErrWritingConfiguration(err)
	}
	PrintInfo(fmt.Sprintf("Switched to \"%s\" (%s)", selectedContext.Name, selectedContext.ID))
	return nil
}
