// Code generated by MockGen. DO NOT EDIT.
// Source: services/edgecluster/helm/contract.go

// Package mock_helm is a generated GoMock package.
package mock_helm

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockHelmHelperContract is a mock of HelmHelperContract interface.
type MockHelmHelperContract struct {
	ctrl     *gomock.Controller
	recorder *MockHelmHelperContractMockRecorder
}

// MockHelmHelperContractMockRecorder is the mock recorder for MockHelmHelperContract.
type MockHelmHelperContractMockRecorder struct {
	mock *MockHelmHelperContract
}

// NewMockHelmHelperContract creates a new mock instance.
func NewMockHelmHelperContract(ctrl *gomock.Controller) *MockHelmHelperContract {
	mock := &MockHelmHelperContract{ctrl: ctrl}
	mock.recorder = &MockHelmHelperContractMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHelmHelperContract) EXPECT() *MockHelmHelperContractMockRecorder {
	return m.recorder
}

// AddRepository mocks base method.
func (m *MockHelmHelperContract) AddRepository(name, url string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRepository", name, url)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddRepository indicates an expected call of AddRepository.
func (mr *MockHelmHelperContractMockRecorder) AddRepository(name, url interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRepository", reflect.TypeOf((*MockHelmHelperContract)(nil).AddRepository), name, url)
}

// InstallChart mocks base method.
func (m *MockHelmHelperContract) InstallChart(kubeconfig, namespace, name, repo, chart string, args map[string]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstallChart", kubeconfig, namespace, name, repo, chart, args)
	ret0, _ := ret[0].(error)
	return ret0
}

// InstallChart indicates an expected call of InstallChart.
func (mr *MockHelmHelperContractMockRecorder) InstallChart(kubeconfig, namespace, name, repo, chart, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstallChart", reflect.TypeOf((*MockHelmHelperContract)(nil).InstallChart), kubeconfig, namespace, name, repo, chart, args)
}

// UpdateCharts mocks base method.
func (m *MockHelmHelperContract) UpdateCharts() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCharts")
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCharts indicates an expected call of UpdateCharts.
func (mr *MockHelmHelperContractMockRecorder) UpdateCharts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCharts", reflect.TypeOf((*MockHelmHelperContract)(nil).UpdateCharts))
}
