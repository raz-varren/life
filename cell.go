package life

const (
	//CharLive = string('\u2b1c')
	//CharDead = string('\u2b1b')
	CharLive = "*"
	CharDead = " "
)

type Cell struct {
	Pos       Pos
	Alive     bool
	NextState State
	Neighbors []*Cell
}

func (c *Cell) init(e *Env) {
	p := c.Pos
	i := 0
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if x == 0 && y == 0 {
				continue
			}
			c.Neighbors[i] = e.Cells[p.Neighbor(x, y, e.Bounds)]
			i++
		}
	}
}

func (c *Cell) CalcNextState() {
	liveNeighbors := 0
	for _, n := range c.Neighbors {
		if n.Alive {
			liveNeighbors++
		}
	}

	if !c.Alive {
		if liveNeighbors == 3 {
			c.NextState = StateAlive
		}
	} else {
		if liveNeighbors < 2 || liveNeighbors > 3 {
			c.NextState = StateDead
		}
	}
}

func (c *Cell) SetNextState() {
	switch c.NextState {
	case StateSame:
		return
	case StateAlive:
		c.Alive = true
	case StateDead:
		c.Alive = false
	}

	c.NextState = StateSame
}

func (c *Cell) String() string {
	if c.Alive {
		return CharLive
	} else {
		return CharDead
	}
}
