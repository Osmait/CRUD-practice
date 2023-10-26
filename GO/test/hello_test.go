package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	filemanager "github.com/osmait/crud/internal/fileManager"
	"github.com/osmait/crud/internal/server"
	"github.com/osmait/crud/internal/storage"
)

func TestFindPost(t *testing.T) {
	// Inicializar el enrutador Gin
	router := gin.Default()
	fileManage := filemanager.NewFileManager("../db.json")
	postRepository := storage.NewPostRepository(fileManage)

	service := server.NewPostService(*postRepository)
	// Define tus rutas y controladores aquí
	router.GET("/", server.FindPost(*service))

	// Crea una solicitud HTTP de prueba
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Verifica la respuesta
	if w.Code != http.StatusOK {
		t.Errorf("El código de estado esperado no coincide. Esperado %d, obtenido %d", http.StatusOK, w.Code)
	}
}
