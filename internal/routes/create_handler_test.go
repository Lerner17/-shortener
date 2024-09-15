package routes

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Lerner17/shortener/internal/helpers"
	"github.com/stretchr/testify/assert"
)

func TestCreateHandler(t *testing.T) {
	r := NewRouter()
	ts := httptest.NewServer(r)
	defer ts.Close()

	resp, _ := helpers.TestRequest(t, ts, http.MethodPost, "/", bytes.NewReader([]byte("https://yandex.ru")))
	defer resp.Body.Close()
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	resp, _ = helpers.TestRequest(t, ts, http.MethodGet, "/", nil)
	defer resp.Body.Close()
	assert.Equal(t, http.StatusMethodNotAllowed, resp.StatusCode)

	resp, _ = helpers.TestRequest(t, ts, http.MethodPost, "/", nil)
	defer resp.Body.Close()
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}