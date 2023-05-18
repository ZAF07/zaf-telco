package ports

import (
	"context"

	"github.com/ZAF07/telco/internal/core/domain"
)

//  This module hosts the interface/contracts which define what the core service can do

type ITelcoAccountService interface {
	// GetTelcoAccountName(ctx context.Context, id int) (string, error)
	// GetTelcoAccountStatus() string
	// GetTelcoAccountSIMNumber() int
	// GetTelcoAccountTelcoStatus() bool
	CreateTelcoAccount(ctx context.Context, user domain.TelcoAccount) (bool, error)
	// UpdateTelcoAccount(user domain.TelcoAccount) bool
}

type ITelcoSettingService interface {
	GetTelcoSetting(ctx context.Context, setting string) (string, error)
}
