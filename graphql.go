package main

import (
	"context"
)

type Resolver struct{}

// // foo
func (r *queryResolver) Fruits(ctx context.Context) ([]*Fruit, error) {
	return []*Fruit{
		{Name: "banana", Color: "yellow"},
		{Name: "apple", Color: "red"},
	}, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
