package mongodb

import (
	"context"
	"online-food/domain"

	"github.com/Kamva/mgm/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type mongoProductRepository struct {
	Logger *zap.Logger
	Coll   *mgm.Collection
}

func (m mongoProductRepository) Create(ctx context.Context, product *domain.Product) (*domain.Product, error) {
	err := m.Coll.Create(product)

	if err != nil {
		m.Logger.Error(err.Error(), zap.Error(err))
		return nil, err
	}

	return product, nil
}

func (m mongoProductRepository) Fetch(ctx context.Context) ([]domain.Product, error) {
	var products []domain.Product

	if err := m.Coll.SimpleFind(&products, bson.D{}, &options.FindOptions{Sort: bson.D{{Key: "created_at", Value: -1}}}); err != nil {
		return nil, err
	}
	return products, nil
}

func (m mongoProductRepository) GetById(ctx context.Context, id string) (*domain.Product, error) {
	var product domain.Product
	err := m.Coll.FindByID(id, &product)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (m mongoProductRepository) GetByProductName(ctx context.Context, name string) (*domain.Product, error) {
	var product []domain.Product
	err := m.Coll.SimpleFind(&product, bson.D{{Key: "name", Value: name}})
	if err != nil {
		return nil, err
	}

	return &product[0], nil
}

func NewProductRepository(logger *zap.Logger) domain.ProductRepository {
	return &mongoProductRepository{
		Logger: logger,
		Coll:   mgm.Coll(&domain.Product{}),
	}
}
