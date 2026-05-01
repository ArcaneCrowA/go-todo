package storage_test

import (
	"os"
	"testing"
	"time"

	"github.com/ArcaneCrowA/go-todo/internal/storage"
	"github.com/ArcaneCrowA/go-todo/internal/task"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestJSONLoad(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		setup   func(t *testing.T) string
		wantLen int
		wantErr bool
	}{
		{
			name: "file does not exist returns empty slice",
			setup: func(t *testing.T) string {
				return t.TempDir() + "/nonexistent.json"
			},
			wantLen: 0,
			wantErr: false,
		},
		{
			name: "empty file returns empty slice",
			setup: func(t *testing.T) string {
				path := t.TempDir() + "/empty.json"
				require.NoError(t, os.WriteFile(path, []byte{}, 0664))
				return path
			},
			wantLen: 0,
			wantErr: false,
		},
		{
			name: "file with items returns all items",
			setup: func(t *testing.T) string {
				path := t.TempDir() + "/items.json"
				st := storage.NewJSONStore(path)
				now := time.Now().Truncate(time.Minute)
				item := task.Item{
					ID:          1,
					Name:        "task1",
					Description: "desc1",
					Status:      task.ToDo,
					Created:     now,
					Updated:     now,
				}
				require.NoError(t, st.Save(item))
				return path
			},
			wantLen: 1,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := tt.setup(t)
			st := storage.NewJSONStore(path)

			items, err := st.Load()

			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Len(t, items, tt.wantLen)
		})
	}
}

func TestJSONSave(t *testing.T) {
	t.Parallel()

	now := time.Now().Truncate(time.Minute)

	tests := []struct {
		name    string
		setup   func(t *testing.T) string
		items   []task.Item
		wantIDs []int
		wantErr bool
	}{
		{
			name: "save first item gets ID 0",
			setup: func(t *testing.T) string {
				return t.TempDir() + "/storage.json"
			},
			items: []task.Item{
				{
					Name:        "first",
					Description: "first task",
					Status:      task.ToDo,
					Created:     now,
					Updated:     now,
				},
			},
			wantIDs: []int{0},
			wantErr: false,
		},
		{
			name: "multiple items get sequential IDs",
			setup: func(t *testing.T) string {
				return t.TempDir() + "/storage.json"
			},
			items: []task.Item{
				{
					Name:        "first",
					Description: "first task",
					Status:      task.ToDo,
					Created:     now,
					Updated:     now,
				},
				{
					Name:        "second",
					Description: "second task",
					Status:      task.InProgress,
					Created:     now,
					Updated:     now,
				},
				{
					Name:        "third",
					Description: "third task",
					Status:      task.Done,
					Created:     now,
					Updated:     now,
				},
			},
			wantIDs: []int{0, 1, 2},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := tt.setup(t)
			st := storage.NewJSONStore(path)

			for _, item := range tt.items {
				err := st.Save(item)
				if tt.wantErr {
					require.Error(t, err)
					return
				}
				require.NoError(t, err)
			}

			loaded, err := st.Load()
			require.NoError(t, err)
			require.Len(t, loaded, len(tt.items))

			for i, item := range loaded {
				assert.Equal(t, tt.items[i].Name, item.Name)
				assert.Equal(t, tt.items[i].Description, item.Description)
				assert.Equal(t, tt.items[i].Status, item.Status)
				if tt.wantIDs != nil {
					assert.Equal(t, tt.wantIDs[i], item.ID, "ID mismatch at index %d", i)
				}
			}
		})
	}
}

func TestJSONSavePersistsToFile(t *testing.T) {
	t.Parallel()

	tmpDir := t.TempDir()
	path := tmpDir + "/storage.json"

	st := storage.NewJSONStore(path)
	now := time.Now().Truncate(time.Minute)

	item := task.Item{
		Name:        "tests",
		Description: "description",
		Status:      task.ToDo,
		Created:     now,
		Updated:     now,
	}

	err := st.Save(item)
	require.NoError(t, err)

	b, err := os.ReadFile(path)
	require.NoError(t, err)
	s := string(b)
	assert.Contains(t, s, "test")
	assert.Contains(t, s, "desc")
}

func TestJSONDelete(t *testing.T) {
	t.Parallel()

	now := time.Now().Truncate(time.Minute)

	tests := []struct {
		name          string
		setup         func(t *testing.T) string
		itemsToSave   []task.Item
		itemsToDelete []task.Item
		wantRemaining int
		wantErr       bool
	}{
		{
			name: "delete existing item removes it",
			setup: func(t *testing.T) string {
				return t.TempDir() + "/storage.json"
			},
			itemsToSave: []task.Item{
				{ID: 0, Name: "delete", Description: "delete me", Status: task.ToDo, Created: now, Updated: now},
			},
			itemsToDelete: []task.Item{
				{ID: 0, Name: "delete", Description: "delete me", Status: task.ToDo, Created: now, Updated: now},
			},
			wantRemaining: 0,
			wantErr:       false,
		},
		{
			name: "delete non-existing ID is no-op",
			setup: func(t *testing.T) string {
				return t.TempDir() + "/storage.json"
			},
			itemsToSave: []task.Item{
				{ID: 0, Name: "keep", Description: "keep me", Status: task.ToDo, Created: now, Updated: now},
			},
			itemsToDelete: []task.Item{
				{ID: 99, Name: "ghost", Description: "not saved", Status: task.ToDo, Created: now, Updated: now},
			},
			wantRemaining: 1,
			wantErr:       false,
		},
		{
			name: "delete one of many items",
			setup: func(t *testing.T) string {
				path := t.TempDir() + "/storage.json"
				st := storage.NewJSONStore(path)
				for _, item := range []task.Item{
					{Name: "alpha", Description: "a", Status: task.ToDo, Created: now, Updated: now},
					{Name: "beta", Description: "b", Status: task.ToDo, Created: now, Updated: now},
					{Name: "gamma", Description: "c", Status: task.ToDo, Created: now, Updated: now},
				} {
					require.NoError(t, st.Save(item))
				}
				return path
			},
			itemsToDelete: []task.Item{
				{ID: 1, Name: "beta", Description: "b", Status: task.ToDo, Created: now, Updated: now},
			},
			wantRemaining: 2,
			wantErr:       false,
		},
		{
			name: "delete from empty store returns no error",
			setup: func(t *testing.T) string {
				return t.TempDir() + "/storage.json"
			},
			itemsToSave: []task.Item{},
			itemsToDelete: []task.Item{
				{ID: 1, Name: "nothing", Description: "nothing", Status: task.ToDo, Created: now, Updated: now},
			},
			wantRemaining: 0,
			wantErr:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := tt.setup(t)
			st := storage.NewJSONStore(path)

			for _, item := range tt.itemsToSave {
				require.NoError(t, st.Save(item))
			}

			for _, item := range tt.itemsToDelete {
				err := st.Delete(item)
				if tt.wantErr {
					require.Error(t, err)
					return
				}
				require.NoError(t, err)
			}

			loaded, err := st.Load()
			require.NoError(t, err)
			assert.Len(t, loaded, tt.wantRemaining)
		})
	}
}
