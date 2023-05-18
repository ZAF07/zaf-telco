package postgresdb

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/ZAF07/telco/internal/core/domain"
	"github.com/ZAF07/telco/internal/core/ports"
)

type TelcoAccountRepoAdapter struct {
	db    *sql.DB
	table string
}

// Verify that TelcoAccountRepoAdapter implements ITelcoAccountRepository interface
var _ ports.ITelcoAccountRepository = (*TelcoAccountRepoAdapter)(nil)

func NewTelcoAccountRepoAdapter(db *sql.DB, table string) *TelcoAccountRepoAdapter {
	return &TelcoAccountRepoAdapter{
		db:    db,
		table: table,
	}
}

func (ta *TelcoAccountRepoAdapter) CreateTelcoAccount(ctx context.Context, user domain.TelcoAccount) (bool, error) {
	return true, nil
}
