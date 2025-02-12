package subscription

import (
	"fmt"

	"github.com/google/uuid"
	pkgerrors "github.com/riweston/aztx/pkg/errors"
	"github.com/riweston/aztx/pkg/finder"
	"github.com/riweston/aztx/pkg/types"
)

type Manager struct {
	types.BaseManager
}

// SetDefaultSubscription marks a subscription as default by its UUID.
func (sm *Manager) SetDefaultSubscription(subscriptionID uuid.UUID) error {
	for i, sub := range sm.Configuration.Subscriptions {
		if sub.ID == subscriptionID {
			sm.Configuration.Subscriptions[i].IsDefault = true
		} else {
			sm.Configuration.Subscriptions[i].IsDefault = false
		}
	}
	return nil
}

// FindSubscription searches for a subscription by name and returns it.
func (sm *Manager) FindSubscription(name string) (*types.Subscription, error) {
	for _, sub := range sm.Configuration.Subscriptions {
		if sub.Name == name {
			return &sub, nil
		}
	}
	return nil, pkgerrors.ErrSubscriptionNotFound
}

// FindSubscriptionIndex uses fuzzy finding to let user select a subscription
func (sm *Manager) FindSubscriptionIndex() (int, error) {
	if len(sm.Configuration.Subscriptions) == 0 {
		return -1, pkgerrors.ErrSubscriptionNotFound
	}

	sub, err := finder.Fuzzy(sm.Configuration.Subscriptions, func(s types.Subscription) string {
		return fmt.Sprintf("%s (%s)", s.Name, s.ID)
	})
	if err != nil {
		return -1, err
	}

	// Find the index of the selected subscription
	for i, s := range sm.Configuration.Subscriptions {
		if s.ID == sub.ID {
			return i, nil
		}
	}

	return -1, pkgerrors.ErrSubscriptionNotFound
}

// FindSubscriptionByID finds a subscription by its ID
func (sm *Manager) FindSubscriptionByID(id uuid.UUID) (*types.Subscription, error) {
	return finder.ByID(sm.Configuration.Subscriptions, id)
}

// FindSubscriptionsByTenant returns subscriptions filtered by tenant ID
func (sm *Manager) FindSubscriptionsByTenant(tenantID uuid.UUID) ([]types.Subscription, error) {
	var tenantSubs []types.Subscription
	for _, sub := range sm.Configuration.Subscriptions {
		if sub.TenantID == tenantID {
			tenantSubs = append(tenantSubs, sub)
		}
	}
	if len(tenantSubs) == 0 {
		return nil, pkgerrors.ErrSubscriptionNotFound
	}
	return tenantSubs, nil
}

// FindSubscriptionIndexByTenant uses fuzzy finding to select a subscription from a specific tenant
func (sm *Manager) FindSubscriptionIndexByTenant(tenantID uuid.UUID) (*types.Subscription, error) {
	subs, err := sm.FindSubscriptionsByTenant(tenantID)
	if err != nil {
		return nil, err
	}

	return finder.Fuzzy(subs, func(s types.Subscription) string {
		return fmt.Sprintf("%s (%s)", s.Name, s.ID)
	})
}
