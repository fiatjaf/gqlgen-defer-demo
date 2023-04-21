package main

import (
	"context"
	"math/rand"
	"time"
)

type Fruit struct {
	Name  string
	Color string

	price int

	Availability Availability
}

func (f *Fruit) Price(ctx context.Context) int {
	return f.price
}

type Availability struct{}

func (_ Availability) Here(ctx context.Context) int {
	return rand.Intn(10)
}

func (_ Availability) There(ctx context.Context) int {
	time.Sleep(time.Duration(rand.Intn(5)+2) * time.Second)
	return rand.Intn(10)
}
