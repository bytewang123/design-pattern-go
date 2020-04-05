package main

import "fmt"

type flyBehaviour interface {
	fly()
}

type flyWithWings struct {
}

func (f *flyWithWings) fly() {
	fmt.Println("fly with wings!")
}

type cantFly struct {
}

func (f *cantFly) fly() {
	fmt.Println("cant fly...")
}

type flyWithRocketPower struct {
}

func (f *flyWithRocketPower) fly() {
	fmt.Println("fly with rocket power!")
}

type flyWithSonicRocketPower struct {
}

func (f *flyWithSonicRocketPower) fly() {
	fmt.Println("fly with sonic rocket power!")
}
