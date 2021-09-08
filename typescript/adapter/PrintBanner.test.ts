import { assertEquals } from "https://deno.land/std@0.81.0/testing/asserts.ts";
import { Banner } from "./Banner.ts";
import { PrintBanner } from "./PrintBanner.ts";

Deno.test({
  name: "PrintBanner.printWeak() should return string.",
  fn(): void {
    const printBanner = new PrintBanner("Hello");
    const expected = "(Hello)";
    assertEquals(printBanner.printWeak(), expected);
  },
});

Deno.test({
  name: "PrintBanner.printStrong() should return string.",
  fn(): void {
    const printBanner = new PrintBanner("Hello");
    const expected = "*Hello*";
    assertEquals(printBanner.printStrong(), expected);
  },
});

