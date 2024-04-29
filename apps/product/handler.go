package product

import (
	"github.com/gofiber/fiber/v2"
	infrafiber "go-online-shop/infra/fiber"
	"go-online-shop/infra/response"
	"net/http"
)

type handler struct {
	svc service
}

func newHandler(svc service) handler {
	return handler{svc: svc}
}

func (h handler) CreateProduct(ctx *fiber.Ctx) error {
	var req = CreateProductRequestPayload{}

	if err := ctx.BodyParser(&req); err != nil {
		return infrafiber.NewResponse(
			infrafiber.WithMessage("invalid payload"),
			infrafiber.WithError(response.ErrorBadRequest),
		).Send(ctx)
	}

	if err := h.svc.CreateProduct(ctx.UserContext(), req); err != nil {
		myErr, ok := response.ErrorMapping[err.Error()]
		if !ok {
			myErr = response.ErrorGeneral
		}

		return infrafiber.NewResponse(
			infrafiber.WithMessage("invalid payload"),
			infrafiber.WithError(myErr),
		).Send(ctx)
	}

	return infrafiber.NewResponse(
		infrafiber.WithHttpCode(http.StatusCreated),
		infrafiber.WithMessage("product created successfully"),
	).Send(ctx)
}
