export class Book {
  private name: string;

  constructor(_name: string) {
    this.name = _name;
  }

  getName = (): string => this.name;
}

