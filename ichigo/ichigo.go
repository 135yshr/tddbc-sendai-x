package ichigo

import "fmt"

type Ichigo struct {
	Name string
	Size string
}

func (c *Ichigo) String() string {
	return fmt.Sprintf("%s: %s", c.Name, c.Size)
}
