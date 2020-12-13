package bookshelf

// import (
// 	"github.com/hitsuji-haneta/design-pattern/go/iterator/book"
// )

type Book interface {
	GetName() string
}

type bookshelf struct {
	books   []Book
	last    int
	maxsize int
}

func New(maxsize int) *bookshelf {
	return &bookshelf{
		books:   []Book{},
		last:    0,
		maxsize: maxsize,
	}
}

func (bs *bookshelf) GetBookAt(index int) Book {
	return bs.books[index]
}

func (bs *bookshelf) AppendBook(book Book) {
	bs.books = append(bs.books, book)
}

func (bs *bookshelf) GetLength() int {
	return len(bs.books)
}
