package ichigo

import "fmt"

// Ichigo - いちごオブジェクト
type Ichigo struct {
	Variety string
	Size    string
}

func (c *Ichigo) String() string {
	return fmt.Sprintf("%s: %s", c.Variety, c.Size)
}

// NewWithVarietyAndWeight - 品種と重さ(g)からいちごを作成する
func NewWithVarietyAndWeight(variety string, weight uint) *Ichigo {
	switch {
	case weight < 10:
		return &Ichigo{Variety: variety, Size: "S"}
	case weight < 20:
		return &Ichigo{Variety: variety, Size: "M"}
	case weight < 30:
		return &Ichigo{Variety: variety, Size: "L"}
	}
	return &Ichigo{Variety: variety, Size: "S"}
}
