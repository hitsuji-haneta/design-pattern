package bookshelf

type shelfIterator struct {
	shelf *bookShelf
	index int
}

func (si *shelfIterator) HasNext() bool {
	if si.index < si.shelf.GetLength() {
		return true
	}
	return false
}

func (si *shelfIterator) Next() Book {
	if !si.HasNext() {
		return nil
	}
	book := si.shelf.GetBookAt(si.index)
	si.index++
	return book
}
