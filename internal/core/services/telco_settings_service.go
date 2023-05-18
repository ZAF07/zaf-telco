package services

import (
	"context"

	"github.com/ZAF07/telco/internal/core/ports"
)

type TelcoSettingsService struct {
	settingsRepo ports.ITelcoSettingRepository
}

// Implement port service interface to expose to the handlers adapter

func NewTelcoSettingsService(settingsRepo ports.ITelcoSettingRepository) ports.ITelcoSettingService {
	return &TelcoSettingsService{
		settingsRepo: settingsRepo,
	}
}

func (t *TelcoSettingsService) GetTelcoSetting(ctx context.Context, setting string) (string, error) {
	result, err := t.settingsRepo.GetTelcoSetting(ctx, setting)
	if err != nil {
		return "", err
	}
	return result, nil

}
