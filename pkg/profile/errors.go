package profile

import (
	"errors"
	"fmt"
)

// The errors file contains the error types used by the azure_cli package.

var (
	// ErrFileDoesNotExist is returned when the file does not exist.
	ErrFileDoesNotExist = func(err error) error {
		return fmt.Errorf("file does not exist: %w", err)
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

	// ErrWritingFile is returned when there is an error writing the file.
	ErrWritingFile = func(err error) error {
		return fmt.Errorf("error writing file: %w", err)
	}

	// ErrReadingFile is returned when there is an error reading the file.
	ErrReadingFile = func(err error) error {
		return fmt.Errorf("error reading file: %w", err)
	}

	// ErrPathIsEmpty is returned when the sampleConfigFilePath is empty.
	ErrPathIsEmpty = errors.New("sampleConfigFilePath is empty")

	// ErrReadingConfiguration is returned when there is an error reading the configuration.
	ErrReadingConfiguration = func(err error) error {
		return fmt.Errorf("error reading configuration: %w", err)
	}

	// ErrWritingConfiguration is returned when there is an error writing the configuration.
	ErrWritingConfiguration = func(err error) error {
		return fmt.Errorf("error writing configuration: %w", err)
	}

	// ErrSelectingSubscription is returned when there is an error selecting the subscription.
	ErrSelectingSubscription = func(err error) error {
		return fmt.Errorf("error selecting subscription: %w", err)
	}

	// ErrSettingPreviousContext is returned when there is an error setting the previous context.
	ErrSettingPreviousContext = func(err error) error {
		return fmt.Errorf("error setting previous context: %w", err)
	}

	// ErrNoPreviousContext is returned when there is no previous context.
	ErrNoPreviousContext = errors.New("no previous context, check ~/.aztx.yml is present and has content")

	// ErrSubscriptionNotFound is returned when there is no subscription found.
	ErrSubscriptionNotFound = errors.New("no subscription found")

	// ErrFetchingUserProfile is returned when there is an error fetching the user profile.
	ErrFetchingUserProfile = func(err error) error {
		return fmt.Errorf("error fetching user profile: %w", err)
	}
)
