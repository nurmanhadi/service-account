package dto

type RekeningRequestDto struct {
	NoRekening string `json:"no_rekening" validate:"required,max=20"`
	Saldo      int    `json:"saldo" validate:"required"`
}
