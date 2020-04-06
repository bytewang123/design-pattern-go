package main

import (
	"fmt"
)

type houseBlend struct {
	description string
}

func newHouseBlend() *houseBlend {
	return &houseBlend{
		description: "houseBlend",
	}
}

func (h *houseBlend) cost() float64 {
	return 0.98
}

func (h *houseBlend) describe() {
	fmt.Printf("%+v ", h.description)
}

type darkRoast struct {
	description string
}

func newDarkRoast() *darkRoast {
	return &darkRoast{
		description: "darkRoast",
	}
}

func (d *darkRoast) cost() float64 {
	return 1.02
}

func (d *darkRoast) describe() {
	fmt.Println(d.description)
}
