export class Book {
  constructor(
    private name: string,
  ) {}

  getName = (): string => this.name;
}