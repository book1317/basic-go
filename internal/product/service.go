package product

import (
	"academy/internal/app"
	"academy/internal/model"
	"context"
	"errors"
)

type Service struct {
	cv   *app.Config
	repo productRepo
}

func NewService(cv *app.Config, repo productRepo) *Service {
	return &Service{cv: cv, repo: repo}
}

type productRepo interface {
	GetProductsByMerchantId(ctx context.Context, mercahntId string) ([]model.Product, error)
	InsertProducts(ctx context.Context, product model.Product, merchantId string) error
}

func (s Service) GetListAllProductByMerchantId(ctx context.Context, merchantId string) ([]model.Product, error) {
	return s.repo.GetProductsByMerchantId(ctx, merchantId)
}

func (s Service) AddProductByMerchantId(ctx context.Context, product model.Product, merchantId string) error {
	products, err := s.repo.GetProductsByMerchantId(ctx, merchantId)
	if err != nil {
		return err
	}
	if len(products) >= 5 {
		return errors.New("maximum products is 5")
	}

	return s.repo.InsertProducts(ctx, product, merchantId)
}
