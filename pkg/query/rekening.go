package query

const (
	REKENING_ADD               = "INSERT INTO rekening(nasabah_id, no_rekening, saldo, created_at, updated_at) VALUES($1,$2,$3,$4,$5)"
	REKENING_UPDATE_SALDO      = "UPDATE rekening SET saldo = $1, updated_at = $2 WHERE no_rekening = $3 RETURNING saldo"
	REKENING_GET_BY_NOREKENING = "SELECT id, nasabah_id, no_rekening, saldo, created_at, updated_at FROM rekening WHERE no_rekening = $1"
)
