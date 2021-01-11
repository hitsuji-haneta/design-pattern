import {
  assertEquals,
  assertThrows,
} from 'https://deno.land/std@0.81.0/testing/asserts.ts';
import { IDCard } from './idcard/IDCard.ts';
import { IDCardFactory } from './idcard/IDCardFactory.ts';

Deno.test({
  name: 'IDCard.use() should return a text message.',
  fn(): void {
    const factory = new IDCardFactory();
    const card1 = factory.create('Alice');
    const card2 = factory.create('Bob');

    assertEquals(card1.use(), "This is Alice's ID.");
  },
});

Deno.test({
  name: 'Factory.getOwners() should return owner names.',
  fn(): void {
    const factory = new IDCardFactory();
    const card1 = factory.create('Alice');
    const card2 = factory.create('Bob');

    assertEquals(factory.getOwners(), ['Alice', 'Bob']);
  },
});
