package rekening

import (
	"service-account/pkg/dto"
	"service-account/pkg/exception"
	"service-account/pkg/validation"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type RekeningService interface {
	AddSaldo(req *dto.RekeningRequestDto, requestId *string) (map[string]interface{}, error)
	WithdrawSaldo(req *dto.RekeningRequestDto, requestId *string) (map[string]interface{}, error)
	GetSaldoByNoRekening(noRekening *string, requestId *string) (map[string]interface{}, error)
}
type rekeningServiceImpl struct {
	rekeningRepository RekeningRepository
	validation         *validator.Validate
	logger             *logrus.Logger
}

func NewRekeningService(rekeningRepository *RekeningRepository, validation *validator.Validate, logger *logrus.Logger) RekeningService {
	return &rekeningServiceImpl{rekeningRepository: *rekeningRepository, validation: validation, logger: logger}
}
func (s *rekeningServiceImpl) AddSaldo(req *dto.RekeningRequestDto, requestId *string) (map[string]interface{}, error) {
	if err := s.validation.Struct(req); err != nil {
		s.logger.WithFields(logrus.Fields{
			"request_id": requestId,
			"data":       req,
			"task":       "validation",
			"error":      validation.ValidationError(err),
		}).Error("validation error")
		return nil, err
	}
	rekening, err := s.rekeningRepository.GetRekeningByNoRekening(&req.NoRekening)
	if err != nil {
		s.logger.WithFields(logrus.Fields{
			"request_id": requestId,
			"data":       req.NoRekening,
			"task":       "get rekening by no_rekening",
			"error":      exception.RekeningNotFound,
		}).Error("not found")
		return nil, exception.RekeningNotFound
	}
	newSaldo := rekening.Saldo + req.Saldo
	saldo, err := s.rekeningRepository.UpdateSaldo(&req.NoRekening, &newSaldo, time.Now())
	if err != nil {
		s.logger.WithFields(logrus.Fields{
			"request_id": requestId,
			"data":       req,
			"task":       "update saldo to database",
			"error":      err.Error(),
		}).Error("databse error")
		return nil, err
	}
	return map[string]interface{}{"saldo": saldo}, nil
}
func (s *rekeningServiceImpl) WithdrawSaldo(req *dto.RekeningRequestDto, requestId *string) (map[string]interface{}, error) {
	if err := s.validation.Struct(req); err != nil {
		s.logger.WithFields(logrus.Fields{
			"request_id": requestId,
			"data":       req,
			"task":       "validation",
			"error":      validation.ValidationError(err),
		}).Error("validation error")
		return nil, err
	}
	rekening, err := s.rekeningRepository.GetRekeningByNoRekening(&req.NoRekening)
	if err != nil {
		s.logger.WithFields(logrus.Fields{
			"request_id": requestId,
			"data":       req.NoRekening,
			"task":       "get rekening by no_rekening",
			"error":      exception.RekeningNotFound,
		}).Error("not found")
		return nil, exception.RekeningNotFound
	}
	newSaldo := rekening.Saldo - req.Saldo
	if newSaldo < 0 {
		s.logger.WithFields(logrus.Fields{
			"request_id": requestId,
			"data":       newSaldo,
			"task":       "validatation nominal saldo",
			"error":      exception.RekeningSaldoLessThan,
		}).Error("nominal saldo")
		return nil, exception.RekeningSaldoLessThan
	}
	saldo, err := s.rekeningRepository.UpdateSaldo(&req.NoRekening, &newSaldo, time.Now())
	if err != nil {
		s.logger.WithFields(logrus.Fields{
			"request_id": requestId,
			"data":       req,
			"task":       "update saldo to database",
			"error":      err.Error(),
		}).Error("databse error")
		return nil, err
	}
	return map[string]interface{}{"saldo": saldo}, nil
}
func (s *rekeningServiceImpl) GetSaldoByNoRekening(noRekening *string, requestId *string) (map[string]interface{}, error) {
	rekening, err := s.rekeningRepository.GetRekeningByNoRekening(noRekening)
	if err != nil {
		s.logger.WithFields(logrus.Fields{
			"request_id": requestId,
			"data":       noRekening,
			"task":       "get rekening by no_rekening",
			"error":      exception.RekeningNotFound,
		}).Error("not found")
		return nil, exception.RekeningNotFound
	}
	return map[string]interface{}{"saldo": rekening.Saldo}, nil
}
