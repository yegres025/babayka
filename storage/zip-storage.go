package storage

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"os"
)

type ZipStorage struct {
	*Storage
}

func NewZipStorage(filename string) *ZipStorage {
	if len(filename) <= 0 {
		fmt.Println("File name too short")
		return nil
	}

	return &ZipStorage{&Storage{filename: filename + ".zip"}}
}

func (z *ZipStorage) Save(data []byte) error {
	f, err := os.Create(z.filename)
	defer f.Close()
	if err != nil {
		return errors.New("File creation error " + err.Error())
	}

	zw := zip.NewWriter(f)
	defer zw.Close()

	w, err := zw.Create("data")
	_, err = w.Write(data)
	return err
}

func (z *ZipStorage) Load() ([]byte, error) {
	r, err := zip.OpenReader(z.filename)
	defer r.Close()
	if err != nil {
		return nil, errors.New("Read error -" + err.Error())
	}

	if len(r.File) == 0 {
		return nil, errors.New("Archive is empty")
	}

	file := r.File[0]
	rc, err := file.Open()
	defer rc.Close()
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return io.ReadAll(rc)
}
