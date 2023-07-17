package cmd

import (
	"context"
	"os"
	"strconv"
	"time"

	filemanager "github.com/osmait/crud/internal/fileManager"
	internal "github.com/osmait/crud/internal/server"
	storage "github.com/osmait/crud/internal/storage"
)

const (
	path            = "../db.json"
	shutdownTimeout = 10 * time.Second
)

func Run() error {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	port2, _ := strconv.ParseUint(port, 10, 64)
	fileManage := filemanager.NewFileManager(path)
	postRepository := storage.NewPostRepository(fileManage)

	postService := internal.NewPostService(*postRepository)
	ctx, srv := internal.NewServer(context.Background(), host, uint(port2), *postService, shutdownTimeout)
	return srv.Run(ctx)
}
