package graphql

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/graphql-go/handler"
)

func MakeHandler(gqs Service, logger log.Logger, serveropts ...kithttp.ServerOption) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
		kithttp.ServerErrorEncoder(encodeError),
	}

	opts = append(opts, serveropts...)

	graphqlHandler := kithttp.NewServer(
		makeGraphqlEndpoint(gqs),
		decodeGraphqlRequest,
		encodeResponse,
		opts...,
	)

	return graphqlHandler
}

func decodeGraphqlRequest(_ context.Context, req *http.Request) (interface{}, error) {
	return handler.NewRequestOptions(req), nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(response)
}

type errorer interface {
	error() error
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	// TODO: Handling error isn't an internal server error
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
