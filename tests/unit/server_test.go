package tests

import (
	"gotemplate/routes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServer(t *testing.T) {
	router := routes.Router
	routes.GetRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/private/health", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "\"OK\"", w.Body.String())
}
