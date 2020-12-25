export abstract class AbstractDisplay {
  public abstract open(): string;
  public abstract print(isFinal: boolean): string;
  public abstract close(): string;

  display = (times: number): string => {
    const result: string[] = [];

    result.push(this.open());
    for (let i: number = 0; i < times; i++) {
      result.push(this.print(i === times - 1));
    }
    result.push(this.close());

    return result.join('');
  };
}
