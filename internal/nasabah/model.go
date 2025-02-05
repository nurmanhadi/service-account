package nasabah

import (
	"service-account/internal/rekening"
	"time"
)

type Nasabah struct {
	Id        int               `json:"id"`
	Nama      string            `json:"nama"`
	Nik       string            `json:"nik"`
	NoHp      string            `json:"no_hp"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
	Rekening  rekening.Rekening `json:"rekening"`
}
