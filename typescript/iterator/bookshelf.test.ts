import { assertEquals } from "https://deno.land/std@0.81.0/testing/asserts.ts";
import { Book } from './book.ts';
import { BookShelf } from './bookshelf.ts';

Deno.test({
  name: 'Bookshelf should append book and return it.',
  fn(): void {
    const book = new Book('test book');
    const shelf = new BookShelf();
    shelf.appendBook(book);
    assertEquals(shelf.getBookAt(0), book);
  }
});

Deno.test({
  name: 'Bookshelf should append 2 books and return 2nd one by index.',
  fn(): void {
    const book1 = new Book('test book1');
    const book2 = new Book('test book2');
    const shelf = new BookShelf();

    shelf.appendBook(book1);
    shelf.appendBook(book2);
    assertEquals(shelf.getBookAt(1), book2);
  }
});

Deno.test({
  name: 'Bookshelf.getLength() should return the number of books it has.',
  fn(): void {
    const book1 = new Book('test book1');
    const book2 = new Book('test book2');
    const shelf = new BookShelf();

    shelf.appendBook(book1);
    shelf.appendBook(book2);
    assertEquals(shelf.getLength(), 2);
  }
});

Deno.test({
  name: 'Bookshelf.createIterator() should return the iterator.',
  fn(): void {
    const book1 = new Book('test book1');
    const book2 = new Book('test book2');
    const shelf = new BookShelf();

    shelf.appendBook(book1);
    shelf.appendBook(book2);

    const expected = [book1, book2];
    const iterator = shelf.createIterator();

    let i = 0;
    while (iterator.hasNext()) {
      assertEquals(iterator.next(), expected[i]);
      i++;
    }

    assertEquals(iterator.next(), null);
  }
});