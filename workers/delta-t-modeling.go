package workers

import (
	"github.com/gonum/stat/distuv"
	"math"
	"math/rand"
	"time"
)

type DeltaTimeModeling struct {
	randomSeed *rand.Rand
	Source     []int
	Buffer     []Buffer
	Server     []Server
}

func (d *DeltaTimeModeling) Init(lambda float64, arrivalsNum int) {
	d.randomSeed = rand.New(rand.NewSource(time.Now().UnixNano()))
	d.Source = make([]int, arrivalsNum)
	d.Server = make([]Server, arrivalsNum)

	arrivalGenerator := distuv.Exponential{
		Rate:   lambda,
		Source: d.randomSeed,
	}

	d.Source[0] = int(math.Round(arrivalGenerator.Rand()))
	for i := 1; i < arrivalsNum; i++ {
		d.Source[i] = d.Source[i-1] + int(math.Round(arrivalGenerator.Rand()))
	}
}

func (d *DeltaTimeModeling) Modeling(mu float64) {
	servingGenerator := distuv.Exponential{
		Rate:   1 / mu,
		Source: d.randomSeed,
	}

	for i := range d.Source {
		// начальное состояние стистемы
		var duration = int(math.Round(servingGenerator.Rand()))
		if duration == 0 {
			duration = 1
		}

		// Если это первая обработка
		if i == 0 {
			d.Server[i] = Server{
				serveStartDay: d.Source[i],
				duration:      duration,
				serveEndDay:   d.Source[i] + duration,
			}
		} else { // Если не первая
			// Если предыдущая обработка закончиилась раньше начала пришедшей новой заявки
			if d.Server[i-1].serveEndDay < d.Source[i] {
				d.Server[i] = Server{
					serveStartDay: d.Source[i],
					duration:      duration,
					serveEndDay:   d.Source[i] + duration,
				}
			} else { // Если предыдущая обработка ещё не закончиилась
				d.Server[i] = Server{
					serveStartDay: d.Server[i-1].serveEndDay,
					duration:      duration,
					serveEndDay:   d.Server[i-1].serveEndDay + duration,
				}
			}

			if d.Server[i-1].serveEndDay > d.Source[i] {
				d.Buffer = append(m.Buffer, Buffer{
					Day:   m.Source[i],
					Event: 1,
				})
			}
		}
		d.Buffer = append(d.Buffer, Buffer{
			Day:   d.Server[i].serveEndDay,
			Event: -1,
		})
	}
}
