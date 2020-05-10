package service

import (
	"encoding/json"
	"fmt"
	"github.com/aclk/goblog/common/model"
	"github.com/aclk/goblog/common/tracing"
	"github.com/aclk/goblog/dataservice/cmd"
	"github.com/aclk/goblog/dataservice/internal/app/dbclient"
	"github.com/aclk/goblog/dataservice/internal/app/dbclient/mock_dbclient"
	"github.com/golang/mock/gomock"
	"github.com/opentracing/opentracing-go"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

var serviceName = "dataservice"

// Run this first in each test, poor substitute for a proper @Before func
func setup(mockRepo dbclient.IGormClient) *Server {
	tracing.SetTracer(opentracing.NoopTracer{})
	h := NewHandler(mockRepo)
	s := NewServer(h, cmd.DefaultConfiguration())
	s.SetupRoutes()
	return s
}

func TestGetAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mock_dbclient.NewMockIGormClient(ctrl)
	mockRepo.EXPECT().QueryAccount(gomock.Any(), "123").Return(model.AccountData{ID: "123", Name: "Person_123"}, nil)

	s := setup(mockRepo)

	req := httptest.NewRequest("GET", "/accounts/123", nil)
	resp := httptest.NewRecorder()

	s.r.ServeHTTP(resp, req)

	account := model.AccountData{}
	_ = json.Unmarshal(resp.Body.Bytes(), &account)

	assert.Equal(t, 200, resp.Code)
	assert.Equal(t, "123", account.ID)
	assert.Equal(t, "Person_123", account.Name)
}

func TestGetNonExistingAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mock_dbclient.NewMockIGormClient(ctrl)
	mockRepo.EXPECT().QueryAccount(gomock.Any(), "456").Return(model.AccountData{}, fmt.Errorf(""))

	s := setup(mockRepo)
	req := httptest.NewRequest("GET", "/accounts/456", nil)
	resp := httptest.NewRecorder()

	s.r.ServeHTTP(resp, req)

	assert.Equal(t, 404, resp.Code)
}

func TestHealth(t *testing.T) {
	s := setup(nil)

	req := httptest.NewRequest("GET", "/health", nil)
	resp := httptest.NewRecorder()

	s.r.ServeHTTP(resp, req)
	assert.Equal(t, 200, resp.Code)
}

func TestGetAccountWrongPath(t *testing.T) {
	s := setup(nil)

	req := httptest.NewRequest("GET", "/invalid/123", nil)
	resp := httptest.NewRecorder()

	s.r.ServeHTTP(resp, req)

	assert.Equal(t, 404, resp.Code)
}
