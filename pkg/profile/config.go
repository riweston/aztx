package profile

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/riweston/aztx/pkg/errors"
	"github.com/riweston/aztx/pkg/state"
	"github.com/riweston/aztx/pkg/storage"
	"github.com/riweston/aztx/pkg/subscription"
	"github.com/riweston/aztx/pkg/tenant"
	"github.com/riweston/aztx/pkg/types"
	"github.com/spf13/viper"
)

type ConfigurationAdapter struct {
	reader              storage.FileAdapter
	writer              storage.FileAdapter
	tenantService       tenant.TenantManager
	subscriptionService subscription.SubscriptionManager
}

func NewConfigurationAdapter(
	reader storage.FileAdapter,
	writer storage.FileAdapter,
) *ConfigurationAdapter {
	return &ConfigurationAdapter{
		reader: reader,
		writer: writer,
	}
}

func (c *ConfigurationAdapter) SelectWithFinder() (*types.Subscription, error) {
	config, err := c.reader.ReadConfig()
	if err != nil {
		return nil, errors.ErrReadingConfiguration(err)
	}

	subManager := subscription.SubscriptionManager{Configuration: config}
	idx, err := subManager.FindSubscriptionIndex()
	if err != nil {
		return nil, err
	}

	return &config.Subscriptions[idx], nil
}

func (c *ConfigurationAdapter) SetContext(subscriptionID uuid.UUID) error {
	config, err := c.reader.ReadConfig()
	if err != nil {
		return errors.ErrReadingConfiguration(err)
	}

	// Save current context before switching
	lc := state.NewStateReaderWriter(&state.ViperAdapter{Viper: viper.GetViper()})
	for _, sub := range config.Subscriptions {
		if sub.IsDefault {
			lc.WriteLastContext(sub.ID.String(), sub.Name)
			break
		}
	}

	subManager := subscription.SubscriptionManager{Configuration: config}
	if err := subManager.SetDefaultSubscription(subscriptionID); err != nil {
		return err
	}

	return c.writer.WriteConfig(config)
}

func (c *ConfigurationAdapter) SetPreviousContext(lastContext *state.LastContext) error {
	lastContextId := lastContext.ReadLastContextId()
	lastContextName := lastContext.ReadLastContextDisplayName()

	if lastContextId == "" || lastContextName == "" {
		return errors.ErrNoPreviousContext
	}

	// Get current context before switching
	config, err := c.reader.ReadConfig()
	if err != nil {
		return errors.ErrReadingConfiguration(err)
	}

	// Find current default subscription
	for _, sub := range config.Subscriptions {
		if sub.IsDefault {
			// Save current as last before switching
			lastContext.WriteLastContext(sub.ID.String(), sub.Name)
			break
		}
	}

	id, err := uuid.Parse(lastContextId)
	if err != nil {
		return fmt.Errorf("invalid last context ID: %w", err)
	}

	return c.SetContext(id)
}

func (c *ConfigurationAdapter) SaveTenant(id uuid.UUID, name string) error {
	config, err := c.reader.ReadConfig()
	if err != nil {
		return errors.ErrReadingConfiguration(err)
	}

	tenantManager := tenant.TenantManager{Configuration: config}
	if err := tenantManager.SaveTenantName(id, name); err != nil {
		return err
	}

	return c.writer.WriteConfig(config)
}
