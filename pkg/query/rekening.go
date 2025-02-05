package query

const (
	REKENING_INSERT = "INSERT INTO rekening(nasabah_id, no_rekening, saldo, created_at, updated_at) VALUES($1,$2,$3,$4,$5)"
)
