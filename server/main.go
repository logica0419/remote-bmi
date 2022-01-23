package main

import "github.com/logica0419/remote-bmi/server/router"

func main() {
	e := router.SetupEcho()

	e.Logger.Panic(e.Start(":3000"))
}
