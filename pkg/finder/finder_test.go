package finder

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type testItem struct {
	id   uuid.UUID
	name string
}

func (t testItem) GetID() uuid.UUID {
	return t.id
}

func TestByID(t *testing.T) {
	id1 := uuid.MustParse("a1a2a3a4-b1b2-c1c2-d1d2-d3d4d5d6d7d8")
	id2 := uuid.MustParse("b1b2b3b4-c1c2-d1d2-e1e2-e3e4e5e6e7e8")
	nonExistentID := uuid.MustParse("f1f2f3f4-e1e2-d1d2-c1c2-c3c4c5c6c7c8")

	tests := []struct {
		name      string
		items     []testItem
		searchID  uuid.UUID
		want      *testItem
		wantErr   bool
		errString string
	}{
		{
			name: "find existing item",
			items: []testItem{
				{id: id1, name: "item1"},
				{id: id2, name: "item2"},
			},
			searchID:  id1,
			want:      &testItem{id: id1, name: "item1"},
			wantErr:   false,
			errString: "",
		},
		{
			name: "item not found",
			items: []testItem{
				{id: id1, name: "item1"},
				{id: id2, name: "item2"},
			},
			searchID:  nonExistentID,
			want:      nil,
			wantErr:   true,
			errString: "item not found",
		},
		{
			name:      "empty slice returns error",
			items:     []testItem{},
			searchID:  id1,
			want:      nil,
			wantErr:   true,
			errString: "item not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ByID(tt.items, tt.searchID)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.errString, err.Error())
				assert.Nil(t, got)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestFuzzy(t *testing.T) {
	// We'll only test the error cases since the interactive part shouldn't be tested
	tests := []struct {
		name        string
		items       []testItem
		displayFunc func(testItem) string
		wantErr     bool
		errString   string
	}{
		{
			name:      "empty slice returns error",
			items:     []testItem{},
			wantErr:   true,
			errString: "no items to select from",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			displayFunc := func(i testItem) string { return i.name }
			_, err := Fuzzy(tt.items, displayFunc)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.errString, err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
