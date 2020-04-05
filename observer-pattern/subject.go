package main

import "fmt"

type subject interface {
	registerObserver(string, observer)
	removeObserver(string)
	notifyObservers()
}

type weatherData struct {
	tem       float64
	hum       float64
	pre       float64
	observers map[string]observer
}

func (w *weatherData) init() {
	w.observers = map[string]observer{}
}

func (w *weatherData) registerObserver(name string, o observer) {
	w.observers[name] = o
}

func (w *weatherData) removeObserver(name string) {
	delete(w.observers, name)
}

func (w *weatherData) notifyObservers() {
	for k, v := range w.observers {
		fmt.Printf("notify observer %+v ...\n", k)
		v.update(w.tem, w.hum, w.pre)
		fmt.Printf("notify observer %+v success\n", k)
	}
}

func (w *weatherData) measurementsChanged() {
	w.notifyObservers()
}

func (w *weatherData) setMeasurements(tem, hum, pre float64) {
	w.tem = tem
	w.hum = hum
	w.pre = pre
	w.measurementsChanged()
}
