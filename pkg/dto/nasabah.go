package dto

type NasabahRequestDto struct {
	Nama string `json:"nama" validate:"required,max=100"`
	Nik  string `json:"nik" validate:"required,min=10,max=20"`
	NoHp string `json:"no_hp" validate:"required,min=9,max=20"`
}
