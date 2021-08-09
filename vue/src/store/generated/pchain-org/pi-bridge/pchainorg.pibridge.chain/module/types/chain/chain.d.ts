import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "pchainorg.pibridge.chain";
export interface Chain {
    creator: string;
    index: string;
    sourceChainID: number;
    height: number;
    proof: string;
    proxyAddress: string;
    extra: string;
    headerOrCrossChainMsg: string;
}
export declare const Chain: {
    encode(message: Chain, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Chain;
    fromJSON(object: any): Chain;
    toJSON(message: Chain): unknown;
    fromPartial(object: DeepPartial<Chain>): Chain;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
