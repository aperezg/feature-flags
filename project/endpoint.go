package project

import (
	"context"
	"net/http"

	"github.com/aperezg/feature-flags/http/response"
	"github.com/go-kit/kit/endpoint"
	"github.com/pkg/errors"
)

type createProjectRequest struct {
	Name string
}

func makeCreateProjectEndpoint(ps Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(createProjectRequest)
		p, err := ps.CreateProject(req.Name)

		r := response.NewResponse(p,
			response.WithError(err),
			response.WithCode(errorCode(err)),
		)

		return r, nil
	}
}

func errorCode(err error) (errCode int) {
	if err != nil {
		switch errors.Cause(err).Error() {
		case errorProjectAlreadyExists:
			errCode = http.StatusConflict
		case errorCreatingProject:
			errCode = http.StatusUnprocessableEntity
		default:
			errCode = http.StatusInternalServerError
		}
	}

	return
}
