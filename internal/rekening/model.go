package rekening

import "time"

type Rekening struct {
	Id         int       `json:"id"`
	NasabahId  int       `json:"nasabah_id"`
	NoRekening string    `json:"no_rekening"`
	Saldo      int       `json:"saldo"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
