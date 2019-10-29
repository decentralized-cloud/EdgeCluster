// Package memory implements im-memory repository services
package memory

import (
	"context"

	"github.com/decentralized-cloud/edge-cluster/models"
	"github.com/decentralized-cloud/edge-cluster/services/repository"

	"github.com/lucsky/cuid"
)

type edgeClusterWithTenant struct {
	tenantID    string
	edgeCluster models.EdgeCluster
}

type repositoryService struct {
	edgeClusterWithTenants map[string]*edgeClusterWithTenant
}

// NewRepositoryService creates new instance of the RepositoryService, setting up all dependencies and returns the instance
// Returns the new service or error if something goes wrong
func NewRepositoryService() (repository.RepositoryContract, error) {
	return &repositoryService{
		edgeClusterWithTenants: make(map[string]*edgeClusterWithTenant),
	}, nil
}

// CreateEdgeCluster creates a new edge cluster.
// context: Optional The reference to the context
// request: Mandatory. The request to create a new edge cluster
// Returns either the result of creating new edge cluster or error if something goes wrong.
func (service *repositoryService) CreateEdgeCluster(
	ctx context.Context,
	request *repository.CreateEdgeClusterRequest) (*repository.CreateEdgeClusterResponse, error) {
	edgeClusterID := cuid.New()
	service.edgeClusterWithTenants[edgeClusterID] = &edgeClusterWithTenant{
		tenantID:    request.TenantID,
		edgeCluster: request.EdgeCluster,
	}

	return &repository.CreateEdgeClusterResponse{
		EdgeClusterID: edgeClusterID,
	}, nil
}

// ReadEdgeCluster read an existing edge cluster
// context: Optional The reference to the context
// request: Mandatory. The request to read an existing edge cluster
// Returns either the result of reading an exiting edge cluster or error if something goes wrong.
func (service *repositoryService) ReadEdgeCluster(
	ctx context.Context,
	request *repository.ReadEdgeClusterRequest) (*repository.ReadEdgeClusterResponse, error) {
	edgeClusterWithTenant, ok := service.edgeClusterWithTenants[request.EdgeClusterID]
	if !ok {
		return nil, repository.NewEdgeClusterNotFoundError(request.EdgeClusterID)
	}

	return &repository.ReadEdgeClusterResponse{
		TenantID:    edgeClusterWithTenant.tenantID,
		EdgeCluster: edgeClusterWithTenant.edgeCluster,
	}, nil
}

// UpdateEdgeCluster update an existing edge cluster
// context: Optional The reference to the context
// request: Mandatory. The request to update an existing edge cluster
// Returns either the result of updateing an exiting edge cluster or error if something goes wrong.
func (service *repositoryService) UpdateEdgeCluster(
	ctx context.Context,
	request *repository.UpdateEdgeClusterRequest) (*repository.UpdateEdgeClusterResponse, error) {
	edgeClusterWithTenant, ok := service.edgeClusterWithTenants[request.EdgeClusterID]
	if !ok {
		return nil, repository.NewEdgeClusterNotFoundError(request.EdgeClusterID)
	}

	edgeClusterWithTenant.edgeCluster = request.EdgeCluster

	return &repository.UpdateEdgeClusterResponse{}, nil
}

// DeleteEdgeCluster delete an existing edge cluster
// context: Optional The reference to the context
// request: Mandatory. The request to delete an existing edge cluster
// Returns either the result of deleting an exiting edge cluster or error if something goes wrong.
func (service *repositoryService) DeleteEdgeCluster(
	ctx context.Context,
	request *repository.DeleteEdgeClusterRequest) (*repository.DeleteEdgeClusterResponse, error) {
	_, ok := service.edgeClusterWithTenants[request.EdgeClusterID]
	if !ok {
		return nil, repository.NewEdgeClusterNotFoundError(request.EdgeClusterID)
	}

	delete(service.edgeClusterWithTenants, request.EdgeClusterID)

	return &repository.DeleteEdgeClusterResponse{}, nil
}
