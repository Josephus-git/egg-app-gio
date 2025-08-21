package main

import (
	"log"

	"gioui.org/app"
	"gioui.org/unit"
)

func main() {
	go func() {
		// create a new window
		w := new(app.Window)
		w.Option(app.Title("Egg timer"))
		w.Option(app.Size(unit.Dp(400), unit.Dp(600)))

		if err := draw(w); err != nil {
			log.Fatal(err)
		}

	}()
	app.Main()
}
