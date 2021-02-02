export class Singleton {
    private static instance: Singleton;
    private _name: string;

    private constructor(name: string) {
        this._name = name;
    }

    static getInstance = (name: string): Singleton => {
        if (!Singleton.instance) {
            Singleton.instance = new Singleton(name);
        }

        return Singleton.instance;
    }

    get name(): string {
        return this._name;
    };
}