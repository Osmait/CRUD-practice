package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/osmait/crud/internal/domain"
)

func FindPost(postService PostService) gin.HandlerFunc {

	return func(c *gin.Context) {
		posts := postService.FindAll()
		c.JSON(200, posts)
	}
}

func CreatedPost(postService PostService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request domain.Post
		if err := ctx.BindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		postService.Created(ctx, request)
		ctx.Status(http.StatusCreated)
	}
}

func Upload() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		file, _ := ctx.FormFile("file")
		log.Println(file.Filename)
		dts := fmt.Sprintf("../upload/%s", file.Filename)
		ctx.SaveUploadedFile(file, dts)

		ctx.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))

	}
}
