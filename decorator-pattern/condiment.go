package main

import (
	"fmt"
)

type mocha struct {
	bev beverage
}

func newMocha(b beverage) *mocha {
	return &mocha{
		bev: b,
	}
}

func (m *mocha) cost() float64 {
	return 0.2 + m.bev.cost()
}

func (m *mocha) describe() {
	m.bev.describe()
	fmt.Println("mocha")
}

type whip struct {
	bev beverage
}

func newWhip(b beverage) *whip {
	return &whip{
		bev: b,
	}
}

func (w *whip) cost() float64 {
	return 0.2 + w.bev.cost()
}

func (w *whip) describe() {
	w.bev.describe()
	fmt.Println("whip")
}
