package merchant

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"academy/internal/app"
	"academy/internal/model"
	"errors"
)

type Repo struct {
	cv *app.Config
}

func NewRepo(cv *app.Config) *Repo {
	return &Repo{cv: cv}
}

func (r Repo) InserMerchant(ctx context.Context, merchant model.Merchant) (model.Merchant, error) {
	coll := r.cv.MongoDB.Client.Database("quiz").Collection("merchant")
	_, err := coll.InsertOne(ctx, merchant)
	if err != nil {
		return merchant, err
	}

	return merchant, nil
}

func (r Repo) GetMerchantByBankAccount(ctx context.Context, merchant model.Merchant) error {
	var tempMerchant model.Merchant
	coll := r.cv.MongoDB.Client.Database("quiz").Collection("merchant")
	coll.FindOne(ctx, bson.M{"bank_account": merchant.BankAccount}).Decode(&tempMerchant)
	if (tempMerchant != model.Merchant{}) {
		return errors.New("bank account already exist")
	}
	return nil
}

func (r Repo) GetMerchantById(ctx context.Context, merchantId string) (model.Merchant, error) {
	var merchant model.Merchant
	coll := r.cv.MongoDB.Client.Database("quiz").Collection("merchant")
	merchantObjectId, err := primitive.ObjectIDFromHex(merchantId)
	if err != nil {
		return merchant, err
	}

	coll.FindOne(ctx, bson.M{"_id": merchantObjectId}).Decode(&merchant)
	if (merchant == model.Merchant{}) {
		return merchant, errors.New("no merchant")
	}
	return merchant, nil
}

func (r Repo) UpdateMerchant(ctx context.Context, merchant model.Merchant) error {
	coll := r.cv.MongoDB.Client.Database("quiz").Collection("merchant")
	_, err := coll.UpdateOne(ctx, bson.M{"username": merchant.Username}, bson.M{"$set": bson.M{"name": merchant.Name}})
	if err != nil {
		return err
	}
	return nil
}
