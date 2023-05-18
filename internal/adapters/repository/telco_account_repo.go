package repository

import (
	"context"
	"database/sql"

	"github.com/ZAF07/telco/internal/core/domain"
)

type TelcoAccountRepo struct {
	db    *sql.DB
	table string
}

func NewTelcoAccountRepo(db interface{}, table string) *TelcoAccountRepo {
	return &TelcoAccountRepo{
		db:    db.(*sql.DB),
		table: table,
	}
}

func (ta *TelcoAccountRepo) CreateTelcoAccount(ctx context.Context, user domain.TelcoAccount) (bool, error) {
	return true, nil
}
