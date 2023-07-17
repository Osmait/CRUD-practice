package filemanager

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/osmait/crud/internal/domain"
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
	var post []domain.Post
	if err != nil {

		os.Create("db.json")
		return post

	}

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
