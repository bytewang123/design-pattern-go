package main

import "fmt"

type basicDuck interface {
	quackBehaviour
	flyBehaviour
	display()
	swim()
}

type duck struct {
	quackBehaviour
	flyBehaviour
	name string
}

func (d *duck) display() {
	fmt.Printf("I am %+v\n", d.name)
}

func (d *duck) swim() {
	fmt.Printf("I am %+v, I can swim\n", d.name)
}

type rocketDuck struct {
	*duck
}

type toyDuck struct {
}
