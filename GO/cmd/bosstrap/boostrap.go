package cmd

import (
	"context"
	"time"

	filemanager "github.com/osmait/crud/internal/fileManager"
	internal "github.com/osmait/crud/internal/server"
	storage "github.com/osmait/crud/internal/storage"
)

const (
	host            = "localhost"
	port            = 8000
	path            = "../db.json"
	shutdownTimeout = 10 * time.Second
)

func Run() error {
	fileManage := filemanager.NewFileManager(path)
	postRepository := storage.NewPostRepository(fileManage)

	postService := internal.NewPostService(*postRepository)
	ctx, srv := internal.NewServer(context.Background(), host, port, *postService, shutdownTimeout)
	return srv.Run(ctx)
}
