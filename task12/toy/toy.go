package toy

type Toy struct {
	Name   string
	Weight int
	
	onHand int
	sold   int
}

func New(name string, weight int) *Toy {
	return &Toy{
		Name:   name,
		Weight: weight,
	}
}

func (t *Toy) OnHand() int {
	return t.onHand
}

func (t *Toy) UpdateOnHand(count int) int {
	t.onHand += count
	return t.onHand
}

func (t *Toy) Sold() int {
	return t.sold
}

func (t *Toy) UpdateSold(count int) int {
	t.sold += count
	return t.sold
}
