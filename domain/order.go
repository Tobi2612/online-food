package domain

import (
	"context"

	"github.com/Kamva/mgm/v2"
)

type Item struct {
	ProductID string `json:"productId" bson:"productId"`
	Quantity  int    `json:"quantity" bson:"quantity"`
}

type Order struct {
	mgm.DefaultModel `bson:",inline"`
	CouponCode       string `json:"couponCode" bson:"couponCode"`
	Items            []Item `json:"items" bson:"items"`
}

type OrderReqDTO struct {
	CouponCode string `json:"couponCode,omitempty"`
	Items      []Item `json:"items,omitempty"`
}

type OrderRepository interface {
	Create(ctx context.Context, order *Order) (*Order, error)
	Fetch(ctx context.Context) ([]Order, error)
	GetById(ctx context.Context, id string) (*Order, error)
}
