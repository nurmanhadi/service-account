package handler

import (
	"service-account/internal/rekening"
	"service-account/pkg/dto"
	"service-account/pkg/response"

	"github.com/gofiber/fiber/v2"
)

type RekeningHandlerImpl struct {
	rekeningService rekening.RekeningService
}

func NewRekeningHandler(rekeningService *rekening.RekeningService) *RekeningHandlerImpl {
	return &RekeningHandlerImpl{rekeningService: *rekeningService}
}
func (h *RekeningHandlerImpl) AddSaldo(c *fiber.Ctx) error {
	requestId, _ := c.Locals("request_id").(string)
	req := new(dto.RekeningRequestDto)
	if err := c.BodyParser(&req); err != nil {
		return response.ErrorResponse(c, 400, "cannot parse to json")
	}
	result, err := h.rekeningService.AddSaldo(req, &requestId)
	if err != nil {
		return response.ResponseError(c, err)
	}
	return response.Success(c, 200, result)
}
func (h *RekeningHandlerImpl) WithdrawSaldo(c *fiber.Ctx) error {
	requestId, _ := c.Locals("request_id").(string)
	req := new(dto.RekeningRequestDto)
	if err := c.BodyParser(&req); err != nil {
		return response.ErrorResponse(c, 400, "cannot parse to json")
	}
	result, err := h.rekeningService.WithdrawSaldo(req, &requestId)
	if err != nil {
		return response.ResponseError(c, err)
	}
	return response.Success(c, 200, result)
}
func (h *RekeningHandlerImpl) GetSaldoByNoRekening(c *fiber.Ctx) error {
	requestId, _ := c.Locals("request_id").(string)
	noRekening := c.Params("no_rekening")
	result, err := h.rekeningService.GetSaldoByNoRekening(&noRekening, &requestId)
	if err != nil {
		return response.ResponseError(c, err)
	}
	return response.Success(c, 200, result)
}
