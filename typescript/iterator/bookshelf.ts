import { Aggregate } from './interfaces/aggregate.ts';
import { Iterator } from './interfaces/iterator.ts';
import { Book } from './book.ts';

export class BookShelf implements Aggregate {
  private books: Book[];

  constructor() {
    this.books = [];
  }

  appendBook = (book: Book): void => {
    this.books.push(book);
  };

  getBookAt = (index: number): Book => {
    return this.books[index];
  };

  getLength = (): number => {
    return this.books.length;
  };

  createIterator = (): Iterator => {
    return new ShelfIterator(this);
  }
}

class ShelfIterator implements Iterator {
  private bookShelf: BookShelf;
  private index: number;

  constructor(bookShelf: BookShelf) {
    this.bookShelf = bookShelf;
    this.index = 0;
  }

  hasNext = (): boolean => {
    if (this.index < this.bookShelf.getLength()) {
      return true;
    }
    return false;
  };

  next = (): Book | null => {
    if (!this.hasNext()) {
      return null;
    }
    const nextBook = this.bookShelf.getBookAt(this.index);
    this.index++;
    return nextBook;
  }
}