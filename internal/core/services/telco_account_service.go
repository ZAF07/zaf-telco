package services

import (
	"context"

	"github.com/ZAF07/telco/internal/core/domain"
	"github.com/ZAF07/telco/internal/core/ports"
)

type TelcoAccountService struct {
	accountRepo ports.ITelcoAccountRepository
}

// Verify that TelcoAccountService implements ITelcoAccountService interface
var _ ports.ITelcoAccountService = (*TelcoAccountService)(nil)

// Constructor to create a new TelcoAccountService instance
func NewTelcoAccountService(accountRepo ports.ITelcoAccountRepository) *TelcoAccountService {
	return &TelcoAccountService{
		accountRepo: accountRepo,
	}
}

// 💡Telco Account domain specific business logic lives here

func (t *TelcoAccountService) CreateTelcoAccount(ctx context.Context, account domain.TelcoAccount) (bool, error) {
	success, err := t.accountRepo.CreateTelcoAccount(ctx, account)
	if err != nil {
		return false, err
	}

	return success, nil
}
