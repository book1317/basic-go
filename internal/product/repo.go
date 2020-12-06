package product

import (
	"academy/internal/app"
	"academy/internal/model"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Repo struct {
	cv *app.Config
}

func NewRepo(cv *app.Config) *Repo {
	return &Repo{cv: cv}
}

func (r Repo) GetProductsByMerchantId(ctx context.Context, mercahntId string) ([]model.Product, error) {
	var products []model.Product
	coll := r.cv.MongoDB.Client.Database("quiz").Collection("product")
	merchantObjectId, err := primitive.ObjectIDFromHex(mercahntId)
	if err != nil {
		return nil, err
	}

	cur, err := coll.Find(ctx, bson.M{"merchant_id": merchantObjectId})
	if err != nil {
		return nil, err
	}

	if err := cur.All(ctx, &products); err != nil {
		return nil, err
	}
	return products, nil
}

func (r Repo) InsertProducts(ctx context.Context, product model.Product, merchantId string) error {
	merchantObjectId, err := primitive.ObjectIDFromHex(merchantId)
	if err != nil {
		return err
	}

	coll := r.cv.MongoDB.Client.Database("quiz").Collection("product")
	_, err = coll.InsertOne(ctx, bson.M{"name": product.Name, "amount": product.Amount, "merchant_id": merchantObjectId, "stocks": product.Stocks})
	if err != nil {
		return err
	}
	return nil
}
