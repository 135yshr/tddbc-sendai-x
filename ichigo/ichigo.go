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
	return &Ichigo{Variety: variety, Size: "S"}
}
