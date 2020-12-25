import { AbstractDisplay } from './AbstractDisplay.ts';

export class StringDisplay extends AbstractDisplay {
  private text: string;

  constructor(text: string) {
    super();
    this.text = text;
  }

  public open = (): string => {
    return '"';
  };

  public print = (isFinal: boolean): string => {
    if (isFinal) return this.text;
    return `${this.text}, `;
  };

  public close = (): string => {
    return '"';
  };
}
