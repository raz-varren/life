package main

import (
	"fmt"
	"github.com/raz-varren/life"
	"time"
)

func main() {
	b := life.Bounds{30, 20}
	e := life.NewEnv(b)

	p := []life.Pos{
		{3, 3}, {3, 4}, {4, 2}, {4, 3}, {4, 4},
		{5, 2}, {5, 3}, {5, 5}, {6, 3}, {6, 4},
		{6, 5}, {7, 4}, {25, 4}, {25, 5}, {25, 6},
	}

	e.SetLivingCells(p...)

	e.PrintLife()
	time.Sleep(time.Second * 3)

	for {
		e.Next()
		e.PrintLife()
		time.Sleep(time.Millisecond * 100)
		if e.Dead {
			fmt.Println("Your environment has stagnated")
			break
		}
	}
}
