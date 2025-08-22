package json_storage

import (
	"fmt"
	"os"
)

type JsonStorage struct {
	*Storage
}

func NewJsonStorage(filename string) *JsonStorage {
	if len(filename) <= 0 {
		fmt.Println("File name too short")
		return &JsonStorage{}
	}

	return &JsonStorage{&Storage{filename: filename + ".json"}}
}

func (s *JsonStorage) Save(data []byte) error {
	err := os.WriteFile(s.filename, data, 0644)
	return err
}

func (s *JsonStorage) Load() ([]byte, error) {
	data, err := os.ReadFile(s.filename)
	return data, err
}
