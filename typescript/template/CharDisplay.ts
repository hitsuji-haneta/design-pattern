import { AbstractDisplay } from './AbstractDisplay.ts';

export class CharDisplay extends AbstractDisplay {
  private char: string;

  constructor(char: string) {
    super();
    if (char.length === 1) {
      this.char = char;
    } else {
      throw new Error('Argument should be 1 character.');
    }
  }

  public open = (): string => {
    return '<';
  };

  public print = (_: boolean): string => {
    return this.char;
  };

  public close = (): string => {
    return '>';
  };
}
