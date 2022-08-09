package examples

type Person struct {
	weight int
	height int
}

func (p *Person) GetWeight() int {
	return p.weight
}

func (p *Person) GetHeight() int {
	return p.height
}

func (p *Person) SetWeight(weight int) {
	p.weight = weight
}

func (p *Person) SetHeight(height int) {
	p.height = height
}
