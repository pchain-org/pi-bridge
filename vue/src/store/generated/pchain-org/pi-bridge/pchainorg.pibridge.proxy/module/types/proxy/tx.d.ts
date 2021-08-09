import { Reader, Writer } from 'protobufjs/minimal';
export declare const protobufPackage = "pchainorg.pibridge.proxy";
/** this line is used by starport scaffolding # proto/tx/message */
export interface MsgCreateProxy {
    creator: string;
    index: string;
    addressList: string;
    address: string;
}
export interface MsgCreateProxyResponse {
}
export interface MsgUpdateProxy {
    creator: string;
    index: string;
    addressList: string;
    address: string;
}
export interface MsgUpdateProxyResponse {
}
export interface MsgDeleteProxy {
    creator: string;
    index: string;
}
export interface MsgDeleteProxyResponse {
}
export declare const MsgCreateProxy: {
    encode(message: MsgCreateProxy, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateProxy;
    fromJSON(object: any): MsgCreateProxy;
    toJSON(message: MsgCreateProxy): unknown;
    fromPartial(object: DeepPartial<MsgCreateProxy>): MsgCreateProxy;
};
export declare const MsgCreateProxyResponse: {
    encode(_: MsgCreateProxyResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateProxyResponse;
    fromJSON(_: any): MsgCreateProxyResponse;
    toJSON(_: MsgCreateProxyResponse): unknown;
    fromPartial(_: DeepPartial<MsgCreateProxyResponse>): MsgCreateProxyResponse;
};
export declare const MsgUpdateProxy: {
    encode(message: MsgUpdateProxy, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateProxy;
    fromJSON(object: any): MsgUpdateProxy;
    toJSON(message: MsgUpdateProxy): unknown;
    fromPartial(object: DeepPartial<MsgUpdateProxy>): MsgUpdateProxy;
};
export declare const MsgUpdateProxyResponse: {
    encode(_: MsgUpdateProxyResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateProxyResponse;
    fromJSON(_: any): MsgUpdateProxyResponse;
    toJSON(_: MsgUpdateProxyResponse): unknown;
    fromPartial(_: DeepPartial<MsgUpdateProxyResponse>): MsgUpdateProxyResponse;
};
export declare const MsgDeleteProxy: {
    encode(message: MsgDeleteProxy, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteProxy;
    fromJSON(object: any): MsgDeleteProxy;
    toJSON(message: MsgDeleteProxy): unknown;
    fromPartial(object: DeepPartial<MsgDeleteProxy>): MsgDeleteProxy;
};
export declare const MsgDeleteProxyResponse: {
    encode(_: MsgDeleteProxyResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteProxyResponse;
    fromJSON(_: any): MsgDeleteProxyResponse;
    toJSON(_: MsgDeleteProxyResponse): unknown;
    fromPartial(_: DeepPartial<MsgDeleteProxyResponse>): MsgDeleteProxyResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    /** this line is used by starport scaffolding # proto/tx/rpc */
    CreateProxy(request: MsgCreateProxy): Promise<MsgCreateProxyResponse>;
    UpdateProxy(request: MsgUpdateProxy): Promise<MsgUpdateProxyResponse>;
    DeleteProxy(request: MsgDeleteProxy): Promise<MsgDeleteProxyResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    CreateProxy(request: MsgCreateProxy): Promise<MsgCreateProxyResponse>;
    UpdateProxy(request: MsgUpdateProxy): Promise<MsgUpdateProxyResponse>;
    DeleteProxy(request: MsgDeleteProxy): Promise<MsgDeleteProxyResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
