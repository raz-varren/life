package life

const(
	//CharLive = string('\u2b1c')
	//CharDead = string('\u2b1b')
	CharLive = "*"
	CharDead = " "
)

type Cell struct {
	Pos Pos
	Alive bool
	NextState State
	Neighbors []*Cell
}

func (c *Cell) init(e *Env) {
	p := c.Pos
	c.Neighbors[0] = e.Cells[p.Neighbor(0, 1, e.Bounds)]
	c.Neighbors[1] = e.Cells[p.Neighbor(1, 1, e.Bounds)]
	c.Neighbors[2] = e.Cells[p.Neighbor(1, 0, e.Bounds)]
	c.Neighbors[3] = e.Cells[p.Neighbor(1, -1, e.Bounds)]
	c.Neighbors[4] = e.Cells[p.Neighbor(0, -1, e.Bounds)]
	c.Neighbors[5] = e.Cells[p.Neighbor(-1, -1, e.Bounds)]
	c.Neighbors[6] = e.Cells[p.Neighbor(-1, 0, e.Bounds)]
	c.Neighbors[7] = e.Cells[p.Neighbor(-1, 1, e.Bounds)]
}

func (c *Cell) CalcNextState() {
	liveNeighbors := 0
	for _, n := range c.Neighbors{
		if n.Alive{
			liveNeighbors++
		}
	}
	
	if !c.Alive{
		if liveNeighbors == 3{
			c.NextState = StateAlive
		}
	}else{
		if liveNeighbors < 2 || liveNeighbors > 3{
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
	}else{
		return CharDead
	}
}