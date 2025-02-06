package exception

import "errors"

var (
	RekeningNotFound      = errors.New("no rekening not found")
	RekeningSaldoLessThan = errors.New("saldo less than nominal")
)
