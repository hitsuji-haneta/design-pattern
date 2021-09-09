import { assertEquals } from "https://deno.land/std@0.81.0/testing/asserts.ts";
import { Manager } from "./manager/Manager.ts";
import { Product } from "./manager/Product.ts";
import { MessageBox } from "./messageBox/MessageBox.ts";
import { Underline } from "./underline/Underline.ts";

Deno.test({
  name: "【Main】Prototype pattern",
  fn(): void {
    const proto1 = new MessageBox("*");
    const proto2 = new MessageBox("=");
    const proto3 = new Underline("~");
    const proto4 = new Underline("=");

    const manager = new Manager();
    manager.register("star box", proto1);
    manager.register("double box", proto2);
    manager.register("wave line", proto3);
    manager.register("double line", proto4);

    const expected1 = `
*********
* Hello *
*********
`;
    const expected2 = `
=========
= Hello =
=========
`;
    const expected3 = `
Hello
~~~~~
`;
    const expected4 = `
Hello
=====
`;

    assertEquals(manager.create("star box").use("Hello"), expected1);
    assertEquals(manager.create("double box").use("Hello"), expected2);
    assertEquals(manager.create("wave line").use("Hello"), expected3);
    assertEquals(manager.create("double line").use("Hello"), expected4);
  },
});
