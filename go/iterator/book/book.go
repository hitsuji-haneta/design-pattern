package book

type Book struct {
	name string
}

func New(name string) *Book {
	return &Book{name}
}

func (b *Book) GetName() string {
	return b.name
}
