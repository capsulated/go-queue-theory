package main

import (
	"fmt"
	"github.com/logiqone/go-queue-theory/workers"
)

/*
Задача 1
Моделируем региональный центр САИ
М/М/1
Региональный центр САИ прнимает поправки к АНИ от входящиих в его состав аэропортовых САИ
Обозначим его как объект Server
*/

func main() {
	// Возьмём за единиицу времении - один день - uint64
	// Приходит заявка на изменение аэронаигационных случайным образом:
	// В среднем за месяц приходит три заявки λ = 1 / 4 = 0.25 [ед/день]
	// Необходимо промоделировать этот процесс в течении T = 1000 дней
	// μ = 1 / λ = 1 / .25 = 4, тоесть в среднем одна заявка раз в четыре дня

	// Время обработки заявки в среднем 2 дня, соответсвенно μ = 2
	// λ = 1 / 4 = 0.25 [ед/день]

	// Количество входных заявок, которые собраемся моделировать
	// arrivalsNum = 20

	λ := .25
	μ := 2.
	arrivalsNum := 100

	qs := &workers.QueuingSystem{}
	qs.Init(λ, arrivalsNum)
	qs.Modeling(μ)

	qs.CalcBuffer()

	fmt.Printf("%v\n", qs.Source)
	fmt.Printf("%v\n", qs.Server)
	fmt.Printf("%v\n", qs.Buffer)

	drawer := &workers.Drawer{}
	drawer.Init(qs)
	drawer.DrawNumbers()
	drawer.DrawSource()
	drawer.DrawServer()

	drawer.DrawBuffer()

	if err := drawer.Save(); err != nil {
		fmt.Printf("Err: %s\n", err)
	}
}
