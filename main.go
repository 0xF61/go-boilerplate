package main

import "boilerplate/app"

func main() {
	myapp, err := app.New()
	if err != nil {
		panic("Can't connect DB")
	}

	myapp.SetupROutes()
	myapp.Start()
}
