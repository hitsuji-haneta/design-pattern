package iterator

import (
	"fmt"

	"github.com/hitsuji-haneta/design-pattern/go/iterator/book"
	"github.com/hitsuji-haneta/design-pattern/go/iterator/bookshelf"
)

func Run() {
	fmt.Println("hello, world")
	b := book.New("sample")
	fmt.Println(b.GetName())

	bs := bookshelf.New(3)
	fmt.Println(bs.GetLength())

	bs.AppendBook(b)
	newBook := bs.GetBookAt(0)
	fmt.Println(newBook.GetName())
}
