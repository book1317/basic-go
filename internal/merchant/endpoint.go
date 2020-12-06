package merchant

import (
	"context"
	"net/http"

	"github.com/labstack/echo"

	"academy/internal/app"
	"academy/internal/model"
)

type Endpoint struct {
	cv  *app.Config
	srv merchantService
}

func NewEndpoint(cv *app.Config, srv merchantService) *Endpoint {
	return &Endpoint{cv: cv, srv: srv}
}

type merchantService interface {
	RegisterMerchant(ctx context.Context, merchant model.Merchant) (model.Merchant, error)
	GetMerchantById(ctx context.Context, merchantId string) (model.Merchant, error)
	UpdateMerchant(ctx context.Context, merchant model.Merchant) error
}

func (e Endpoint) RegisterMerchant(c echo.Context) error {
	var merchant model.Merchant
	if err := c.Bind(&merchant); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	merchant, err := e.srv.RegisterMerchant(c.Request().Context(), merchant)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, merchant)
}

func (e Endpoint) GetMerchantInformation(c echo.Context) error {
	var merchant model.Merchant
	if err := c.Bind(&merchant); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	merchantId := c.Param("id")
	merchant, err := e.srv.GetMerchantById(c.Request().Context(), merchantId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, merchant)
}

func (e Endpoint) UpdateMerchant(c echo.Context) error {
	var merchant model.Merchant
	if err := c.Bind(&merchant); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err := e.srv.UpdateMerchant(c.Request().Context(), merchant)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, "update complete")
}
