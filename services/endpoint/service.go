// Package endpoint implements different endpoint services required by the edge-cluster service
package endpoint

import (
	"context"

	"github.com/decentralized-cloud/edge-cluster/models"
	"github.com/decentralized-cloud/edge-cluster/services/business"
	"github.com/go-kit/kit/endpoint"
	commonErrors "github.com/micro-business/go-core/system/errors"
)

type endpointCreatorService struct {
	businessService business.BusinessContract
}

// NewEndpointCreatorService creates new instance of the EndpointCreatorService, setting up all dependencies and returns the instance
// businessService: Mandatory. Reference to the instance of the Edge Cluster  service
// Returns the new service or error if something goes wrong
func NewEndpointCreatorService(
	businessService business.BusinessContract) (EndpointCreatorContract, error) {
	if businessService == nil {
		return nil, commonErrors.NewArgumentNilError("businessService", "businessService is required")
	}

	return &endpointCreatorService{
		businessService: businessService,
	}, nil
}

// CreateEdgeClusterEndpoint creates Create Edge Cluster endpoint
// Returns the Create Edge Cluster endpoint
func (service *endpointCreatorService) CreateEdgeClusterEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		if ctx == nil {
			return &business.CreateEdgeClusterResponse{
				Err: commonErrors.NewArgumentNilError("ctx", "ctx is required"),
			}, nil
		}

		if request == nil {
			return &business.CreateEdgeClusterResponse{
				Err: commonErrors.NewArgumentNilError("request", "request is required"),
			}, nil
		}

		castedRequest := request.(*business.CreateEdgeClusterRequest)
		parsedToken := ctx.Value(models.ContextKeyParsedToken).(models.ParsedToken)
		castedRequest.UserEmail = parsedToken.Email

		if err := castedRequest.Validate(); err != nil {
			return &business.CreateEdgeClusterResponse{
				Err: commonErrors.NewArgumentErrorWithError("request", "", err),
			}, nil
		}

		return service.businessService.CreateEdgeCluster(ctx, castedRequest)
	}
}

// ReadEdgeClusterEndpoint creates Read Edge Cluster endpoint
// Returns the Read Edge Cluster endpoint
func (service *endpointCreatorService) ReadEdgeClusterEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		if ctx == nil {
			return &business.ReadEdgeClusterResponse{
				Err: commonErrors.NewArgumentNilError("ctx", "ctx is required"),
			}, nil
		}

		if request == nil {
			return &business.ReadEdgeClusterResponse{
				Err: commonErrors.NewArgumentNilError("request", "request is required"),
			}, nil
		}

		castedRequest := request.(*business.ReadEdgeClusterRequest)
		parsedToken := ctx.Value(models.ContextKeyParsedToken).(models.ParsedToken)
		castedRequest.UserEmail = parsedToken.Email

		if err := castedRequest.Validate(); err != nil {
			return &business.ReadEdgeClusterResponse{
				Err: commonErrors.NewArgumentErrorWithError("request", "", err),
			}, nil
		}

		return service.businessService.ReadEdgeCluster(ctx, castedRequest)
	}
}

// UpdateEdgeClusterEndpoint creates Update Edge Cluster endpoint
// Returns the Update Edge Cluster endpoint
func (service *endpointCreatorService) UpdateEdgeClusterEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		if ctx == nil {
			return &business.UpdateEdgeClusterResponse{
				Err: commonErrors.NewArgumentNilError("ctx", "ctx is required"),
			}, nil
		}

		if request == nil {
			return &business.UpdateEdgeClusterResponse{
				Err: commonErrors.NewArgumentNilError("request", "request is required"),
			}, nil
		}

		castedRequest := request.(*business.UpdateEdgeClusterRequest)
		parsedToken := ctx.Value(models.ContextKeyParsedToken).(models.ParsedToken)
		castedRequest.UserEmail = parsedToken.Email

		if err := castedRequest.Validate(); err != nil {
			return &business.UpdateEdgeClusterResponse{
				Err: commonErrors.NewArgumentErrorWithError("request", "", err),
			}, nil
		}

		return service.businessService.UpdateEdgeCluster(ctx, castedRequest)
	}
}

// DeleteEdgeClusterEndpoint creates Delete Edge Cluster endpoint
// Returns the Delete Edge Cluster endpoint
func (service *endpointCreatorService) DeleteEdgeClusterEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		if ctx == nil {
			return &business.DeleteEdgeClusterResponse{
				Err: commonErrors.NewArgumentNilError("ctx", "ctx is required"),
			}, nil
		}

		if request == nil {
			return &business.DeleteEdgeClusterResponse{
				Err: commonErrors.NewArgumentNilError("request", "request is required"),
			}, nil
		}

		castedRequest := request.(*business.DeleteEdgeClusterRequest)
		parsedToken := ctx.Value(models.ContextKeyParsedToken).(models.ParsedToken)
		castedRequest.UserEmail = parsedToken.Email

		if err := castedRequest.Validate(); err != nil {
			return &business.DeleteEdgeClusterResponse{
				Err: commonErrors.NewArgumentErrorWithError("request", "", err),
			}, nil
		}

		return service.businessService.DeleteEdgeCluster(ctx, castedRequest)
	}
}

// SearchEndpoint creates Search Edge Cluster endpoint
// Returns the Search Edge Cluster endpoint
func (service *endpointCreatorService) SearchEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		if ctx == nil {
			return &business.SearchResponse{
				Err: commonErrors.NewArgumentNilError("ctx", "ctx is required"),
			}, nil
		}

		if request == nil {
			return &business.SearchResponse{
				Err: commonErrors.NewArgumentNilError("request", "request is required"),
			}, nil
		}

		castedRequest := request.(*business.SearchRequest)
		parsedToken := ctx.Value(models.ContextKeyParsedToken).(models.ParsedToken)
		castedRequest.UserEmail = parsedToken.Email

		if err := castedRequest.Validate(); err != nil {
			return &business.SearchResponse{
				Err: commonErrors.NewArgumentErrorWithError("request", "", err),
			}, nil
		}

		return service.businessService.Search(ctx, castedRequest)
	}
}

// ListEdgeClusterNodesEndpoint creates List Edge Cluster Nodes endpoint
// Returns the List Edge Cluster Nodes endpoint
func (service *endpointCreatorService) ListEdgeClusterNodesEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		if ctx == nil {
			return &business.ListEdgeClusterNodesResponse{
				Err: commonErrors.NewArgumentNilError("ctx", "ctx is required"),
			}, nil
		}

		if request == nil {
			return &business.ListEdgeClusterNodesResponse{
				Err: commonErrors.NewArgumentNilError("request", "request is required"),
			}, nil
		}

		castedRequest := request.(*business.ListEdgeClusterNodesRequest)
		parsedToken := ctx.Value(models.ContextKeyParsedToken).(models.ParsedToken)
		castedRequest.UserEmail = parsedToken.Email

		if err := castedRequest.Validate(); err != nil {
			return &business.ListEdgeClusterNodesResponse{
				Err: commonErrors.NewArgumentErrorWithError("request", "", err),
			}, nil
		}

		return service.businessService.ListEdgeClusterNodes(ctx, castedRequest)
	}
}
