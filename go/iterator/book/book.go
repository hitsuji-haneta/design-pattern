package book

type book struct {
	name string
}

func New(name string) *book {
	return &book{name}
}

func (b *book) GetName() string {
	return b.name
}
