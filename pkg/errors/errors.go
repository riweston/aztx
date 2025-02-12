// Package errors provides a centralized error handling system for the aztx application.
// It defines custom error types and error wrapping functions to provide consistent
// error handling and reporting across the application.
package errors

import (
	"errors"
	"fmt"
)

var (
	// Storage related errors

	// ErrFileDoesNotExist is returned when attempting to access a non-existent file
	ErrFileDoesNotExist = errors.New("file does not exist")
	// ErrFetchingHomePath is returned when unable to determine the user's home directory
	ErrFetchingHomePath = errors.New("could not fetch home directory")
	// ErrPathIsEmpty is returned when a required file path is empty
	ErrPathIsEmpty = errors.New("path is empty")

	// Configuration related errors

	// ErrReadingConfiguration wraps errors that occur while reading configuration files
	ErrReadingConfiguration = func(err error) error {
		return fmt.Errorf("error reading configuration: %w", err)
	}
	// ErrWritingConfiguration wraps errors that occur while writing configuration files
	ErrWritingConfiguration = func(err error) error {
		return fmt.Errorf("error writing configuration: %w", err)
	}

	// Context related errors

	// ErrNoPreviousContext is returned when attempting to switch to a previous context that doesn't exist
	ErrNoPreviousContext = errors.New("no previous context, check ~/.aztx.yml is present and has content")

	// Subscription related errors

	// ErrSubscriptionNotFound is returned when a requested subscription cannot be found
	ErrSubscriptionNotFound = errors.New("subscription not found")

	// File operation errors

	// ErrFileOperation wraps errors that occur during file operations with context about the operation
	ErrFileOperation = func(op string, err error) error {
		return fmt.Errorf("error %s file: %w", op, err)
	}

	// Generic operation errors

	// ErrOperation wraps generic operation errors with context about the operation
	ErrOperation = func(op string, err error) error {
		return fmt.Errorf("error during %s: %w", op, err)
	}

	// ErrMarshallingJSON wraps errors that occur during JSON marshalling
	ErrMarshallingJSON = func(err error) error {
		return fmt.Errorf("error marshalling JSON: %w", err)
	}

	// ErrUnmarshallingJSON wraps errors that occur during JSON unmarshalling
	ErrUnmarshallingJSON = func(err error) error {
		return fmt.Errorf("error unmarshalling JSON: %w", err)
	}

	// ErrReadingFile wraps errors that occur while reading files
	ErrReadingFile = func(err error) error {
		return fmt.Errorf("error reading file: %w", err)
	}

	// ErrWritingFile wraps errors that occur while writing files
	ErrWritingFile = func(err error) error {
		return fmt.Errorf("error writing file: %w", err)
	}

	// ErrSelectingSubscription wraps errors that occur during subscription selection
	ErrSelectingSubscription = func(err error) error {
		return fmt.Errorf("error selecting subscription: %w", err)
	}

	// ErrSettingPreviousContext wraps errors that occur while setting the previous context
	ErrSettingPreviousContext = func(err error) error {
		return fmt.Errorf("error setting previous context: %w", err)
	}

	// Context validation errors

	// ErrInvalidContext is returned when a context is missing required fields
	ErrInvalidContext = errors.New("invalid context: missing required fields")
	// ErrInvalidSubscriptionID is returned when a subscription ID has an invalid format
	ErrInvalidSubscriptionID = errors.New("invalid subscription ID format")

	// State related errors

	// ErrNoDefaultSubscription is returned when no default subscription is configured
	ErrNoDefaultSubscription = errors.New("no default subscription found in configuration")

	// Validation errors

	// ErrEmptyConfiguration is returned when the configuration is empty or nil
	ErrEmptyConfiguration = errors.New("configuration is empty or nil")
	// ErrInvalidTenantID is returned when a tenant ID has an invalid format
	ErrInvalidTenantID = errors.New("invalid tenant ID format")
	// ErrEmptyTenantName is returned when a tenant name is empty
	ErrEmptyTenantName = errors.New("tenant name cannot be empty")

	// Tenant operation errors

	// ErrTenantOperation wraps errors that occur during tenant operations
	ErrTenantOperation = func(op string, err error) error {
		return fmt.Errorf("error during tenant operation %s: %w", op, err)
	}

	// ErrTenantNotFound is returned when a tenant is not found
	ErrTenantNotFound = errors.New("tenant not found")
)

// WrapError is a helper function that wraps an error with operation context.
// It takes an operation name and an error, and returns a new error with additional context.
func WrapError(op string, err error) error {
	return fmt.Errorf("operation %s failed: %w", op, err)
}
