package response

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

// Status define the type of a Jsend Status.
type Status string

const (
	// StatusSuccess represent the type success on a Jsend Schema.
	StatusSuccess Status = "success"
	// StatusError represent the type error on a Jsend Schema.
	StatusError Status = "error"
)

// Response specification for schema Jsend.
type Response struct {
	Status  Status      `json:"status"`
	Code    int         `json:"code,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

// Option type to allow modeling the response dynamically.
type Option func(*Response)

// NewResponse initialize a Response with the data that will be printed by the api.
func NewResponse(data interface{}, opts ...Option) *Response {
	r := &Response{Status: StatusSuccess, Data: data}
	for _, opt := range opts {
		opt(r)
	}
	return r
}

// WithError modify the Response to model it as an error.
func WithError(err error) Option {
	return func(r *Response) {
		if err != nil {
			r.Status = StatusError
			r.Data = nil
			r.Message = err.Error()
		}
	}
}

// WithCode add the http code to the Response.
func WithCode(code int) Option {
	return func(r *Response) {
		r.Code = code
	}
}

// EncodeResponse format the answer that will end up printing the api.
func EncodeResponse(ctx context.Context, w http.ResponseWriter, r interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	resp := r.(*Response)
	if resp.Code != 0 {
		w.WriteHeader(resp.Code)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	return json.NewEncoder(w).Encode(r)
}

const (
	// HTTPStatusNotFound error for define Status Not Found api error.
	HTTPStatusNotFound = "StatusNotFound"
	// HTTPStatusBadRequest error for define Status Bad Request api error.
	HTTPStatusBadRequest = "StatusBadRequest"
)

// EncodeError format the errors that will end up printing the api.
func EncodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	var errCode int
	switch errors.Cause(err).Error() {
	case HTTPStatusNotFound:
		errCode = http.StatusNotFound
	case HTTPStatusBadRequest:
		errCode = http.StatusBadRequest
	default:
		errCode = http.StatusInternalServerError
	}
	w.WriteHeader(errCode)
	r := NewResponse(nil,
		WithError(err),
		WithCode(errCode))

	json.NewEncoder(w).Encode(r)
}
