package project

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aperezg/feature-flags/http/response"
	"github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

func MakeHandler(ps Service, logger log.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
		kithttp.ServerErrorEncoder(response.EncodeError),
	}

	createProjectHandler := kithttp.NewServer(
		makeCreateProjectEndpoint(ps),
		decodeCreateProjectRequest,
		response.EncodeResponse,
		opts...,
	)

	r := mux.NewRouter()
	r.Handle("/projects", createProjectHandler).Methods("POST")

	return r
}

func decodeCreateProjectRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body struct {
		Name string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		err = errors.Wrap(err, response.HTTPStatusBadRequest)
		return nil, err
	}

	return createProjectRequest{
		Name: body.Name,
	}, nil
}
