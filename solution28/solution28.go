package solution28

type Counter struct {
	Value int
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func (c *Counter) Inc(delta int) {
	if delta <= 0 {
		c.Value++
	} else {
		c.Value += delta
	}
}

func (c *Counter) Dec(delta int) {
	if delta <= 0 {
		c.Value--
	} else {
		c.Value -= delta
	}
	if c.Value < 0 {
		c.Value = 0
	}
}
