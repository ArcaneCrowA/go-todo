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

func TestJSONSave(t *testing.T) {
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
