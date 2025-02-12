package errors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorWrapping(t *testing.T) {
	baseErr := errors.New("base error")

	tests := []struct {
		name    string
		err     error
		wrapper func(error) error
		wantMsg string
	}{
		{
			name:    "wrap configuration read error",
			err:     baseErr,
			wrapper: ErrReadingConfiguration,
			wantMsg: "error reading configuration: base error",
		},
		{
			name:    "wrap configuration write error",
			err:     baseErr,
			wrapper: ErrWritingConfiguration,
			wantMsg: "error writing configuration: base error",
		},
		{
			name:    "wrap JSON marshalling error",
			err:     baseErr,
			wrapper: ErrMarshallingJSON,
			wantMsg: "error marshalling JSON: base error",
		},
		{
			name:    "wrap JSON unmarshalling error",
			err:     baseErr,
			wrapper: ErrUnmarshallingJSON,
			wantMsg: "error unmarshalling JSON: base error",
		},
		{
			name:    "wrap file read error",
			err:     baseErr,
			wrapper: ErrReadingFile,
			wantMsg: "error reading file: base error",
		},
		{
			name:    "wrap file write error",
			err:     baseErr,
			wrapper: ErrWritingFile,
			wantMsg: "error writing file: base error",
		},
		{
			name:    "wrap subscription selection error",
			err:     baseErr,
			wrapper: ErrSelectingSubscription,
			wantMsg: "error selecting subscription: base error",
		},
		{
			name:    "wrap previous context error",
			err:     baseErr,
			wrapper: ErrSettingPreviousContext,
			wantMsg: "error setting previous context: base error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wrappedErr := tt.wrapper(tt.err)
			assert.EqualError(t, wrappedErr, tt.wantMsg)
			assert.True(t, errors.Is(wrappedErr, baseErr), "wrapped error should contain the original error")
		})
	}
}

func TestStaticErrors(t *testing.T) {
	tests := []struct {
		name string
		err  error
		msg  string
	}{
		{
			name: "file does not exist error",
			err:  ErrFileDoesNotExist,
			msg:  "file does not exist",
		},
		{
			name: "fetching home path error",
			err:  ErrFetchingHomePath,
			msg:  "could not fetch home directory",
		},
		{
			name: "path is empty error",
			err:  ErrPathIsEmpty,
			msg:  "path is empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.EqualError(t, tt.err, tt.msg)
		})
	}
}
