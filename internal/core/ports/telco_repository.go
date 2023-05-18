package ports

import (
	"context"

	"github.com/ZAF07/telco/internal/core/domain"
)

//  This module hosts the interface/contracts which define what the core service can do

type ITelcoRepository interface {
	GetTelcoStatus(id string)
	// CreateAccount(account domain.TelcoAccount)
	// UpdateSetting(setting string)
}

// Interface/Ports for any operation to the Telco Account Database
type ITelcoAccountRepository interface {
	// GetTelcoAccountName() string
	// GetTelcoAccountStatus() string
	// GetTelcoAccountSIMNumber() int
	// GetTelcoAccountTelcoStatus() bool
	CreateTelcoAccount(ctx context.Context, user domain.TelcoAccount) (bool, error)
	// UpdateTelcoAccount(user domain.TelcoAccount) bool
}

// Interface/Ports for any operation to the Telco Settings Database
type ITelcoSettingRepository interface {
	GetTelcoSetting(ctx context.Context, setting string) (string, error)
	// GetTelcoAccountSettingStatus(settingType string) string
	// UpdateTelcoSetting(settingType, value string) bool
}
