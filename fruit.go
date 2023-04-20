package main

import (
	"context"
	"fmt"
	"math/rand"
)

type Fruit struct {
	Name  string
	Color string
}

func (f *Fruit) Price(ctx context.Context) int {
	fmt.Printf("checking %s price\n", f.Name)
	return rand.Int()
}

func (f *Fruit) Availability(ctx context.Context) int {
	fmt.Printf("checking %s availability\n", f.Name)
	return rand.Int()
}
