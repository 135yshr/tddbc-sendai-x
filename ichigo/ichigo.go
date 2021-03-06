package ichigo

import "fmt"

type Ichigo struct {
	Variety string
	Size    string
}

func (c *Ichigo) String() string {
	return fmt.Sprintf("%s: %s", c.Variety, c.Size)
}
