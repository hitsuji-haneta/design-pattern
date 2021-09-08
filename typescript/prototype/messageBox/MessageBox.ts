export class MessageBox {
  private decorator: string;

  constructor(_decorator: string) {
    this.decorator = _decorator;
  }

  public use = (msg: string): string => {
    let line = "";
    for (let i = 0; i < msg.length + 4; i++) line += this.decorator;

    return `\n${line}\n${this.decorator} ${msg} ${this.decorator}\n${line}\n`;
  };
}
