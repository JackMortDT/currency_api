package error_utils_test

import (
	"currency_api/utils/error_utils"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNotFoundError(t *testing.T) {
	errMsg := "test error message"
	err := error_utils.NewNotFoundError(errMsg)

	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status())
	assert.EqualValues(t, "not_found", err.Error())
	assert.EqualValues(t, errMsg, err.Message())
}

func TestNewBadRequestError(t *testing.T) {
	errMsg := "test error message"
	err := error_utils.NewBadRequestError(errMsg)

	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status())
	assert.EqualValues(t, "bad_request", err.Error())
	assert.EqualValues(t, errMsg, err.Message())
}

func TestNewInternalServerError(t *testing.T) {
	errMsg := "test error message"
	err := error_utils.NewInternalServerError(errMsg)

	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status())
	assert.EqualValues(t, "server_error", err.Error())
	assert.EqualValues(t, errMsg, err.Message())
}

func TestNewServiceUnavailableError(t *testing.T) {
	errMsg := "test error message"
	err := error_utils.NewServiceUnavailableError(errMsg)

	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusServiceUnavailable, err.Status())
	assert.EqualValues(t, "service_unavailable", err.Error())
	assert.EqualValues(t, errMsg, err.Message())
}
