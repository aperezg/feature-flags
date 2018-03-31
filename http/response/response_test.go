package response_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aperezg/feature-flags/http/response"
	"github.com/aperezg/feature-flags/identity"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestNewResponse(t *testing.T) {
	r := responseOK()
	assert.IsType(t, &response.Response{}, r)
	assert.Equal(t, r.Status, response.StatusSuccess)
	assert.Empty(t, r.Message)
}

func TestNewResponseWithError(t *testing.T) {
	r := responseError()
	assert.IsType(t, &response.Response{}, r)
	assert.Empty(t, r.Data)
	assert.Equal(t, r.Status, response.StatusError)
	assert.Equal(t, r.Code, http.StatusInternalServerError)
}

var encodeTest = []*response.Response{
	responseOK(),
	responseError(),
}

func TestEncodeResponse(t *testing.T) {
	ctx := context.Background()
	w := httptest.NewRecorder()

	for _, r := range encodeTest {
		err := response.EncodeResponse(ctx, w, r)
		assert.NoError(t, err, "Error encoding response: %v", r)
	}
}

var responseData = map[string]string{
	identity.NewID(): "Alderaan",
	identity.NewID(): "Yavin IV",
	identity.NewID(): "Hoth",
}

func responseOK() *response.Response {
	return response.NewResponse(responseData)
}

func responseError() *response.Response {
	return response.NewResponse(
		nil,
		response.WithError(errors.New("error")),
		response.WithCode(http.StatusInternalServerError),
	)
}
