// Package finder provides utilities for finding and selecting items using fuzzy search
// and ID-based lookups. It is primarily used for interactive selection of Azure
// resources like tenants and subscriptions.
package finder

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/ktr0731/go-fuzzyfinder"
)

// IDGetter is an interface that both Tenant and Subscription implement
type IDGetter interface {
	GetID() uuid.UUID
}

// Fuzzy is a utility function that provides interactive fuzzy finding capabilities
func Fuzzy[T any](items []T, displayFunc func(T) string) (*T, error) {
	if len(items) == 0 {
		return nil, fmt.Errorf("no items to select from")
	}

	idx, err := fuzzyfinder.Find(
		items,
		func(i int) string {
			return displayFunc(items[i])
		},
	)
	if err != nil {
		return nil, err
	}

	return &items[idx], nil
}

// ByID finds an item by its UUID in a slice of items that implement IDGetter
func ByID[T IDGetter](items []T, id uuid.UUID) (*T, error) {
	for _, item := range items {
		if item.GetID() == id {
			return &item, nil
		}
	}
	return nil, fmt.Errorf("item not found")
}
