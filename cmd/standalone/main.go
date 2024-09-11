package main

import (
	"github.com/dmccarthy-dev/roadrunner-demo-fiber/internal/app/fiberapp"
)

func main() {

	//boot the fiber app
	fiberApp := fiberapp.Boot()

	err := fiberApp.Listen("localhost:8888")
	if err != nil {
		println("App crashed: ", err.Error())
	}

}
