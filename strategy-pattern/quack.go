package main

import "fmt"

type quackBehaviour interface {
	quack()
}

type quackWithGuaGua struct {
}

func (q *quackWithGuaGua) quack() {
	fmt.Println("quack with guagua!")
}

type quackWithGuGu struct {
}

func (q *quackWithGuGu) quack() {
	fmt.Println("quack with gugu!")
}

type quackWithEngine struct {
}

func (q *quackWithEngine) quack() {
	fmt.Println("quack with engine!")
}
