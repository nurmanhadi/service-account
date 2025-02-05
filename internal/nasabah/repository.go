package nasabah

import (
	"context"
	"database/sql"
	"service-account/pkg/query"
)

type NasabahRepository interface {
	Add(nasabah *Nasabah) (int64, error)
	Count(nik *string, noHp *string) (int64, error)
}
type nasabahRepositoryImpl struct {
	db  *sql.DB
	ctx context.Context
}

func NewNasabahRepository(db *sql.DB, ctx context.Context) NasabahRepository {
	return &nasabahRepositoryImpl{db: db, ctx: ctx}
}
func (r *nasabahRepositoryImpl) Add(nasabah *Nasabah) (int64, error) {
	stmt, err := r.db.PrepareContext(r.ctx, query.NASABAH_INSERT)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	var nasabahId int64
	err = stmt.QueryRowContext(r.ctx, &nasabah.Nama, &nasabah.Nik, &nasabah.NoHp, &nasabah.CreatedAt, &nasabah.UpdatedAt).Scan(&nasabahId)
	if err != nil {
		return 0, err
	}
	return nasabahId, nil
}
func (r *nasabahRepositoryImpl) Count(nik *string, noHp *string) (int64, error) {
	var count int64
	stmt, err := r.db.PrepareContext(r.ctx, query.NASABAH_COUNT)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	rows, err := stmt.QueryContext(r.ctx, &nik, &noHp)
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			return 0, err
		}
	}
	return count, nil
}
