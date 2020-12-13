package bookshelf

import (
	"fmt"

	"github.com/hitsuji-haneta/design-pattern/go/iterator/book"
)

type Iterator interface {
	HasNext() bool
	Next() *book.Book
}

type bookShelf struct {
	books   []*book.Book
	last    int
	maxsize int
}

func NewShelf(maxsize int) *bookShelf {
	return &bookShelf{
		books:   []*book.Book{},
		last:    0,
		maxsize: maxsize,
	}
}

func (bs *bookShelf) GetBookAt(index int) *book.Book {
	return bs.books[index]
}

func (bs *bookShelf) AppendBook(book *book.Book) error {
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
	return newIterator(bs)
}
