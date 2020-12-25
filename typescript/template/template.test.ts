import {
  assertEquals,
  assertThrows,
} from 'https://deno.land/std@0.81.0/testing/asserts.ts';
import { AbstractDisplay } from './AbstractDisplay.ts';
import { CharDisplay } from './CharDisplay.ts';
import { StringDisplay } from './StringDisplay.ts';

Deno.test({
  name: 'CharDisplay.display() should returns five characters with <>',
  fn(): void {
    const sut: AbstractDisplay = new CharDisplay('A');
    const expected = '<AAAAA>';
    assertEquals(sut.display(5), expected);
  },
});

Deno.test({
  name: 'CharDisplay should throw exception when the argument is not 1 char',
  fn(): void {
    assertThrows(
      () => {
        const sut: AbstractDisplay = new CharDisplay('AA');
      },
      Error,
      'Argument should be 1 character.'
    );
  },
});

Deno.test({
  name: 'StringDisplay.display() should returns five sentences with ""',
  fn(): void {
    const sut: AbstractDisplay = new StringDisplay('Hello');
    const expected = '"Hello, Hello, Hello, Hello, Hello"';
    assertEquals(sut.display(5), expected);
  },
});
