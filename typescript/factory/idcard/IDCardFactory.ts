import { IDCard } from './IDCard.ts';
import { Factory } from '../framework/Factory.ts';
import { Product } from '../framework/Product.ts';

export class IDCardFactory extends Factory {
  private owners: string[];

  constructor() {
    super();
    this.owners = [];
  }

  protected createdProduct = (owner: string): Product => {
    return new IDCard(owner);
  };

  protected registerProduct = (product: Product): void => {
    if (!this._isIDCard(product)) return;
    this.owners.push((product as IDCard).getOwner());
  };

  public getOwners = (): string[] => {
    return this.owners;
  };

  _isIDCard = (item: any): item is IDCard => {
    return item.getOwner() !== undefined && typeof item.getOwner() === 'string';
  };
}
