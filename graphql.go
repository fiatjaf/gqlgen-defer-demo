package main

import (
	"context"
)

type Resolver struct{}

// // foo
func (r *queryResolver) Fruits(ctx context.Context) ([]*Fruit, error) {
	return []*Fruit{
		{Name: "banana", Color: "yellow", price: 1},
		{Name: "apple", Color: "red", price: 5},
	}, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
