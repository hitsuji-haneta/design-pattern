package bookshelf

import "github.com/hitsuji-haneta/design-pattern/go/iterator/book"

type Book interface {
	GetName() string
}

type shelfIterator struct {
	shelf *bookShelf
	index int
}

func newIterator(shelf *bookShelf) *shelfIterator {
	return &shelfIterator{
		shelf: shelf,
		index: 0,
	}
}

func (si *shelfIterator) HasNext() bool {
	if si.index < si.shelf.GetLength() {
		return true
	}
	return false
}

func (si *shelfIterator) Next() *book.Book {
	book := si.shelf.GetBookAt(si.index)
	si.index++
	return book
}
