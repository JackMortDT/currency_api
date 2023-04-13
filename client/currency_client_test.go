package client

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApiRequest(t *testing.T) {
	// Crear un servidor mock
	server := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		// Mock de la respuesta de la API
		res.WriteHeader(http.StatusOK)
		res.Write([]byte(`{"key": "value"}`))
	}))
	defer server.Close()

	// Guardar los valores originales de las variables de entorno
	originalCurrencyURL := os.Getenv("CURRENCY_URL")
	originalAPIKey := os.Getenv("API_KEY")
	originalTimeout := os.Getenv("TIMEOUT")

	// Restaurar los valores de las variables de entorno al final de la prueba
	defer func() {
		os.Setenv("CURRENCY_URL", originalCurrencyURL)
		os.Setenv("API_KEY", originalAPIKey)
		os.Setenv("TIMEOUT", originalTimeout)
	}()

	// Establecer los valores de las variables de entorno necesarias para la prueba
	os.Setenv("CURRENCY_URL", server.URL+"/")
	os.Setenv("API_KEY", "my-api-key")
	os.Setenv("TIMEOUT", "5")

	// Llamar a la función que se está probando
	resp, err := ApiRequest()

	// Verificar que no haya errores
	assert.Nil(t, err)

	// Verificar que la respuesta no sea nula
	assert.NotNil(t, resp)

	// Verificar que el estado de la respuesta sea 200 OK
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
