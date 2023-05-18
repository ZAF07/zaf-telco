package postgresdb

import (
	"context"
	"database/sql"

	"github.com/ZAF07/telco/internal/core/ports"
	_ "github.com/lib/pq"
)

//  The Repo Adapter is the concrete implementation of the ITelcoReposiroty interface. They are the ones who talk to the outside world (the database in this case) and is also isolated from the internal services via abstracting with an interface as well. You can freely change the implementation here (stuff like connecting to a diff database or what library you use to interact with the database) as long as this package implements the Irepository interface

type TelcoSettingRepoAdapter struct {
	db    *sql.DB
	table string
}

// Verify that TelcoSettingRepoAdapter implements ITelcoSettingRepository interface
var _ ports.ITelcoSettingRepository = (*TelcoSettingRepoAdapter)(nil)

func NewTelcoSettingRepoAdapter(db *sql.DB, table string) *TelcoSettingRepoAdapter {
	return &TelcoSettingRepoAdapter{
		db:    db,
		table: table,
	}
}

func (t *TelcoSettingRepoAdapter) GetTelcoSetting(ctx context.Context, setting string) (string, error) {
	return "This was returned from the TelcoSetingRepoAdapter", nil
}
