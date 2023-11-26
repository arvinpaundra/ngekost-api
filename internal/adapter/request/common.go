package request

type (
	Common struct {
		Limit  int
		Offset int
		Search string
	}
)

func (c *Common) GetLimit() int {
	if c.Limit < 10 {
		return 25
	}
	return c.Limit
}

func (c *Common) GetOffset() int {
	if c.Offset-1 < 0 {
		return 0
	}
	return c.Offset - 1
}
