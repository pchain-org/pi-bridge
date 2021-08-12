import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "pchainorg.pibridge.block";
export interface Block {
    creator: string;
    index: string;
    chainID: number;
    address: string;
    headers: string[];
}
export declare const Block: {
    encode(message: Block, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Block;
    fromJSON(object: any): Block;
    toJSON(message: Block): unknown;
    fromPartial(object: DeepPartial<Block>): Block;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
