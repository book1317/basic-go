package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"academy/internal/app"
	"academy/internal/merchant"
	"academy/internal/product"
)

type route struct {
	Desc           string
	Group          string
	Path           string
	HttpMethod     string
	HandlerFunc    echo.HandlerFunc
	MiddlewareFunc []echo.MiddlewareFunc
}

func NewRouter(e *echo.Echo, c *app.Config) error {
	merchantRepo := merchant.NewRepo(c)
	merchantService := merchant.NewService(c, merchantRepo)
	merchantEndpoint := merchant.NewEndpoint(c, merchantService)

	productRepo := product.NewRepo(c)
	productService := product.NewService(c, productRepo)
	productEndpoint := product.NewEndpoint(c, productService)

	routes := []route{
		{
			Desc:        "Register Merchant",
			Group:       "merchant",
			Path:        "/register",
			HttpMethod:  http.MethodPost,
			HandlerFunc: merchantEndpoint.RegisterMerchant,
		},
		{
			Desc:           "Merchant Information",
			Group:          "merchant",
			Path:           "/information/:id",
			HttpMethod:     http.MethodGet,
			HandlerFunc:    merchantEndpoint.GetMerchantInformation,
			MiddlewareFunc: []echo.MiddlewareFunc{verifyPassword(c)},
		},
		{
			Desc:           "Update Merchant",
			Group:          "merchant",
			Path:           "/update",
			HttpMethod:     http.MethodPost,
			HandlerFunc:    merchantEndpoint.UpdateMerchant,
			MiddlewareFunc: []echo.MiddlewareFunc{verifyPassword(c)},
		},
		{
			Desc:           "List All Products",
			Group:          "merchant/:id",
			Path:           "/products",
			HttpMethod:     http.MethodGet,
			HandlerFunc:    productEndpoint.GetListAllProduct,
			MiddlewareFunc: []echo.MiddlewareFunc{verifyPassword(c)},
		},
		{
			Desc:        "Add Product",
			Group:       "merchant/:id",
			Path:        "/product",
			HttpMethod:  http.MethodPost,
			HandlerFunc: productEndpoint.AddProduct,
		},
	}

	// middleware
	e.Use(middleware.BodyDumpWithConfig(bodyDumpConfig()))

	for _, r := range routes {
		e.Group(r.Group).Add(r.HttpMethod, r.Path, r.HandlerFunc, r.MiddlewareFunc...)
	}

	return nil
}
