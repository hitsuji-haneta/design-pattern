import { assertEquals } from "https://deno.land/std@0.81.0/testing/asserts.ts";
import { Manager } from "./Manager.ts";
import { Product } from "./Product.ts";

Deno.test({
  name:
    "Manager.create() should retuen the same class with registered products.",
  fn(): void {
    const sut = new Manager();
    const mock: Product = {
      use: (msg: string) => msg,
    };
    sut.register("dummy", mock);
    const actual = sut.create("dummy")
    assertEquals(actual.use("Hello"), "Hello");
  },
});
