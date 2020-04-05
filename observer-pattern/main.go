package main

func main() {
	wd := &weatherData{}
	wd.init()

	ccd := &currentConditionDisplay{
		sub: wd,
	}
	fd := &forecastDisplay{
		sub: wd,
	}
	ccd.sub.registerObserver("currentConditionDisplay", ccd)
	ccd.sub.registerObserver("forecastDisplay", fd)

	wd.setMeasurements(10.1, 20.2, 30.3)
}
