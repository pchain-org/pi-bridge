import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "pchainorg.pibridge.trx";
export interface Trx {
    creator: string;
    index: string;
    chainID: number;
    address: string;
    crossChainMsgs: string;
}
export declare const Trx: {
    encode(message: Trx, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Trx;
    fromJSON(object: any): Trx;
    toJSON(message: Trx): unknown;
    fromPartial(object: DeepPartial<Trx>): Trx;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
