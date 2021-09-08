import { assertEquals } from "https://deno.land/std@0.81.0/testing/asserts.ts";
import { MessageBox } from "./MessageBox.ts";

Deno.test({
  name: "MessageBox.use() should retuen strong box",
  fn(): void {
    const mbox = new MessageBox("=");
    const expected = `
=========
= Hello =
=========
`;
    assertEquals(mbox.use("Hello"), expected);
  },
});

Deno.test({
  name: "MessageBox.use() should retuen strong box with wrong message",
  fn(): void {
    const mbox = new MessageBox("*");
    const expected = `
****************************
* Hello, my name is Alice. *
****************************
`;
    assertEquals(mbox.use("Hello, my name is Alice."), expected);
  },
});
