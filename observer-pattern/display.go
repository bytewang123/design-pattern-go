package main

import "fmt"

type displayElement interface {
	display()
}

type currentConditionDisplay struct {
	tem float64
	hum float64
	sub subject
}

func (c *currentConditionDisplay) display() {
	fmt.Printf("current condition display: temprature:%+v, humidity:%+v\n", c.tem, c.hum)
}

func (c *currentConditionDisplay) update(tem, hum, sub float64) {
	c.tem = tem
	c.hum = hum
	c.display()
}

type forecastDisplay struct {
	tem float64
	hum float64
	sub subject
}

func (f *forecastDisplay) display() {
	fmt.Printf("forecast display: temprature:%+v, humidity:%+v\n", f.tem, f.hum)
}

func (f *forecastDisplay) update(tem, hum, sub float64) {
	f.tem = tem
	f.hum = hum
	f.display()
}
