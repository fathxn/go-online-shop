package auth

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
	return handler{
		svc: svc,
	}
}

func (h handler) register(ctx *fiber.Ctx) error {
	var req = RegisterRequestPayload{}
	if err := ctx.BodyParser(&req); err != nil {
		myErr := response.ErrorBadRequest
		return infrafiber.NewResponse(
			infrafiber.WithMessage(err.Error()),
			infrafiber.WithError(myErr),
			infrafiber.WithHttpCode(http.StatusBadRequest),
		).Send(ctx)
	}

	if err := h.svc.register(ctx.UserContext(), req); err != nil {
		myErr, ok := response.ErrorMapping[err.Error()]
		if !ok {
			myErr = response.ErrorGeneral
		}

		return infrafiber.NewResponse(
			infrafiber.WithMessage(err.Error()),
			infrafiber.WithError(myErr),
			infrafiber.WithHttpCode(http.StatusBadRequest),
		).Send(ctx)
	}

	return infrafiber.NewResponse(
		infrafiber.WithHttpCode(http.StatusOK),
	).Send(ctx)
}

func (h handler) login(ctx *fiber.Ctx) error {
	var req = LoginRequestPayload{}
	if err := ctx.BodyParser(&req); err != nil {
		myErr := response.ErrorBadRequest
		return infrafiber.NewResponse(
			infrafiber.WithMessage(err.Error()),
			infrafiber.WithError(myErr),
			infrafiber.WithHttpCode(http.StatusBadRequest),
		).Send(ctx)
	}

	token, err := h.svc.login(ctx.UserContext(), req)

	if err != nil {
		myErr, ok := response.ErrorMapping[err.Error()]
		if !ok {
			myErr = response.ErrorGeneral
		}

		return infrafiber.NewResponse(
			infrafiber.WithMessage(err.Error()),
			infrafiber.WithError(myErr),
			infrafiber.WithHttpCode(http.StatusBadRequest),
		).Send(ctx)
	}

	return infrafiber.NewResponse(
		infrafiber.WithHttpCode(http.StatusOK),
		infrafiber.WithPayload(map[string]interface{}{
			"access_token": token,
		}),
	).Send(ctx)
}
