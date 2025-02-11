package handler

import (
	"service-account/internal/nasabah"
	"service-account/pkg/dto"
	"service-account/pkg/response"

	"github.com/gofiber/fiber/v2"
)

type NasabahHandlerImpl struct {
	nasabahService nasabah.NasabahService
}

func NewNasabahHandler(nasabahService *nasabah.NasabahService) *NasabahHandlerImpl {
	return &NasabahHandlerImpl{nasabahService: *nasabahService}
}
func (h *NasabahHandlerImpl) Daftar(c *fiber.Ctx) error {
	requestId, _ := c.Locals("request_id").(string)
	req := new(dto.NasabahRequestDto)
	if err := c.BodyParser(&req); err != nil {
		return response.ErrorResponse(c, 400, "cannot parse to json")
	}
	result, err := h.nasabahService.Daftar(req, &requestId)
	if err != nil {
		return response.ResponseError(c, err)
	}
	return response.Success(c, 200, result)
}
