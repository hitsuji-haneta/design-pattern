package bookshelf

import (
	"fmt"
)

type Book interface {
	GetName() string
}

type Iterator interface {
	HasNext() bool
	Next() Book
}

type bookShelf struct {
	books   []Book
	last    int
	maxsize int
}

func New(maxsize int) *bookShelf {
	return &bookShelf{
		books:   []Book{},
		last:    0,
		maxsize: maxsize,
	}
}

func (bs *bookShelf) GetBookAt(index int) Book {
	return bs.books[index]
}

func (bs *bookShelf) AppendBook(book Book) error {
	if len(bs.books) >= bs.maxsize {
		return fmt.Errorf("This book shelf had reached the maximum number of books")
	}

	bs.books = append(bs.books, book)
	return nil
}

func (bs *bookShelf) GetLength() int {
	return len(bs.books)
}

func (bs *bookShelf) CreateIterator() Iterator {
	return &shelfIterator{bs, 0}
}
