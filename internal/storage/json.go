package storage

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/ArcaneCrowA/go-todo/internal/task"
)

type JSONStore struct {
	path string
}

func NewJSONStore(path string) *JSONStore {
	return &JSONStore{path: path}
}

func (s *JSONStore) Save(item task.Item) error {
	var items []task.Item
	var id int

	data, err := os.ReadFile(s.path)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}

	if len(data) > 0 {
		err = json.Unmarshal(data, &items)
		if err != nil {
			return err
		}
		id = items[len(items)-1].ID
	}

	item.ID = id
	items = append(items, item)

	data, err = json.Marshal(items)
	if err != nil {
		return err
	}

	return os.WriteFile(s.path, data, 0664)
}
