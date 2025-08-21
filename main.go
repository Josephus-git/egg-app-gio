package main

import (
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func main() {
	go func() {
		// create a new window
		w := new(app.Window)
		w.Option(app.Title("Egg timer"))
		w.Option(app.Size(unit.Dp(400), unit.Dp(600)))

		// ops are the operations from the UI
		var ops op.Ops

		// startButton is a clickable widget
		var startButton widget.Clickable

		// th defines the material design style
		th := material.NewTheme()

		// listen for events in the window
		for {
			// first grab the event
			evt := w.Event()

			// detect the type
			switch typ := evt.(type) {

			// this is sent when the application should re-render.
			case app.FrameEvent:
				gtx := app.NewContext(&ops, typ)

				layout.Flex{
					Axis:    layout.Vertical,
					Spacing: layout.SpaceStart,
				}.Layout(gtx,

					layout.Rigid(
						func(gtx layout.Context) layout.Dimensions {
							btn := material.Button(th, &startButton, "Start")
							return btn.Layout(gtx)
						},
					),
					// ... one rigid to hold an empty spacer
					layout.Rigid(
						layout.Spacer{Height: unit.Dp(25)}.Layout,
					),
				)
				btn := material.Button(th, &startButton, "Start")
				btn.Layout(gtx)

				typ.Frame(gtx.Ops)

			// this is sent when the application should exit
			case app.DestroyEvent:
				os.Exit(0)
			}

		}
	}()
	app.Main()
}
