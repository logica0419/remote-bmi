package main

func main() {
	e := setupEcho()

	e.Logger.Panic(e.Start(":3000"))
}
