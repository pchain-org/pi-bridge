import { Reader, Writer } from 'protobufjs/minimal';
export declare const protobufPackage = "pchainorg.pibridge.chain";
/** this line is used by starport scaffolding # proto/tx/message */
export interface MsgCreateChain {
    creator: string;
    index: string;
    sourceChainID: number;
    height: number;
    proof: string;
    proxyAddress: string;
    extra: string;
    headerOrCrossChainMsg: string;
}
export interface MsgCreateChainResponse {
}
export interface MsgUpdateChain {
    creator: string;
    index: string;
    sourceChainID: number;
    height: number;
    proof: string;
    proxyAddress: string;
    extra: string;
    headerOrCrossChainMsg: string;
}
export interface MsgUpdateChainResponse {
}
export interface MsgDeleteChain {
    creator: string;
    index: string;
}
export interface MsgDeleteChainResponse {
}
export declare const MsgCreateChain: {
    encode(message: MsgCreateChain, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateChain;
    fromJSON(object: any): MsgCreateChain;
    toJSON(message: MsgCreateChain): unknown;
    fromPartial(object: DeepPartial<MsgCreateChain>): MsgCreateChain;
};
export declare const MsgCreateChainResponse: {
    encode(_: MsgCreateChainResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateChainResponse;
    fromJSON(_: any): MsgCreateChainResponse;
    toJSON(_: MsgCreateChainResponse): unknown;
    fromPartial(_: DeepPartial<MsgCreateChainResponse>): MsgCreateChainResponse;
};
export declare const MsgUpdateChain: {
    encode(message: MsgUpdateChain, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateChain;
    fromJSON(object: any): MsgUpdateChain;
    toJSON(message: MsgUpdateChain): unknown;
    fromPartial(object: DeepPartial<MsgUpdateChain>): MsgUpdateChain;
};
export declare const MsgUpdateChainResponse: {
    encode(_: MsgUpdateChainResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateChainResponse;
    fromJSON(_: any): MsgUpdateChainResponse;
    toJSON(_: MsgUpdateChainResponse): unknown;
    fromPartial(_: DeepPartial<MsgUpdateChainResponse>): MsgUpdateChainResponse;
};
export declare const MsgDeleteChain: {
    encode(message: MsgDeleteChain, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteChain;
    fromJSON(object: any): MsgDeleteChain;
    toJSON(message: MsgDeleteChain): unknown;
    fromPartial(object: DeepPartial<MsgDeleteChain>): MsgDeleteChain;
};
export declare const MsgDeleteChainResponse: {
    encode(_: MsgDeleteChainResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteChainResponse;
    fromJSON(_: any): MsgDeleteChainResponse;
    toJSON(_: MsgDeleteChainResponse): unknown;
    fromPartial(_: DeepPartial<MsgDeleteChainResponse>): MsgDeleteChainResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    /** this line is used by starport scaffolding # proto/tx/rpc */
    CreateChain(request: MsgCreateChain): Promise<MsgCreateChainResponse>;
    UpdateChain(request: MsgUpdateChain): Promise<MsgUpdateChainResponse>;
    DeleteChain(request: MsgDeleteChain): Promise<MsgDeleteChainResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    CreateChain(request: MsgCreateChain): Promise<MsgCreateChainResponse>;
    UpdateChain(request: MsgUpdateChain): Promise<MsgUpdateChainResponse>;
    DeleteChain(request: MsgDeleteChain): Promise<MsgDeleteChainResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
