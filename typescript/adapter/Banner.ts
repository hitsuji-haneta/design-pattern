export class Banner {
  protected text: string;

  constructor(text: string) {
    this.text = text;
  }

  showWithParen = (): string => {
    return `(${this.text})`;
  }

  showWithAster = (): string => {
    return `*${this.text}*`;
  }
}