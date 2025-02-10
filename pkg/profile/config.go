package profile

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/riweston/aztx/pkg/errors"
	"github.com/riweston/aztx/pkg/state"
	"github.com/riweston/aztx/pkg/subscription"
	"github.com/riweston/aztx/pkg/tenant"
	"github.com/riweston/aztx/pkg/types"
)

type ConfigurationAdapter struct {
	storage StorageAdapter
	logger  Logger
}

func NewConfigurationAdapter(storage StorageAdapter, logger Logger) *ConfigurationAdapter {
	return &ConfigurationAdapter{
		storage: storage,
		logger:  logger,
	}
}

func (c *ConfigurationAdapter) SelectWithFinder() (*types.Subscription, error) {
	if c.storage == nil {
		c.logger.Error("storage adapter is nil")
		return nil, errors.ErrEmptyConfiguration
	}

	c.logger.Debug("reading azure profile configuration")
	config, err := c.storage.ReadConfig()
	if err != nil {
		c.logger.Error("failed to read configuration: %v", err)
		return nil, errors.WrapError("reading configuration", err)
	}

	if len(config.Subscriptions) == 0 {
		c.logger.Warn("no subscriptions found in configuration")
		return nil, errors.ErrEmptyConfiguration
	}

	c.logger.Debug("initiating subscription selection with fuzzy finder")
	subManager := subscription.Manager{Configuration: config}
	idx, err := subManager.FindSubscriptionIndex()
	if err != nil {
		c.logger.Error("failed to get subscription selection: %v", err)
		return nil, errors.WrapError("finding subscription", err)
	}

	if idx < 0 || idx >= len(config.Subscriptions) {
		c.logger.Error("selected subscription index %d is out of bounds", idx)
		return nil, errors.ErrSubscriptionNotFound
	}

	selected := &config.Subscriptions[idx]
	c.logger.Info("selected subscription: %s (%s)", selected.Name, selected.ID)
	return selected, nil
}

func (c *ConfigurationAdapter) SetContext(subscriptionID uuid.UUID) error {
	if subscriptionID == uuid.Nil {
		c.logger.Error("invalid subscription ID provided")
		return errors.ErrInvalidSubscriptionID
	}

	c.logger.Debug("reading configuration to update context")
	config, err := c.storage.ReadConfig()
	if err != nil {
		c.logger.Error("failed to read configuration: %v", err)
		return errors.WrapError("reading configuration", err)
	}

	found := false
	var newDefault string
	for i, sub := range config.Subscriptions {
		if sub.IsDefault {
			c.logger.Debug("clearing default from subscription: %s", sub.Name)
			config.Subscriptions[i].IsDefault = false
		}
		if sub.ID == subscriptionID {
			c.logger.Debug("setting new default subscription: %s", sub.Name)
			config.Subscriptions[i].IsDefault = true
			found = true
			newDefault = sub.Name
		}
	}

	if !found {
		c.logger.Error("subscription %s not found in configuration", subscriptionID)
		return errors.ErrSubscriptionNotFound
	}

	c.logger.Debug("writing updated configuration")
	if err := c.storage.WriteConfig(config); err != nil {
		c.logger.Error("failed to write configuration: %v", err)
		return errors.WrapError("writing configuration", err)
	}

	c.logger.Info("successfully switched context to: %s", newDefault)
	return nil
}

func (c *ConfigurationAdapter) SetPreviousContext(state state.StateManager) error {
	if state == nil {
		c.logger.Error("state manager is nil")
		return errors.ErrInvalidContext
	}

	lastId, lastName := state.GetLastContext()
	if lastId == "" || lastName == "" {
		c.logger.Warn("no previous context found")
		return errors.ErrNoPreviousContext
	}

	c.logger.Debug("reading configuration to switch to previous context")
	config, err := c.storage.ReadConfig()
	if err != nil {
		c.logger.Error("failed to read configuration: %v", err)
		return errors.WrapError("reading configuration", err)
	}

	var currentDefault *types.Subscription
	for _, sub := range config.Subscriptions {
		if sub.IsDefault {
			currentDefault = &sub
			break
		}
	}

	if currentDefault == nil {
		c.logger.Error("no default subscription found in configuration")
		return errors.ErrNoDefaultSubscription
	}

	c.logger.Debug("saving current context: %s", currentDefault.Name)
	if err := state.SetLastContext(currentDefault.ID.String(), currentDefault.Name); err != nil {
		c.logger.Error("failed to save current context: %v", err)
		return errors.WrapError("saving last context", err)
	}

	id, err := uuid.Parse(lastId)
	if err != nil {
		c.logger.Error("failed to parse previous subscription ID: %v", err)
		return errors.WrapError("parsing subscription ID", err)
	}

	c.logger.Info("switching to previous context: %s", lastName)
	return c.SetContext(id)
}

func (c *ConfigurationAdapter) SaveTenant(id uuid.UUID, name string) error {
	if id == uuid.Nil {
		return errors.ErrInvalidTenantID
	}

	if name == "" {
		return errors.ErrEmptyTenantName
	}

	config, err := c.storage.ReadConfig()
	if err != nil {
		return errors.WrapError("reading configuration", err)
	}

	tenantManager := tenant.Manager{Configuration: config}
	if err := tenantManager.SaveTenantName(id, name); err != nil {
		return errors.WrapError("saving tenant name", err)
	}

	if err := c.storage.WriteConfig(config); err != nil {
		return errors.WrapError("writing configuration", err)
	}

	return nil
}

// Add context to key operations
func (c *ConfigurationAdapter) SelectWithFinderContext(ctx context.Context) (*types.Subscription, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		return c.SelectWithFinder()
	}
}

func (c *ConfigurationAdapter) SetContextWithTimeout(subscriptionID uuid.UUID, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	done := make(chan error, 1)
	go func() {
		done <- c.SetContext(subscriptionID)
	}()

	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
