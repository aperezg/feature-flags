package graphql

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/graphql-go/handler"
)

func makeGraphqlEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*handler.RequestOptions)
		res := s.Do(ctx, req)
		return res, nil
	}
}
