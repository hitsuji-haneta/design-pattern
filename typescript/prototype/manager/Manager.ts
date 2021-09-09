import { Product } from "./Product.ts";

export class Manager {
  private showcase: { [key: string]: Product };

  constructor() {
    this.showcase = {};
  }

  public register = (name: string, proto: Product) => {
    this.showcase[name] = proto;
  };

  public create = (name: string): Product => {
    const prd: Product = Object.assign(
      Object.create(
        Object.getPrototypeOf(this.showcase[name]),
      ),
      this.showcase[name],
    );

    return prd;
  };
}
