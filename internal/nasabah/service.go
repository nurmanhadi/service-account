package nasabah

import (
	"fmt"
	"math/rand/v2"
	"service-account/internal/rekening"
	"service-account/pkg/dto"
	"service-account/pkg/exception"
	"service-account/pkg/validation"
	"service-account/utils"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type NasabahService interface {
	Daftar(req *dto.NasabahRequestDto) (map[string]interface{}, error)
}
type nasabahServiceImpl struct {
	nasabahRepository  NasabahRepository
	rekeningRepository rekening.RekeningRepository
	validation         *validator.Validate
	logger             *logrus.Logger
}

func NewNasabahService(nasabahRepository *NasabahRepository, rekeningRepository *rekening.RekeningRepository, validation *validator.Validate, logger *logrus.Logger) NasabahService {
	return &nasabahServiceImpl{nasabahRepository: *nasabahRepository, rekeningRepository: *rekeningRepository, validation: validation, logger: logger}
}
func (s *nasabahServiceImpl) Daftar(req *dto.NasabahRequestDto) (map[string]interface{}, error) {
	if err := s.validation.Struct(req); err != nil {
		s.logger.WithFields(logrus.Fields{
			"id":    rand.Int64N(99999999999),
			"data":  req,
			"task":  "validation",
			"error": validation.ValidationError(err),
		}).Error("validation error")
		return nil, err
	}
	countNasabah, err := s.nasabahRepository.Count(&req.Nik, &req.NoHp)
	if err != nil {
		s.logger.WithFields(logrus.Fields{
			"id":    rand.Int64N(99999999999),
			"data":  fmt.Sprintf("%s, %s", req.Nik, req.NoHp),
			"task":  "count nasabah to database",
			"error": err,
		}).Fatal("count error")
		return nil, err
	}
	if countNasabah == 1 {
		s.logger.WithFields(logrus.Fields{
			"id":    rand.Int64N(99999999999),
			"data":  fmt.Sprintf("%s, %s", req.Nik, req.NoHp),
			"task":  "count nasabah",
			"error": "nik atau no hp sudah terdaftar",
		}).Error("exist")
		return nil, exception.NasabahCountAlreadyExist
	}
	nasabah := &Nasabah{
		Nama:      req.Nama,
		Nik:       req.Nik,
		NoHp:      req.NoHp,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	nasabahId, err := s.nasabahRepository.Add(nasabah)
	if err != nil {
		s.logger.WithFields(logrus.Fields{
			"id":    rand.Int64N(99999999999),
			"data":  nasabah,
			"task":  "add nasabah to database",
			"error": err,
		}).Fatal("database error")
		return nil, err
	}
	noRekening := utils.GenerateRandomRekening()
	rekening := &rekening.Rekening{
		NasabahId:  int(nasabahId),
		NoRekening: noRekening,
		Saldo:      0,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	if err := s.rekeningRepository.Add(rekening); err != nil {
		s.logger.WithFields(logrus.Fields{
			"id":    rand.Int64N(99999999999),
			"data":  rekening,
			"task":  "add rekening to database",
			"error": err,
		}).Fatal("database error")
		return nil, err
	}
	return map[string]interface{}{"no_rekening": noRekening}, nil
}
