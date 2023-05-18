package mockdb

import "context"

type MockTelcoSettingRepoAdapter struct{}

func NewMockTelcoSettingRepoAdapter() *MockTelcoSettingRepoAdapter {
	return &MockTelcoSettingRepoAdapter{}
}

func (m *MockTelcoSettingRepoAdapter) GetTelcoSetting(ctx context.Context, setting string) (string, error) {
	return "This is from the mock telco setting repo adapter", nil
}
