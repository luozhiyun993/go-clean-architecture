// Code generated by MockGen. DO NOT EDIT.
// Source: .\service\service.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	models "my-clean-architecture/models"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockIArticleService is a mock of IArticleService interface.
type MockIArticleService struct {
	ctrl     *gomock.Controller
	recorder *MockIArticleServiceMockRecorder
}

// MockIArticleServiceMockRecorder is the mock recorder for MockIArticleService.
type MockIArticleServiceMockRecorder struct {
	mock *MockIArticleService
}

// NewMockIArticleService creates a new mock instance.
func NewMockIArticleService(ctrl *gomock.Controller) *MockIArticleService {
	mock := &MockIArticleService{ctrl: ctrl}
	mock.recorder = &MockIArticleServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIArticleService) EXPECT() *MockIArticleServiceMockRecorder {
	return m.recorder
}

// Fetch mocks base method.
func (m *MockIArticleService) Fetch(ctx context.Context, createdDate time.Time, num int) ([]models.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fetch", ctx, createdDate, num)
	ret0, _ := ret[0].([]models.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fetch indicates an expected call of Fetch.
func (mr *MockIArticleServiceMockRecorder) Fetch(ctx, createdDate, num interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fetch", reflect.TypeOf((*MockIArticleService)(nil).Fetch), ctx, createdDate, num)
}
