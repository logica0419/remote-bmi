package main

import (
	"time"

	"github.com/logica0419/remote-bmi/server/cmd"
)

func init() {
	const location = "Asia/Tokyo"

	loc, err := time.LoadLocation(location)
	if err != nil {
		loc = time.FixedZone(location, 9*60*60)
	}

	time.Local = loc
}

func main() {
	cmd.Execute()
}
