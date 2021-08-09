import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "pchainorg.pibridge.proxy";
export interface Proxy {
    creator: string;
    index: string;
    addressList: string;
    address: string;
}
export declare const Proxy: {
    encode(message: Proxy, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Proxy;
    fromJSON(object: any): Proxy;
    toJSON(message: Proxy): unknown;
    fromPartial(object: DeepPartial<Proxy>): Proxy;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
