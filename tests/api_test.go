package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/raphaelreis/goLabs/internal/api"
	"github.com/stretchr/testify/assert"
)

func TestHomeHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := api.SetupRouter()

	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetUsersHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := api.SetupRouter()

	req, _ := http.NewRequest("GET", "/users", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
