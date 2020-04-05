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

func (d *duck) setQuackBehaviour(qb quackBehaviour) {
	d.quackBehaviour = qb
}

func (d *duck) setFlyBehaviour(fb flyBehaviour) {
	d.flyBehaviour = fb
}

type rocketDuck struct {
	*duck
}

type toyDuck struct {
}
