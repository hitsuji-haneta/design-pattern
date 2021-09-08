export class Underline {
  private decorator: string;

  constructor(_decorator: string) {
    this.decorator = _decorator;
  }

  public use = (msg: string): string => {
    let line = "";
    for (let i = 0; i < msg.length; i++) line += this.decorator;

    return `\n${msg}\n${line}\n`;
  };
}
