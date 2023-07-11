package filemanager

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/osmait/crud/internal/domain"
	// storage "github.com/osmait/crud/internal/storage"
)

type FileManager struct {
	FilePath string
}

func NewFileManager(path string) *FileManager {
	return &FileManager{
		FilePath: path,
	}
}

func (f *FileManager) Get() []domain.Post {
	file, err := ioutil.ReadFile(f.FilePath)
	if err != nil {
		log.Fatalf("Error Read File")
	}
	var post []domain.Post
	err = json.Unmarshal(file, &post)

	if err != nil {
		log.Fatalf("Error deserial Json")
	}
	return post
}

func (f *FileManager) Write(post domain.Post) error {
	result := f.Get()
	result = append(result, post)
	modif, err := json.Marshal(&result)
	if err != nil {
		return err
	}
	ioutil.WriteFile(f.FilePath, modif, 0644)
	return nil
}
