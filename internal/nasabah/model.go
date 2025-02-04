package nasabah

import "time"

type Nasabah struct {
	Id        int       `json:"id"`
	Nama      string    `json:"nama"`
	Nik       string    `json:"nik"`
	NoHp      int       `json:"no_hp"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
