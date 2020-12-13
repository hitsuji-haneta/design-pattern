package iterator

import (
	"fmt"

	"github.com/hitsuji-haneta/design-pattern/go/iterator/book"
	"github.com/hitsuji-haneta/design-pattern/go/iterator/bookshelf"
)

func Run() (err error) {
	bs := bookshelf.New(4)
	if err = bs.AppendBook(book.New("Around the World in 80 days")); err != nil {
		return err
	}
	if err = bs.AppendBook(book.New("Bible")); err != nil {
		return err
	}
	if err = bs.AppendBook(book.New("Cinderella")); err != nil {
		return err
	}
	if err = bs.AppendBook(book.New("Daddy-Long-Legs")); err != nil {
		return err
	}

	it := bs.CreateIterator()
	for it.HasNext() {
		book := it.Next()
		fmt.Println(book.GetName())
	}
	return nil
}
