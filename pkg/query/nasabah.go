package query

const (
	NASABAH_INSERT = "INSERT INTO nasabah(nama, nik, no_hp, created_at, updated_at) VALUES($1,$2,$3,$4,$5) RETURNING id"
	NASABAH_COUNT  = "SELECT COUNT(*) FROM nasabah WHERE nik= $1 OR no_hp= $2"
)
