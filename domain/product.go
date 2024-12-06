package domain

import (
	"context"

	"github.com/Kamva/mgm/v2"
)

type Product struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string  `json:"name" bson:"name"`
	Price            float64 `json:"price" bson:"price"`
	Category         string  `json:"category" bson:"category"`
}

type ProductDto struct {
	Name     string  `json:"name,omitempty"`
	Price    float64 `json:"price,omitempty"`
	Category string  `json:"category,omitempty"`
}

type ProductRepository interface {
	Create(ctx context.Context, product *Product) (*Product, error)
	Fetch(ctx context.Context) ([]Product, error)
	GetById(ctx context.Context, id string) (*Product, error)
	GetByProductName(ctx context.Context, name string) (*Product, error)
}
