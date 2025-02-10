package subscription

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/ktr0731/go-fuzzyfinder"
	"github.com/riweston/aztx/pkg/types"
)

var ErrSubscriptionNotFound = errors.New("subscription not found")

type SubscriptionManager struct {
	Configuration *types.Configuration
}

// SetDefaultSubscription marks a subscription as default by its UUID.
func (sm *SubscriptionManager) SetDefaultSubscription(subscriptionID uuid.UUID) error {
	for i, sub := range sm.Configuration.Subscriptions {
		if sub.ID == subscriptionID {
			sm.Configuration.Subscriptions[i].IsDefault = true
		} else {
			sm.Configuration.Subscriptions[i].IsDefault = false
		}
	}
	return nil
}

// FindSubscription searches for a subscription by name and returns its index.
func (sm *SubscriptionManager) FindSubscription(name string) (*types.Subscription, error) {
	for _, sub := range sm.Configuration.Subscriptions {
		if sub.Name == name {
			return &sub, nil
		}
	}
	return nil, ErrSubscriptionNotFound
}

// FindSubscriptionIndex uses fuzzy finding to let user select a subscription
func (sm *SubscriptionManager) FindSubscriptionIndex() (int, error) {
	items := make([]string, len(sm.Configuration.Subscriptions))
	for i, sub := range sm.Configuration.Subscriptions {
		items[i] = fmt.Sprintf("%s (%s)", sub.Name, sub.ID)
	}

	idx, err := fuzzyfinder.Find(
		items,
		func(i int) string {
			return items[i]
		},
	)
	if err != nil {
		return -1, err
	}

	return idx, nil
}

// FindSubscriptionByID finds a subscription by its ID
func (sm *SubscriptionManager) FindSubscriptionByID(id uuid.UUID) (*types.Subscription, error) {
	for _, sub := range sm.Configuration.Subscriptions {
		if sub.ID == id {
			return &sub, nil
		}
	}
	return nil, ErrSubscriptionNotFound
}
