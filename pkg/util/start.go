// Package util implements different utilities required by the EdgeCluster service
package util

import (
	"log"
	"os"
	"os/signal"

	"github.com/decentralized-cloud/edge-cluster/services/business"
	"github.com/decentralized-cloud/edge-cluster/services/configuration"
	"github.com/decentralized-cloud/edge-cluster/services/edgecluster"
	edgeClusterTypes "github.com/decentralized-cloud/edge-cluster/services/edgecluster/types"
	"github.com/decentralized-cloud/edge-cluster/services/endpoint"
	"github.com/decentralized-cloud/edge-cluster/services/repository/mongodb"
	"github.com/decentralized-cloud/edge-cluster/services/transport/grpc"
	"github.com/decentralized-cloud/edge-cluster/services/transport/http"
	"github.com/micro-business/go-core/gokit/middleware"
	"go.uber.org/zap"
)

var configurationService configuration.ConfigurationContract
var endpointCreatorService endpoint.EndpointCreatorContract
var middlewareProviderService middleware.MiddlewareProviderContract

// StartService setups all dependecies required to start the EdgeCluster service and
// start the service
func StartService() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = logger.Sync()
	}()

	if err = setupDependencies(logger); err != nil {
		logger.Fatal("Failed to setup dependecies", zap.Error(err))
	}

	grpcTransportService, err := grpc.NewTransportService(
		logger,
		configurationService,
		endpointCreatorService,
		middlewareProviderService)
	if err != nil {
		logger.Fatal("Failed to create gRPC transport service", zap.Error(err))
	}

	httpTansportService, err := http.NewTransportService(
		logger,
		configurationService)
	if err != nil {
		logger.Fatal("Failed to create HTTP transport service", zap.Error(err))
	}

	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan struct{})
	signal.Notify(signalChan, os.Interrupt)

	go func() {
		if serviceErr := grpcTransportService.Start(); serviceErr != nil {
			logger.Fatal("Failed to start gRPC transport service", zap.Error(serviceErr))
		}
	}()

	go func() {
		if serviceErr := httpTansportService.Start(); serviceErr != nil {
			logger.Fatal("Failed to start HTTP transport service", zap.Error(serviceErr))
		}
	}()

	go func() {
		<-signalChan
		logger.Info("Received an interrupt, stopping services...")

		if err := grpcTransportService.Stop(); err != nil {
			logger.Error("Failed to stop gRPC transport service", zap.Error(err))
		}

		if err := httpTansportService.Stop(); err != nil {
			logger.Error("Failed to stop HTTP transport service", zap.Error(err))
		}

		close(cleanupDone)
	}()
	<-cleanupDone
}

func setupDependencies(logger *zap.Logger) (err error) {
	if configurationService, err = configuration.NewEnvConfigurationService(); err != nil {
		return
	}

	if middlewareProviderService, err = middleware.NewMiddlewareProviderService(logger, true, ""); err != nil {
		return
	}

	var edgeClusterFactoryService edgeClusterTypes.EdgeClusterFactoryContract
	if edgeClusterFactoryService, err = edgecluster.NewEdgeClusterFactoryService(logger, configurationService); err != nil {
		return
	}

	repositoryService, err := mongodb.NewMongodbRepositoryService(configurationService)
	if err != nil {
		return
	}

	businessService, err := business.NewBusinessService(repositoryService, edgeClusterFactoryService)
	if err != nil {
		return err
	}

	if endpointCreatorService, err = endpoint.NewEndpointCreatorService(businessService); err != nil {
		return
	}

	return
}
