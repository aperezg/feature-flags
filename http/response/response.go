package response

import (
	"context"
	"encoding/json"
	"net/http"
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
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Code    int         `json:"code,omitempty"`
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
		r.Status = StatusError
		r.Data = nil
		r.Message = err.Error()
	}
}

// WithCode add the http code to the Response.
func WithCode(code int) Option {
	return func(r *Response) {
		r.Code = code
	}
}

// EncodeResponse format the answer that will end up printing the api.
func EncodeResponse(ctx context.Context, w http.ResponseWriter, r *Response) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if r.Code != 0 {
		w.WriteHeader(r.Code)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	return json.NewEncoder(w).Encode(r)
}
