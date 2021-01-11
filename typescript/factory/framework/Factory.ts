import { Product } from './Product.ts';

export abstract class Factory {
  protected abstract createdProduct(owner: string): Product;
  protected abstract registerProduct(product: Product): void;

  public create = (owner: string): Product => {
    const p: Product = this.createdProduct(owner);
    this.registerProduct(p);
    return p;
  };
}
