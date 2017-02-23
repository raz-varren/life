package life

import (
	"bytes"
	"fmt"
)

const (
	StateSame  State = 0
	StateAlive       = 1
	StateDead        = -1
)

type State int

type Bounds struct {
	W, H int
}

type Pos struct {
	X, Y int
}

type Env struct {
	Dead       bool
	Generation int64
	Bounds     Bounds
	Cells      map[Pos]*Cell
	HashCache  []string
}

func NewEnv(bounds Bounds) *Env {
	env := &Env{
		Dead:       false,
		Generation: 0,
		Bounds:     bounds,
		Cells:      make(map[Pos]*Cell),

		//initialize these so they don't match
		HashCache: []string{
			"hello",
			"goodbye",
			"asshole",
		},
	}

	for x := 1; x <= bounds.W; x++ {
		for y := 1; y <= bounds.H; y++ {
			p := Pos{x, y}
			c := &Cell{
				Pos:       p,
				Alive:     false,
				NextState: StateSame,
				Neighbors: make([]*Cell, 8),
			}
			env.Cells[p] = c
		}
	}

	for _, cell := range env.Cells {
		cell.init(env)
	}

	return env
}

func (e *Env) SetLivingCells(p ...Pos) {
	for _, pos := range p {
		if pos.X < 1 || pos.X > e.Bounds.W || pos.Y < 1 || pos.Y > e.Bounds.H {
			continue
		}
		e.Cells[pos].Alive = true
	}
}

func (e *Env) PrintLife() {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("Generation: %d\n", e.Generation))
	for y := 1; y <= e.Bounds.H; y++ {
		for x := 1; x <= e.Bounds.W; x++ {
			if x == 1 {
				buf.WriteByte(' ')
			}

			p := Pos{x, y}
			buf.WriteString(e.Cells[p].String())
			buf.WriteByte(' ')
		}
		buf.WriteByte('\n')
	}
	buf.WriteByte('\n')
	fmt.Print(buf.String())
}

func (e *Env) Next() {
	for _, c := range e.Cells {
		c.CalcNextState()
	}
	var buf bytes.Buffer
	for x := 1; x <= e.Bounds.W; x++ {
		for y := 1; y <= e.Bounds.H; y++ {
			c := e.Cells[Pos{x, y}]
			c.SetNextState()
			if c.Alive {
				buf.WriteByte('1')
			} else {
				buf.WriteByte('0')
			}
		}
	}

	setSum(e.HashCache, buf.String())
	e.Dead = checkSums(e.HashCache)

	e.Generation++
}

func (p Pos) Neighbor(xDir, yDir int, b Bounds) Pos {
	x, y := p.X, p.Y

	if xDir > 1 {
		xDir = 1
	}
	if xDir < -1 {
		xDir = -1
	}
	if yDir > 1 {
		yDir = 1
	}
	if yDir < -1 {
		yDir = -1
	}

	x += xDir
	y += yDir

	if x > b.W {
		x = 1
	}
	if x < 1 {
		x = b.W
	}
	if y > b.H {
		y = 1
	}
	if y < 1 {
		y = b.H
	}

	return Pos{x, y}
}

func setSum(sums []string, newSum string) {
	for i, iLen := 0, len(sums); i < iLen; i++ {
		if i == iLen-1 {
			sums[i] = newSum
			return
		}
		sums[i] = sums[i+1]
	}
}

func checkSums(sums []string) bool {
	for k1, v1 := range sums {
		for k2, sLen := k1+1, len(sums); k2 < sLen; k2++ {
			v2 := sums[k2]
			if v1 == v2 {
				return true
			}
		}
	}
	return false
}
