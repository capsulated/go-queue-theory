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
	for i := 0; i < 1500; i = i + 10 {
		d.Context.DrawRectangle(0, 10+float64(i), 1500, 1)
		d.Context.SetRGB255(200, 200, 200)
		d.Context.Fill()

		d.Context.DrawRectangle(10+float64(i), 0, 1, 500)
		d.Context.SetRGB255(200, 200, 200)
		d.Context.Fill()
	}
}

func (d *Drawer) DrawSource() {
	// Lines
	d.Context.DrawRectangle(100, 100, 1300, 1)
	d.Context.SetRGB255(0, 0, 0)
	d.Context.Fill()

	// draw
	for i, s := range d.Source {
		// it's HORRBLE ! REFACTOR !!
		if i != 0 && s == d.Source[i-1] {
			d.Context.DrawPoint(105+10*float64(s), 110, 5)
			d.Context.SetRGB255(0, 0, 0)
			d.Context.Fill()
			d.Context.DrawPoint(105+10*float64(s), 110, 4)
			d.Context.SetRGB255(255, 5, 5)
			d.Context.Fill()
			continue
		}

		if i != 0 && i != 1 && s == d.Source[i-2] {
			d.Context.DrawPoint(105+10*float64(s), 120, 5)
			d.Context.SetRGB255(0, 0, 0)
			d.Context.Fill()
			d.Context.DrawPoint(105+10*float64(s), 120, 4)
			d.Context.SetRGB255(255, 5, 5)
			d.Context.Fill()
			continue
		}

		d.Context.DrawPoint(105+10*float64(s), 100, 5)
		d.Context.SetRGB255(0, 0, 0)
		d.Context.Fill()

		d.Context.DrawPoint(105+10*float64(s), 100, 4)
		d.Context.SetRGB255(255, 5, 5)
		d.Context.Fill()
	}
}

func (d *Drawer) DrawServer() {
	// Lines
	d.Context.DrawRectangle(100, 200, 1300, 1)
	d.Context.SetRGB255(0, 0, 0)
	d.Context.Fill()

	// draw
	for _, s := range *d.Server {
		// it's HORRBLE ! REFACTOR !!
		resX := 100 + 10*float64(s.serveStartDay)
		width := 10 * float64(s.duration)

		d.Context.DrawRoundedRectangle(resX, 195, width, 10, 5)
		d.Context.SetRGB255(0, 0, 0)
		d.Context.Fill()

		d.Context.DrawRoundedRectangle(resX, 196, width, 8, 5)
		d.Context.SetRGB255(255, 5, 255)
		d.Context.Fill()
	}
}

func (d *Drawer) DrawNumbers() {
	for i := .0; i < 100; i++ {
		if int(i)%2 != 0 {
			d.Context.SetRGB255(0, 0, 0)
			d.Context.DrawString(fmt.Sprintf("%d", int(i)), 103+10*i, 70)
			d.Context.Fill()
		}
	}
}

func (d *Drawer) DrawBuffer() {

	for i := range d.Buffer {
		if i > 0 {
			if d.Buffer[i].Day == d.Buffer[i-1].Day {
				d.Context.SetRGB255(255, 255, 255)
				d.Context.DrawRectangle(102+10*float64(d.Buffer[i].Day)-1, 131, 9, 9)
				d.Context.Fill()
			}
		}
		d.Context.SetRGB255(0, 0, 255)
		d.Context.DrawString(fmt.Sprintf("%d", d.Buffer[i].Queue), 103+10*float64(d.Buffer[i].Day), 140)
		d.Context.Fill()
	}
}

func (d *Drawer) Save() error {
	return d.Context.SavePNG("out.png")
}
