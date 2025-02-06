package rekening

import (
	"context"
	"database/sql"
	"fmt"
	"service-account/pkg/query"
	"time"
)

type RekeningRepository interface {
	Add(rekening *Rekening) error
	UpdateSaldo(noRekening *string, saldo *int, updateAt time.Time) (int, error)
	GetRekeningByNoRekening(noRekening *string) (*Rekening, error)
}
type rekeningRepositoryImpl struct {
	db  *sql.DB
	ctx context.Context
}

func NewRekeningRepository(db *sql.DB, ctx context.Context) RekeningRepository {
	return &rekeningRepositoryImpl{db: db, ctx: ctx}
}
func (r *rekeningRepositoryImpl) Add(rekening *Rekening) error {
	stmt, err := r.db.PrepareContext(r.ctx, query.REKENING_ADD)
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
func (r *rekeningRepositoryImpl) UpdateSaldo(noRekening *string, saldo *int, updateAt time.Time) (int, error) {
	var saldoNow int
	stmt, err := r.db.PrepareContext(r.ctx, query.REKENING_UPDATE_SALDO)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	err = stmt.QueryRowContext(r.ctx, &saldo, &updateAt, &noRekening).Scan(&saldoNow)
	if err != nil {
		return 0, err
	}
	return saldoNow, nil
}
func (r *rekeningRepositoryImpl) GetRekeningByNoRekening(noRekening *string) (*Rekening, error) {
	rekening := new(Rekening)
	stmt, err := r.db.PrepareContext(r.ctx, query.REKENING_GET_BY_NOREKENING)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.QueryContext(r.ctx, &noRekening)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&rekening.Id, &rekening.NasabahId, &rekening.NoRekening, &rekening.Saldo, &rekening.CreatedAt, &rekening.UpdatedAt)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("no_rekening %s not found", *noRekening)
	}
	return rekening, nil
}
