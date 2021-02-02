import { assert, assertEquals } from "https://deno.land/std@0.81.0/testing/asserts.ts";
import { Singleton } from './singleton.ts';

Deno.test({
  name: 'Singleton create the same instance',
  fn(): void {
    const obj1 = Singleton.getInstance('hanako');
    const obj2 = Singleton.getInstance('tamako');
    assert(obj1 === obj2);
    assertEquals(obj1.name, obj2.name)
  }
});