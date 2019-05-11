package workers

import (
	"github.com/gonum/stat/distuv"
	"math"
	"math/rand"
	"sort"
	"time"
)

type Source struct {
	arrivalDay int // время поступления завяки
}

type Server struct {
	serveStartDay int
	duration      int
	serveEndDay   int
}

type Buffer struct {
	Day   int
	Event int
	Queue int
}

type QueuingSystem struct {
	RandomSeed *rand.Rand
	Source     []int
	Buffer     []Buffer
	Server     []Server
}

func (m *QueuingSystem) Init(λ float64, arrivalsNum int) {
	m.RandomSeed = rand.New(rand.NewSource(time.Now().UnixNano()))
	m.Source = make([]int, arrivalsNum)
	m.Server = make([]Server, arrivalsNum)

	arrivalGenerator := distuv.Exponential{
		Rate:   λ,
		Source: m.RandomSeed,
	}

	// Единица времени моделированиия - один день
	for i := 0; i < arrivalsNum; i++ {
		if i == 0 {
			m.Source[i] = int(math.Round(arrivalGenerator.Rand()))
			continue
		}
		m.Source[i] = m.Source[i-1] + int(math.Round(arrivalGenerator.Rand()))
	}
}

func (m *QueuingSystem) Modeling(μ float64) {
	servingGenerator := distuv.Exponential{
		Rate:   1 / μ,
		Source: m.RandomSeed,
	}

	// Единица времени моделированиия - один день
	for i := range m.Source {
		// начальное состояние стистемы
		var duration = int(math.Round(servingGenerator.Rand()))
		if duration == 0 {
			duration = 1
		}

		// Если это первая обработка
		if i == 0 {
			m.Server[i] = Server{
				serveStartDay: m.Source[i],
				duration:      duration,
				serveEndDay:   m.Source[i] + duration,
			}
		} else { // Если не первая
			// Если предыдущая обработка закончиилась раньше начала пришедшей новой заявки
			if m.Server[i-1].serveEndDay < m.Source[i] {
				m.Server[i] = Server{
					serveStartDay: m.Source[i],
					duration:      duration,
					serveEndDay:   m.Source[i] + duration,
				}
			} else { // Если предыдущая обработка ещё не закончиилась
				m.Server[i] = Server{
					serveStartDay: m.Server[i-1].serveEndDay,
					duration:      duration,
					serveEndDay:   m.Server[i-1].serveEndDay + duration,
				}
			}

			if m.Server[i-1].serveEndDay > m.Source[i] {
				m.Buffer = append(m.Buffer, Buffer{
					Day:   m.Source[i],
					Event: 1,
				})
			}
		}
		m.Buffer = append(m.Buffer, Buffer{
			Day:   m.Server[i].serveEndDay,
			Event: -1,
		})
	}
}

func (m *QueuingSystem) CalcBuffer() {
	// Сортировка буфера по дням
	sort.Slice(m.Buffer, func(i, j int) bool {
		return m.Buffer[i].Day < m.Buffer[j].Day
	})

	for i, b := range m.Buffer {
		if i == 0 {
			b.Queue = 0
			continue
		}

		if m.Buffer[i-1].Queue == 0 && m.Buffer[i].Event == -1 {
			m.Buffer[i].Queue = 0
			continue
		}

		m.Buffer[i].Queue = m.Buffer[i-1].Queue + m.Buffer[i].Event

	}
}
