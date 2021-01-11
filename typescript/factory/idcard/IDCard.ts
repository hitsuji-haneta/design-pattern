import { Product } from '../framework/Product.ts';

export class IDCard extends Product {
  private owner: string;

  constructor(owner: string) {
    super();
    console.log(`Created ${owner}'s ID.`);
    this.owner = owner;
  }

  public use = (): string => {
    return `This is ${this.owner}'s ID.`;
  };

  public getOwner = (): string => this.owner;
}
