package ichigo

type Ichigo struct {
	Name string
	Size string
}

func (c *Ichigo) String() string {
	return "さちおとめ: L"
}
