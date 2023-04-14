package client

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApiRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		// Mock de la respuesta de la API
		res.WriteHeader(http.StatusOK)
		res.Write([]byte(`{"key": "value"}`))
	}))
	defer server.Close()

	originalCurrencyURL := os.Getenv("CURRENCY_URL")
	originalAPIKey := os.Getenv("API_KEY")
	originalTimeout := os.Getenv("TIMEOUT")

	defer func() {
		os.Setenv("CURRENCY_URL", originalCurrencyURL)
		os.Setenv("API_KEY", originalAPIKey)
		os.Setenv("TIMEOUT", originalTimeout)
	}()

	os.Setenv("CURRENCY_URL", server.URL+"/")
	os.Setenv("API_KEY", "my-api-key")
	os.Setenv("TIMEOUT", "5")

	resp, err := ApiRequest()

	assert.Nil(t, err)

	assert.NotNil(t, resp)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestApiRequest_InvalidTimeout(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusOK)
		res.Write([]byte(`{"key": "value"}`))
	}))
	defer server.Close()

	originalCurrencyURL := os.Getenv("CURRENCY_URL")
	originalAPIKey := os.Getenv("API_KEY")
	originalTimeout := os.Getenv("TIMEOUT")

	defer func() {
		os.Setenv("CURRENCY_URL", originalCurrencyURL)
		os.Setenv("API_KEY", originalAPIKey)
		os.Setenv("TIMEOUT", originalTimeout)
	}()

	os.Setenv("CURRENCY_URL", server.URL+"/")
	os.Setenv("API_KEY", "my-api-key")
	os.Setenv("TIMEOUT", "wrong_timeout")

	resp, err := ApiRequest()

	assert.Nil(t, resp)
	assert.NotNil(t, err)
}
