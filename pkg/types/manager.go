package types

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/ktr0731/go-fuzzyfinder"
)

// BaseManager provides common functionality for tenant and subscription managers
type BaseManager struct {
	Configuration *Configuration
}

// FuzzyFindHelper is a utility function that can be used by both tenant and subscription managers
func FuzzyFindHelper[T any](items []T, displayFunc func(T) string) (*T, error) {
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

// IDGetter is an interface that both Tenant and Subscription implement
type IDGetter interface {
	GetID() uuid.UUID
}

// FindByIDHelper is a utility function to find an item by UUID
func FindByIDHelper[T IDGetter](items []T, id uuid.UUID) (*T, error) {
	for _, item := range items {
		if item.GetID() == id {
			return &item, nil
		}
	}
	return nil, fmt.Errorf("item not found")
}
