import { assertEquals } from "https://deno.land/std@0.81.0/testing/asserts.ts";
import { Underline } from "./Underline.ts";

Deno.test({
  name: "Underline.use() should retuen strong box",
  fn(): void {
    const mbox = new Underline("=");
    const expected = `
Hello
=====
`;
    assertEquals(mbox.use("Hello"), expected);
  },
});

Deno.test({
  name: "Underline.use() should retuen strong box with wrong message",
  fn(): void {
    const mbox = new Underline("~");
    const expected = `
Hello, my name is Alice.
~~~~~~~~~~~~~~~~~~~~~~~~
`;
    assertEquals(mbox.use("Hello, my name is Alice."), expected);
  },
});
