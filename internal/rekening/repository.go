package rekening

import (
	"context"
	"database/sql"
	"service-account/pkg/query"
)

type RekeningRepository interface {
	Add(rekening *Rekening) error
}
type impl struct {
	db  *sql.DB
	ctx context.Context
}

func NewRekeningRepository(db *sql.DB, ctx context.Context) RekeningRepository {
	return &impl{db: db, ctx: ctx}
}
func (r *impl) Add(rekening *Rekening) error {
	stmt, err := r.db.PrepareContext(r.ctx, query.REKENING_INSERT)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(r.ctx, &rekening.NasabahId, &rekening.NoRekening, &rekening.Saldo, &rekening.CreatedAt, &rekening.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}
