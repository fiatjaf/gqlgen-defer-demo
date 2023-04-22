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
}

func (f *Fruit) Price(ctx context.Context) int {
	return f.price
}

func (f *Fruit) Availability(ctx context.Context) Availability {
	time.Sleep(time.Second*1 + time.Duration(rand.Intn(2)))
	return Availability{}
}

type Availability struct{}

func (_ Availability) Here(ctx context.Context) int {
	time.Sleep(time.Duration(rand.Intn(1)+1) * time.Second)
	return rand.Intn(10)
}

func (_ Availability) There(ctx context.Context) int {
	time.Sleep(time.Duration(rand.Intn(2)+2) * time.Second)
	return rand.Intn(10)
}
