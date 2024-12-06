package mongodb

import (
	"context"
	"online-food/domain"

	"github.com/Kamva/mgm/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type mongoOrderRepository struct {
	Logger *zap.Logger
	Coll   *mgm.Collection
}

func (m mongoOrderRepository) Create(ctx context.Context, order *domain.Order) (*domain.Order, error) {
	err := m.Coll.Create(order)

	if err != nil {
		m.Logger.Error(err.Error(), zap.Error(err))
		return nil, err
	}

	return order, nil
}

func (m mongoOrderRepository) Fetch(ctx context.Context) ([]domain.Order, error) {
	var orders []domain.Order

	if err := m.Coll.SimpleFind(&orders, bson.D{}, &options.FindOptions{Sort: bson.D{{Key: "created_at", Value: -1}}}); err != nil {
		return nil, err
	}
	return orders, nil
}

func (m mongoOrderRepository) GetById(ctx context.Context, id string) (*domain.Order, error) {
	var order domain.Order
	err := m.Coll.FindByID(id, &order)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func NewOrderRepository(logger *zap.Logger) domain.OrderRepository {
	return &mongoOrderRepository{
		Logger: logger,
		Coll:   mgm.Coll(&domain.Order{}),
	}
}
