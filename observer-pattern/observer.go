package main

type observer interface {
	update(float64, float64, float64)
}
