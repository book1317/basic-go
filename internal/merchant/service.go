package merchant

import (
	"context"

	"academy/internal/app"
	"academy/internal/model"
	"math/rand"
	"time"
)

type Service struct {
	cv   *app.Config
	repo merchantRepo
}

func NewService(cv *app.Config, repo merchantRepo) *Service {
	return &Service{cv: cv, repo: repo}
}

type merchantRepo interface {
	InserMerchant(ctx context.Context, mercahnt model.Merchant) (model.Merchant, error)
	GetMerchantByBankAccount(ctx context.Context, mercahnt model.Merchant) error
	GetMerchantById(ctx context.Context, mercahntId string) (model.Merchant, error)
	UpdateMerchant(ctx context.Context, mercahnt model.Merchant) error
}

func (s Service) RegisterMerchant(ctx context.Context, merchant model.Merchant) (model.Merchant, error) {
	if err := s.repo.GetMerchantByBankAccount(ctx, merchant); err != nil {
		return merchant, err
	}

	const usernameCharSet = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const passwordCharSet = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	merchant.Username = StringWithCharset(10, usernameCharSet)
	merchant.Password = StringWithCharset(20, passwordCharSet)
	if _, err := s.repo.InserMerchant(ctx, merchant); err != nil {
		return merchant, err
	}
	return merchant, nil
}

func (s Service) GetMerchantById(ctx context.Context, merchantId string) (model.Merchant, error) {
	merchant, err := s.repo.GetMerchantById(ctx, merchantId)
	if err != nil {
		return merchant, err
	}
	return merchant, nil
}

func StringWithCharset(length int, charset string) string {
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func (s Service) UpdateMerchant(ctx context.Context, merchant model.Merchant) error {
	err := s.repo.UpdateMerchant(ctx, merchant)
	if err != nil {
		return err
	}
	return nil
}
