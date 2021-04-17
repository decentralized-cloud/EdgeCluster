#!/usr/bin/env sh

set -e
set -x

cleanup() {
	docker rm extract-mock-builder
}

trap 'cleanup' EXIT

if [ $# -eq 0 ]; then
	current_directory=$(dirname "$0")
else
	current_directory="$1"
fi

cd "$current_directory"/..

docker build -f docker/Dockerfile.buildMocks -t mock-builder .
docker create --name extract-mock-builder mock-builder
docker cp extract-mock-builder:/src/services/repository/mock/mock-contract.go ./services/repository/mock/mock-contract.go
docker cp extract-mock-builder:/src/services/transport/mock/mock-contract.go ./services/transport/mock/mock-contract.go
docker cp extract-mock-builder:/src/services/business/mock/mock-contract.go ./services/business/mock/mock-contract.go
docker cp extract-mock-builder:/src/services/configuration/mock/mock-contract.go ./services/configuration/mock/mock-contract.go
docker cp extract-mock-builder:/src/services/endpoint/mock/mock-contract.go ./services/endpoint/mock/mock-contract.go
docker cp extract-mock-builder:/src/services/edgecluster/types/mock/mock-contract.go ./services/edgecluster/types/mock/mock-contract.go
docker cp extract-mock-builder:/src/services/edgecluster/helm/mock/mock-contract.go ./services/edgecluster/helm/mock/mock-contract.go
docker cp extract-mock-builder:/src/services/cron/mock/mock-contract.go ./services/cron/mock/mock-contract.go
