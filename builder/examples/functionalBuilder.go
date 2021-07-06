package builder

type PersonFp struct {
	name, position string
}

type personMod func(*PersonFp)
type PersonFpBuilder struct {
	actions []personMod
}

func (b *PersonFpBuilder) Called(name string) *PersonFpBuilder {
	b.actions = append(b.actions, func(p *PersonFp) {
		p.name = name
	})
	return b
}

func (b *PersonFpBuilder) WorksAsA(position string) *PersonFpBuilder {
	b.actions = append(b.actions, func(p *PersonFp) {
		p.position = position
	})
	return b
}

func (b PersonFpBuilder) Build() *PersonFp {
	p := PersonFp{}
	for _, a := range b.actions {
		a(&p)
	}
	return &p
}
