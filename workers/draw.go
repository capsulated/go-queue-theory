package workers

import (
	"fmt"
	"github.com/fogleman/gg"
)

type Drawer struct {
	Source  []int
	Server  *[]Server
	Buffer  []Buffer
	Context *gg.Context
}

func (d *Drawer) Init(modeler *QueuingSystem) {
	d.Source = modeler.Source
	d.Server = &modeler.Server
	d.Buffer = modeler.Buffer

	// init graph
	d.Context = gg.NewContext(5000, 500)

	// Background
	d.Context.DrawRectangle(0, 0, 1500, 500)
	d.Context.SetRGB255(255, 255, 255)
	d.Context.Fill()

	// Cage
	red := 235
	green := 235
	blue := 235
	for i := 0; i < 1500; i = i + 10 {
		d.Context.DrawRectangle(0, 10+float64(i), 1500, 1)
		d.Context.SetRGB255(red, green, blue)
		d.Context.Fill()

		d.Context.DrawRectangle(10+float64(i), 0, 1, 500)
		d.Context.SetRGB255(red, green, blue)
		d.Context.Fill()
	}
}

func (d *Drawer) DrawSource() {
	// draw
	red := 200
	green := 200
	blue := 200
	radius := 5.0

	ident := 5.0
	step := 10.0

	initY := 20.0
	for i, s := range d.Source {
		// it's HORRBLE ! REFACTOR !!
		if i != 0 && s == d.Source[i-1] {
			d.Context.DrawPoint(ident+step*float64(s), initY+step, radius)
			d.Context.SetRGB255(0, 0, 0)
			d.Context.Fill()
			d.Context.DrawPoint(ident+10*float64(s), initY+step, radius-1)
			d.Context.SetRGB255(red, green, blue)
			d.Context.Fill()
			continue
		}

		if i != 0 && i != 1 && s == d.Source[i-2] {
			d.Context.DrawPoint(ident+step*float64(s), initY+step+2*step, radius)
			d.Context.SetRGB255(0, 0, 0)
			d.Context.Fill()
			d.Context.DrawPoint(ident+step*float64(s), initY+step+2*step, radius-1)
			d.Context.SetRGB255(red, green, blue)
			d.Context.Fill()
			continue
		}

		d.Context.DrawPoint(ident+10*float64(s), initY, radius)
		d.Context.SetRGB255(0, 0, 0)
		d.Context.Fill()

		d.Context.DrawPoint(ident+10*float64(s), initY, radius-1)
		d.Context.SetRGB255(red, green, blue)
		d.Context.Fill()
	}
}

func (d *Drawer) DrawServer() {
	red := 200
	green := 200
	blue := 200

	// draw
	for i, s := range *d.Server {
		// it's HORRBLE ! REFACTOR !!
		resX := 10 * float64(s.serveStartDay)
		width := 10 * float64(s.duration)

		d.Context.DrawRoundedRectangle(resX, 50+float64(i)*10, width, 10, 5)
		d.Context.SetRGB255(0, 0, 0)
		d.Context.Fill()

		d.Context.DrawRoundedRectangle(resX, 51+float64(i)*10, width, 8, 5)
		d.Context.SetRGB255(red, green, blue)
		d.Context.Fill()
	}
}

func (d *Drawer) DrawNumbers() {
	for i := .0; i < 100; i++ {
		if int(i)%2 != 0 {
			d.Context.SetRGB255(0, 0, 0)
			d.Context.DrawString(fmt.Sprintf("%d", int(i)), 10*i, 10)
			d.Context.Fill()
		}
	}
}

func (d *Drawer) DrawBuffer() {
	red := 0
	green := 0
	blue := 0

	startY := 500.0
	step := 10.0
	for i := range d.Buffer {
		if i > 0 {
			if d.Buffer[i].Day == d.Buffer[i-1].Day {
				d.Context.SetRGB255(255, 255, 255)
				d.Context.DrawRectangle(step*float64(d.Buffer[i].Day)-1, startY-9, 9, 9)
				d.Context.Fill()
			}
		}
		d.Context.SetRGB255(red, green, blue)
		d.Context.DrawString(fmt.Sprintf("%d", d.Buffer[i].Queue), step*float64(d.Buffer[i].Day), startY)
		d.Context.Fill()
	}
}

func (d *Drawer) Save() error {
	return d.Context.SavePNG("out.png")
}
