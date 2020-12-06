package product

import (
	"academy/internal/app"
	"academy/internal/model"
	"context"
	"net/http"

	"github.com/labstack/echo"
)

type Endpoint struct {
	cv  *app.Config
	srv productService
}

func NewEndpoint(cv *app.Config, srv productService) *Endpoint {
	return &Endpoint{cv: cv, srv: srv}
}

type productService interface {
	GetListAllProductByMerchantId(ctx context.Context, merchantId string) ([]model.Product, error)
	AddProductByMerchantId(ctx context.Context, product model.Product, merchantId string) error
}

func (e Endpoint) GetListAllProduct(c echo.Context) error {
	var merchant model.Merchant
	if err := c.Bind(&merchant); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	merchantId := c.Param("id")
	products, err := e.srv.GetListAllProductByMerchantId(c.Request().Context(), merchantId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, products)
}

func (e Endpoint) AddProduct(c echo.Context) error {
	var product model.Product
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	merchantId := c.Param("id")
	err := e.srv.AddProductByMerchantId(c.Request().Context(), product, merchantId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, "add product success")
	return nil
}
