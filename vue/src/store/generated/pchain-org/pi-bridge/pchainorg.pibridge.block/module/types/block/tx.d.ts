import { Reader, Writer } from 'protobufjs/minimal';
export declare const protobufPackage = "pchainorg.pibridge.block";
/** this line is used by starport scaffolding # proto/tx/message */
export interface MsgCreateBlock {
    creator: string;
    index: string;
    chainID: number;
    address: string;
    headers: string[];
}
export interface MsgCreateBlockResponse {
}
export interface MsgUpdateBlock {
    creator: string;
    index: string;
    chainID: number;
    address: string;
    headers: string[];
}
export interface MsgUpdateBlockResponse {
}
export interface MsgDeleteBlock {
    creator: string;
    index: string;
}
export interface MsgDeleteBlockResponse {
}
export declare const MsgCreateBlock: {
    encode(message: MsgCreateBlock, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateBlock;
    fromJSON(object: any): MsgCreateBlock;
    toJSON(message: MsgCreateBlock): unknown;
    fromPartial(object: DeepPartial<MsgCreateBlock>): MsgCreateBlock;
};
export declare const MsgCreateBlockResponse: {
    encode(_: MsgCreateBlockResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateBlockResponse;
    fromJSON(_: any): MsgCreateBlockResponse;
    toJSON(_: MsgCreateBlockResponse): unknown;
    fromPartial(_: DeepPartial<MsgCreateBlockResponse>): MsgCreateBlockResponse;
};
export declare const MsgUpdateBlock: {
    encode(message: MsgUpdateBlock, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateBlock;
    fromJSON(object: any): MsgUpdateBlock;
    toJSON(message: MsgUpdateBlock): unknown;
    fromPartial(object: DeepPartial<MsgUpdateBlock>): MsgUpdateBlock;
};
export declare const MsgUpdateBlockResponse: {
    encode(_: MsgUpdateBlockResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateBlockResponse;
    fromJSON(_: any): MsgUpdateBlockResponse;
    toJSON(_: MsgUpdateBlockResponse): unknown;
    fromPartial(_: DeepPartial<MsgUpdateBlockResponse>): MsgUpdateBlockResponse;
};
export declare const MsgDeleteBlock: {
    encode(message: MsgDeleteBlock, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteBlock;
    fromJSON(object: any): MsgDeleteBlock;
    toJSON(message: MsgDeleteBlock): unknown;
    fromPartial(object: DeepPartial<MsgDeleteBlock>): MsgDeleteBlock;
};
export declare const MsgDeleteBlockResponse: {
    encode(_: MsgDeleteBlockResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteBlockResponse;
    fromJSON(_: any): MsgDeleteBlockResponse;
    toJSON(_: MsgDeleteBlockResponse): unknown;
    fromPartial(_: DeepPartial<MsgDeleteBlockResponse>): MsgDeleteBlockResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    /** this line is used by starport scaffolding # proto/tx/rpc */
    CreateBlock(request: MsgCreateBlock): Promise<MsgCreateBlockResponse>;
    UpdateBlock(request: MsgUpdateBlock): Promise<MsgUpdateBlockResponse>;
    DeleteBlock(request: MsgDeleteBlock): Promise<MsgDeleteBlockResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    CreateBlock(request: MsgCreateBlock): Promise<MsgCreateBlockResponse>;
    UpdateBlock(request: MsgUpdateBlock): Promise<MsgUpdateBlockResponse>;
    DeleteBlock(request: MsgDeleteBlock): Promise<MsgDeleteBlockResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
