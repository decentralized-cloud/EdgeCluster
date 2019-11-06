package business_test

import (
	"context"
	"errors"
	"math/rand"
	"strings"
	"testing"

	"github.com/decentralized-cloud/edge-cluster/models"
	"github.com/decentralized-cloud/edge-cluster/services/business"
	repository "github.com/decentralized-cloud/edge-cluster/services/repository"
	repsoitoryMock "github.com/decentralized-cloud/edge-cluster/services/repository/mock"
	"github.com/golang/mock/gomock"
	"github.com/lucsky/cuid"
	"github.com/micro-business/go-core/common"
	commonErrors "github.com/micro-business/go-core/system/errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBusinessService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Business Service Tests")
}

var _ = Describe("Business Service Tests", func() {
	var (
		mockCtrl              *gomock.Controller
		sut                   business.BusinessContract
		mockRepositoryService *repsoitoryMock.MockRepositoryContract
		ctx                   context.Context
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())

		mockRepositoryService = repsoitoryMock.NewMockRepositoryContract(mockCtrl)
		sut, _ = business.NewBusinessService(mockRepositoryService)
		ctx = context.Background()
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("user tries to instantiate BusinessService", func() {
		When("edge cluster repository service is not provided and NewBusinessService is called", func() {
			It("should return ArgumentNilError", func() {
				service, err := business.NewBusinessService(nil)
				Ω(service).Should(BeNil())
				assertArgumentNilError("repositoryService", "", err)
			})
		})

		When("all dependencies are resolved and NewBusinessService is called", func() {
			It("should instantiate the new BusinessService", func() {
				service, err := business.NewBusinessService(mockRepositoryService)
				Ω(err).Should(BeNil())
				Ω(service).ShouldNot(BeNil())
			})
		})
	})

	Describe("CreateEdgeCluster", func() {
		var (
			request business.CreateEdgeClusterRequest
		)

		BeforeEach(func() {
			request = business.CreateEdgeClusterRequest{
				EdgeCluster: models.EdgeCluster{
					TenantID: cuid.New(),
					Name:     cuid.New(),
				}}
		})

		Context("edge cluster service is instantiated", func() {
			When("CreateEdgeCluster is called", func() {
				It("should call edge cluster repository CreateEdgeCluster method", func() {
					mockRepositoryService.
						EXPECT().
						CreateEdgeCluster(ctx, gomock.Any()).
						Do(func(_ context.Context, mappedRequest *repository.CreateEdgeClusterRequest) {
							Ω(mappedRequest.EdgeCluster).Should(Equal(request.EdgeCluster))
						}).
						Return(&repository.CreateEdgeClusterResponse{
							EdgeClusterID: cuid.New(),
							EdgeCluster: models.EdgeCluster{
								TenantID: cuid.New(),
								Name:     cuid.New(),
							}}, nil)

					response, err := sut.CreateEdgeCluster(ctx, &request)
					Ω(err).Should(BeNil())
					Ω(response.Err).Should(BeNil())
				})

				When("And edge cluster repository CreateEdgeCluster return EdgeClusterAlreadyExistError", func() {
					It("should return EdgeClusterAlreadyExistsError", func() {
						expectedError := repository.NewEdgeClusterAlreadyExistsError()
						mockRepositoryService.
							EXPECT().
							CreateEdgeCluster(gomock.Any(), gomock.Any()).
							Return(nil, expectedError)

						response, err := sut.CreateEdgeCluster(ctx, &request)
						Ω(err).Should(BeNil())
						assertEdgeClusterAlreadyExistsError(response.Err, expectedError)
					})
				})

				When("And edge cluster repository CreateEdgeCluster return any other error", func() {
					It("should return UnknownError", func() {
						expectedError := errors.New(cuid.New())
						mockRepositoryService.
							EXPECT().
							CreateEdgeCluster(gomock.Any(), gomock.Any()).
							Return(nil, expectedError)

						response, err := sut.CreateEdgeCluster(ctx, &request)
						Ω(err).Should(BeNil())
						assertUnknowError(expectedError.Error(), response.Err, expectedError)
					})
				})

				When("And edge cluster repository CreateEdgeCluster return no error", func() {
					It("should return expected details", func() {
						expectedResponse := repository.CreateEdgeClusterResponse{
							EdgeClusterID: cuid.New(),
							EdgeCluster: models.EdgeCluster{
								TenantID: cuid.New(),
								Name:     cuid.New(),
							}}

						mockRepositoryService.
							EXPECT().
							CreateEdgeCluster(gomock.Any(), gomock.Any()).
							Return(&expectedResponse, nil)

						response, err := sut.CreateEdgeCluster(ctx, &request)
						Ω(err).Should(BeNil())
						Ω(response.Err).Should(BeNil())
						Ω(response.EdgeClusterID).ShouldNot(BeNil())
						Ω(response.EdgeClusterID).Should(Equal(expectedResponse.EdgeClusterID))
						assertEdgeCluster(response.EdgeCluster, expectedResponse.EdgeCluster)
					})
				})
			})
		})
	})

	Describe("ReadEdgeCluster", func() {
		var (
			request business.ReadEdgeClusterRequest
		)

		BeforeEach(func() {
			request = business.ReadEdgeClusterRequest{
				EdgeClusterID: cuid.New(),
			}
		})

		Context("edge cluster service is instantiated", func() {
			When("ReadEdgeCluster is called", func() {
				It("should call edge cluster repository ReadEdgeCluster method", func() {
					mockRepositoryService.
						EXPECT().
						ReadEdgeCluster(ctx, gomock.Any()).
						Do(func(_ context.Context, mappedRequest *repository.ReadEdgeClusterRequest) {
							Ω(mappedRequest.EdgeClusterID).Should(Equal(request.EdgeClusterID))
						}).
						Return(&repository.ReadEdgeClusterResponse{EdgeCluster: models.EdgeCluster{Name: cuid.New()}}, nil)

					response, err := sut.ReadEdgeCluster(ctx, &request)
					Ω(err).Should(BeNil())
					Ω(response.Err).Should(BeNil())
				})
			})

			When("And edge cluster repository ReadEdgeCluster cannot find provided edge cluster", func() {
				It("should return EdgeClusterNotFoundError", func() {
					expectedError := repository.NewEdgeClusterNotFoundError(request.EdgeClusterID)
					mockRepositoryService.
						EXPECT().
						ReadEdgeCluster(gomock.Any(), gomock.Any()).
						Return(nil, expectedError)

					response, err := sut.ReadEdgeCluster(ctx, &request)
					Ω(err).Should(BeNil())
					assertEdgeClusterNotFoundError(request.EdgeClusterID, response.Err, expectedError)
				})
			})

			When("And edge cluster repository ReadEdgeCluster return any other error", func() {
				It("should return UnknownError", func() {
					expectedError := errors.New(cuid.New())
					mockRepositoryService.
						EXPECT().
						ReadEdgeCluster(gomock.Any(), gomock.Any()).
						Return(nil, expectedError)

					response, err := sut.ReadEdgeCluster(ctx, &request)
					Ω(err).Should(BeNil())
					assertUnknowError(expectedError.Error(), response.Err, expectedError)
				})
			})

			When("And edge cluster repository ReadEdgeCluster return no error", func() {
				It("should return the edgeClusterID", func() {
					edgeCluster := models.EdgeCluster{Name: cuid.New()}
					mockRepositoryService.
						EXPECT().
						ReadEdgeCluster(gomock.Any(), gomock.Any()).
						Return(&repository.ReadEdgeClusterResponse{EdgeCluster: edgeCluster}, nil)

					response, err := sut.ReadEdgeCluster(ctx, &request)
					Ω(err).Should(BeNil())
					Ω(response.Err).Should(BeNil())
					assertEdgeCluster(response.EdgeCluster, edgeCluster)
				})
			})
		})
	})

	Describe("UpdateEdgeCluster", func() {
		var (
			request business.UpdateEdgeClusterRequest
		)

		BeforeEach(func() {
			request = business.UpdateEdgeClusterRequest{
				EdgeClusterID: cuid.New(),
				EdgeCluster:   models.EdgeCluster{Name: cuid.New()},
			}
		})

		Context("edge cluster service is instantiated", func() {
			When("UpdateEdgeCluster is called", func() {
				It("should call edge cluster repository UpdateEdgeCluster method", func() {
					mockRepositoryService.
						EXPECT().
						UpdateEdgeCluster(ctx, gomock.Any()).
						Do(func(_ context.Context, mappedRequest *repository.UpdateEdgeClusterRequest) {
							Ω(mappedRequest.EdgeClusterID).Should(Equal(request.EdgeClusterID))
							Ω(mappedRequest.EdgeCluster.Name).Should(Equal(request.EdgeCluster.Name))
						}).
						Return(&repository.UpdateEdgeClusterResponse{}, nil)

					response, err := sut.UpdateEdgeCluster(ctx, &request)
					Ω(err).Should(BeNil())
					Ω(response.Err).Should(BeNil())
				})
			})

			When("And edge cluster repository UpdateEdgeCluster cannot find provided edge cluster", func() {
				It("should return EdgeClusterNotFoundError", func() {
					expectedError := repository.NewEdgeClusterNotFoundError(request.EdgeClusterID)
					mockRepositoryService.
						EXPECT().
						UpdateEdgeCluster(gomock.Any(), gomock.Any()).
						Return(nil, expectedError)

					response, err := sut.UpdateEdgeCluster(ctx, &request)
					Ω(err).Should(BeNil())
					assertEdgeClusterNotFoundError(request.EdgeClusterID, response.Err, expectedError)
				})
			})

			When("And edge cluster repository UpdateEdgeCluster return any other error", func() {
				It("should return UnknownError", func() {
					expectedError := errors.New(cuid.New())
					mockRepositoryService.
						EXPECT().
						UpdateEdgeCluster(gomock.Any(), gomock.Any()).
						Return(nil, expectedError)

					response, err := sut.UpdateEdgeCluster(ctx, &request)
					Ω(err).Should(BeNil())
					assertUnknowError(expectedError.Error(), response.Err, expectedError)
				})
			})

			When("And edge cluster repository UpdateEdgeCluster return no error", func() {
				It("should return expected details", func() {
					expectedResponse := repository.UpdateEdgeClusterResponse{
						EdgeCluster: models.EdgeCluster{
							Name: cuid.New(),
						}}
					mockRepositoryService.
						EXPECT().
						UpdateEdgeCluster(gomock.Any(), gomock.Any()).
						Return(&expectedResponse, nil)

					response, err := sut.UpdateEdgeCluster(ctx, &request)
					Ω(err).Should(BeNil())
					Ω(response.Err).Should(BeNil())
					assertEdgeCluster(response.EdgeCluster, expectedResponse.EdgeCluster)
				})
			})
		})
	})

	Describe("DeleteEdgeCluster is called", func() {
		var (
			request business.DeleteEdgeClusterRequest
		)

		BeforeEach(func() {
			request = business.DeleteEdgeClusterRequest{
				EdgeClusterID: cuid.New(),
			}
		})

		Context("edge cluster service is instantiated", func() {
			When("DeleteEdgeCluster is called", func() {
				It("should call edge cluster repository DeleteEdgeCluster method", func() {
					mockRepositoryService.
						EXPECT().
						DeleteEdgeCluster(ctx, gomock.Any()).
						Do(func(_ context.Context, mappedRequest *repository.DeleteEdgeClusterRequest) {
							Ω(mappedRequest.EdgeClusterID).Should(Equal(request.EdgeClusterID))
						}).
						Return(&repository.DeleteEdgeClusterResponse{}, nil)

					response, err := sut.DeleteEdgeCluster(ctx, &request)
					Ω(err).Should(BeNil())
					Ω(response.Err).Should(BeNil())
				})
			})

			When("edge cluster repository DeleteEdgeCluster cannot find provided edge cluster", func() {
				It("should return EdgeClusterNotFoundError", func() {
					expectedError := repository.NewEdgeClusterNotFoundError(request.EdgeClusterID)
					mockRepositoryService.
						EXPECT().
						DeleteEdgeCluster(gomock.Any(), gomock.Any()).
						Return(nil, expectedError)

					response, err := sut.DeleteEdgeCluster(ctx, &request)
					Ω(err).Should(BeNil())
					assertEdgeClusterNotFoundError(request.EdgeClusterID, response.Err, expectedError)
				})
			})

			When("edge cluster repository DeleteEdgeCluster is faced with any other error", func() {
				It("should return UnknownError", func() {
					expectedError := errors.New(cuid.New())
					mockRepositoryService.
						EXPECT().
						DeleteEdgeCluster(gomock.Any(), gomock.Any()).
						Return(nil, expectedError)

					response, err := sut.DeleteEdgeCluster(ctx, &request)
					Ω(err).Should(BeNil())
					assertUnknowError(expectedError.Error(), response.Err, expectedError)
				})
			})

			When("edge cluster repository DeleteEdgeCluster completes successfully", func() {
				It("should return no error", func() {
					mockRepositoryService.
						EXPECT().
						DeleteEdgeCluster(gomock.Any(), gomock.Any()).
						Return(&repository.DeleteEdgeClusterResponse{}, nil)

					response, err := sut.DeleteEdgeCluster(ctx, &request)
					Ω(err).Should(BeNil())
					Ω(response.Err).Should(BeNil())
				})
			})
		})
	})

	Describe("Search is called", func() {
		var (
			request        business.SearchRequest
			edgeClusterIDs []string
			tenantIDs      []string
		)

		BeforeEach(func() {
			rand.Seed(42)
			edgeClusterIDs = []string{}
			for idx := 0; idx < rand.Intn(20)+1; idx++ {
				edgeClusterIDs = append(edgeClusterIDs, cuid.New())
			}

			tenantIDs = []string{}
			for idx := 0; idx < rand.Intn(20)+1; idx++ {
				tenantIDs = append(tenantIDs, cuid.New())
			}

			request = business.SearchRequest{
				Pagination: common.Pagination{
					After:  cuid.New(),
					First:  rand.Intn(1000),
					Before: cuid.New(),
					Last:   rand.Intn(1000),
				},
				SortingOptions: []common.SortingOptionPair{
					common.SortingOptionPair{
						Name:      cuid.New(),
						Direction: common.Ascending,
					},
					common.SortingOptionPair{
						Name:      cuid.New(),
						Direction: common.Descending,
					},
				},
				EdgeClusterIDs: edgeClusterIDs,
				TenantIDs:      tenantIDs,
			}
		})

		Context("edge cluster service is instantiated", func() {
			When("Search is called", func() {
				It("should call edge cluster repository Search method", func() {
					mockRepositoryService.
						EXPECT().
						Search(ctx, gomock.Any()).
						Do(func(_ context.Context, mappedRequest *repository.SearchRequest) {
							Ω(mappedRequest.Pagination).Should(Equal(request.Pagination))
							Ω(mappedRequest.SortingOptions).Should(Equal(request.SortingOptions))
							Ω(mappedRequest.EdgeClusterIDs).Should(Equal(request.EdgeClusterIDs))
							Ω(mappedRequest.TenantIDs).Should(Equal(request.TenantIDs))
						}).
						Return(&repository.SearchResponse{}, nil)

					response, err := sut.Search(ctx, &request)
					Ω(err).Should(BeNil())
					Ω(response.Err).Should(BeNil())
				})
			})

			When("edge cluster repository Search is faced with any other error", func() {
				It("should return UnknownError", func() {
					expectedError := errors.New(cuid.New())
					mockRepositoryService.
						EXPECT().
						Search(gomock.Any(), gomock.Any()).
						Return(nil, expectedError)

					response, err := sut.Search(ctx, &request)
					Ω(err).Should(BeNil())
					assertUnknowError(expectedError.Error(), response.Err, expectedError)
				})
			})

			When("edge cluster repository Search completes successfully", func() {
				It("should return the list of matched edgeClusterIDs", func() {
					edgeClusters := []models.EdgeClusterWithCursor{}

					for idx := 0; idx < rand.Intn(20)+1; idx++ {
						edgeClusters = append(edgeClusters, models.EdgeClusterWithCursor{
							EdgeClusterID: cuid.New(),
							EdgeCluster: models.EdgeCluster{
								TenantID: cuid.New(),
								Name:     cuid.New(),
							},
							Cursor: cuid.New(),
						})
					}

					expectedResponse := repository.SearchResponse{
						HasPreviousPage: (rand.Intn(10) % 2) == 0,
						HasNextPage:     (rand.Intn(10) % 2) == 0,
						EdgeClusters:    edgeClusters,
					}

					mockRepositoryService.
						EXPECT().
						Search(gomock.Any(), gomock.Any()).
						Return(&expectedResponse, nil)

					response, err := sut.Search(ctx, &request)
					Ω(err).Should(BeNil())
					Ω(response.Err).Should(BeNil())
					Ω(response.HasPreviousPage).Should(Equal(expectedResponse.HasPreviousPage))
					Ω(response.HasNextPage).Should(Equal(expectedResponse.HasNextPage))
					Ω(response.EdgeClusters).Should(Equal(expectedResponse.EdgeClusters))
				})
			})
		})
	})
})

func assertArgumentNilError(expectedArgumentName, expectedMessage string, err error) {
	Ω(commonErrors.IsArgumentNilError(err)).Should(BeTrue())

	var argumentNilErr commonErrors.ArgumentNilError
	_ = errors.As(err, &argumentNilErr)

	if expectedArgumentName != "" {
		Ω(argumentNilErr.ArgumentName).Should(Equal(expectedArgumentName))
	}

	if expectedMessage != "" {
		Ω(strings.Contains(argumentNilErr.Error(), expectedMessage)).Should(BeTrue())
	}
}

func assertUnknowError(expectedMessage string, err error, nestedErr error) {
	Ω(business.IsUnknownError(err)).Should(BeTrue())

	var unknownErr business.UnknownError
	_ = errors.As(err, &unknownErr)

	Ω(strings.Contains(unknownErr.Error(), expectedMessage)).Should(BeTrue())
	Ω(errors.Unwrap(err)).Should(Equal(nestedErr))
}

func assertEdgeClusterAlreadyExistsError(err error, nestedErr error) {
	Ω(business.IsEdgeClusterAlreadyExistsError(err)).Should(BeTrue())
	Ω(errors.Unwrap(err)).Should(Equal(nestedErr))
}

func assertEdgeClusterNotFoundError(expectedEdgeClusterID string, err error, nestedErr error) {
	Ω(business.IsEdgeClusterNotFoundError(err)).Should(BeTrue())

	var edgeClusterNotFoundErr business.EdgeClusterNotFoundError
	_ = errors.As(err, &edgeClusterNotFoundErr)

	Ω(edgeClusterNotFoundErr.EdgeClusterID).Should(Equal(expectedEdgeClusterID))
	Ω(errors.Unwrap(err)).Should(Equal(nestedErr))
}

func assertEdgeCluster(edgeCluster, expectedEdgeCluster models.EdgeCluster) {
	Ω(edgeCluster).ShouldNot(BeNil())
	Ω(edgeCluster.TenantID).Should(Equal(expectedEdgeCluster.TenantID))
	Ω(edgeCluster.Name).Should(Equal(expectedEdgeCluster.Name))
}
