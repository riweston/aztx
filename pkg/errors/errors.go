package errors

import (
	"errors"
	"fmt"
)

var (
	// Storage related errors
	ErrFileDoesNotExist = errors.New("file does not exist")
	ErrFetchingHomePath = errors.New("could not fetch home directory")
	ErrPathIsEmpty      = errors.New("path is empty")

	// Configuration related errors
	ErrReadingConfiguration = func(err error) error {
		return fmt.Errorf("error reading configuration: %w", err)
	}
	ErrWritingConfiguration = func(err error) error {
		return fmt.Errorf("error writing configuration: %w", err)
	}

	// Context related errors
	ErrNoPreviousContext = errors.New("no previous context, check ~/.aztx.yml is present and has content")

	// Subscription related errors
	ErrSubscriptionNotFound = errors.New("subscription not found")

	// File operation errors
	ErrFileOperation = func(op string, err error) error {
		return fmt.Errorf("error %s file: %w", op, err)
	}

	// Generic operation errors
	ErrOperation = func(op string, err error) error {
		return fmt.Errorf("error during %s: %w", op, err)
	}

	// ErrGettingHomeDirectory is returned when there is an error getting the home directory.
	ErrGettingHomeDirectory = func(err error) error {
		return fmt.Errorf("error getting home directory: %w", err)
	}

	// ErrMarshallingJSON is returned when there is an error marshalling JSON.
	ErrMarshallingJSON = func(err error) error {
		return fmt.Errorf("error marshalling JSON: %w", err)
	}

	// ErrUnmarshallingJSON is returned when there is an error unmarshalling JSON.
	ErrUnmarshallingJSON = func(err error) error {
		return fmt.Errorf("error unmarshalling JSON: %w", err)
	}

	// ErrReadingFile is returned when there is an error reading the file.
	ErrReadingFile = func(err error) error {
		return fmt.Errorf("error reading file: %w", err)
	}

	// ErrWritingFile is returned when there is an error writing the file.
	ErrWritingFile = func(err error) error {
		return fmt.Errorf("error writing file: %w", err)
	}

	// ErrSelectingSubscription is returned when there is an error selecting the subscription.
	ErrSelectingSubscription = func(err error) error {
		return fmt.Errorf("error selecting subscription: %w", err)
	}

	// ErrSettingPreviousContext is returned when there is an error setting the previous context.
	ErrSettingPreviousContext = func(err error) error {
		return fmt.Errorf("error setting previous context: %w", err)
	}

	// ErrFetchingUserProfile is returned when there is an error fetching the user profile.
	ErrFetchingUserProfile = func(err error) error {
		return fmt.Errorf("error fetching user profile: %w", err)
	}

	// Context validation errors
	ErrInvalidContext        = errors.New("invalid context: missing required fields")
	ErrInvalidSubscriptionID = errors.New("invalid subscription ID format")

	// State related errors
	ErrNoDefaultSubscription = errors.New("no default subscription found in configuration")

	// Validation errors
	ErrEmptyConfiguration = errors.New("configuration is empty or nil")
	ErrInvalidTenantID    = errors.New("invalid tenant ID format")

	// ErrEmptyTenantName is returned when the tenant name is empty.
	ErrEmptyTenantName = errors.New("tenant name cannot be empty")

	// Add new error types
	ErrInvalidEnvironmentName = errors.New("invalid environment name")
	ErrInvalidUserType        = errors.New("invalid user type")
	ErrDuplicateTenant        = errors.New("tenant already exists")

	// Add error wrapper for tenant operations
	ErrTenantOperation = func(op string, err error) error {
		return fmt.Errorf("error during tenant operation %s: %w", op, err)
	}
)

// Add helper for wrapping errors with operation context
func WrapError(op string, err error) error {
	return fmt.Errorf("operation %s failed: %w", op, err)
}
