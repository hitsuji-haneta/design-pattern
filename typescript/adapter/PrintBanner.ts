import { Print } from "./interface/Print.ts";
import { Banner } from "./Banner.ts";

export class PrintBanner extends Banner implements Print {
  constructor(text: string) {
    super(text);
  }

  printWeak = (): string => {
    return this.showWithParen();
  };

  printStrong = (): string => {
    return this.showWithAster();
  };
}

