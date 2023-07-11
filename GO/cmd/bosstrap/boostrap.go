package cmd

import (
	filemanager "github.com/osmait/crud/internal/fileManager"
	internal "github.com/osmait/crud/internal/server"
	storage "github.com/osmait/crud/internal/storage"
)

const (
	host = "localhost"
	port = 3000
	path = "../db.json"
)

func Run() error {
	fileManage := filemanager.NewFileManager(path)
	postRepository := storage.NewPostRepository(fileManage)

	postService := internal.NewPostService(*postRepository)
	srv := internal.NewServer(host, port, *postService)
	return srv.Run()
}
