package storage

import (
	"encoding/json"
	"os"

	"github.com/ArcaneCrowA/go-todo/internal/task"
)

type JSONStore struct {
	path string
}

func NewJSONStore(path string) JSONStore {
	return JSONStore{path: path}
}

func (j JSONStore) Save(item task.Item) error {
	file, err := os.OpenFile(j.path, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	var tasks []task.Item
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)
	if err != nil {
		return err
	}

	return nil
}
