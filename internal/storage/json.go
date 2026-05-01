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
	items, err := s.Load()
	if err != nil {
		return err
	}

	item.ID = items[len(items)-1].ID
	items = append(items, item)

	return s.saveItems(items)
}

func (s *JSONStore) Delete(item task.Item) error {
	items, err := s.Load()
	if err != nil {
		return err
	}

	newItems := make([]task.Item, 0, len(items)-1)
	for _, it := range items {
		if item.ID == it.ID {
			continue
		}
		newItems = append(newItems, it)
	}

	return s.saveItems(newItems)
}

func (s *JSONStore) Load() ([]task.Item, error) {
	var items []task.Item

	data, err := os.ReadFile(s.path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {

		} else {
			return nil, err
		}
	}

	if len(data) > 0 {
		err = json.Unmarshal(data, &items)
		if err != nil {
			return nil, err
		}
	}

	return items, nil
}

func (s *JSONStore) saveItems(items []task.Item) error {
	data, err := json.Marshal(items)
	if err != nil {
		return err
	}

	return os.WriteFile(s.path, data, 0664)
}
