package mockdb

import (
	"context"

	"github.com/ZAF07/telco/internal/core/domain"
)

type MockTelcoAccountRepoAdapter struct{}

func NewMockTelcoAccountRepoAdapter() *MockTelcoAccountRepoAdapter {
	return &MockTelcoAccountRepoAdapter{}
}

func (t *MockTelcoAccountRepoAdapter) CreateTelcoAccount(ctx context.Context, user domain.TelcoAccount) (bool, error) {
	return true, nil
}
