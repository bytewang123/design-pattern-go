package main

func main() {
	d := rocketDuck{
		&duck{
			&quackWithEngine{},
			&flyWithRocketPower{},
			"rocket duck",
		},
	}
	d.setFlyBehaviour(&flyWithSonicRocketPower{})
	doAction(d)
}

func doAction(d basicDuck) {
	d.display()
	d.swim()
	d.quack()
	d.fly()
}
