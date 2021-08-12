import { Reader, Writer } from 'protobufjs/minimal';
export declare const protobufPackage = "pchainorg.pibridge.trx";
/** this line is used by starport scaffolding # proto/tx/message */
export interface MsgCrossChainTransfer {
    creator: string;
    asset: string;
    amount: number;
    fromChain: string;
    fromAddress: string;
    toChain: string;
    toAddress: string;
}
export interface MsgCrossChainTransferResponse {
}
export interface MsgCreateTrx {
    creator: string;
    index: string;
    chainID: number;
    address: string;
    crossChainMsgs: string;
    proof: string;
}
export interface MsgCreateTrxResponse {
}
export interface MsgUpdateTrx {
    creator: string;
    index: string;
    chainID: number;
    address: string;
    crossChainMsgs: string;
}
export interface MsgUpdateTrxResponse {
}
export interface MsgDeleteTrx {
    creator: string;
    index: string;
}
export interface MsgDeleteTrxResponse {
}
export declare const MsgCrossChainTransfer: {
    encode(message: MsgCrossChainTransfer, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCrossChainTransfer;
    fromJSON(object: any): MsgCrossChainTransfer;
    toJSON(message: MsgCrossChainTransfer): unknown;
    fromPartial(object: DeepPartial<MsgCrossChainTransfer>): MsgCrossChainTransfer;
};
export declare const MsgCrossChainTransferResponse: {
    encode(_: MsgCrossChainTransferResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCrossChainTransferResponse;
    fromJSON(_: any): MsgCrossChainTransferResponse;
    toJSON(_: MsgCrossChainTransferResponse): unknown;
    fromPartial(_: DeepPartial<MsgCrossChainTransferResponse>): MsgCrossChainTransferResponse;
};
export declare const MsgCreateTrx: {
    encode(message: MsgCreateTrx, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateTrx;
    fromJSON(object: any): MsgCreateTrx;
    toJSON(message: MsgCreateTrx): unknown;
    fromPartial(object: DeepPartial<MsgCreateTrx>): MsgCreateTrx;
};
export declare const MsgCreateTrxResponse: {
    encode(_: MsgCreateTrxResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateTrxResponse;
    fromJSON(_: any): MsgCreateTrxResponse;
    toJSON(_: MsgCreateTrxResponse): unknown;
    fromPartial(_: DeepPartial<MsgCreateTrxResponse>): MsgCreateTrxResponse;
};
export declare const MsgUpdateTrx: {
    encode(message: MsgUpdateTrx, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateTrx;
    fromJSON(object: any): MsgUpdateTrx;
    toJSON(message: MsgUpdateTrx): unknown;
    fromPartial(object: DeepPartial<MsgUpdateTrx>): MsgUpdateTrx;
};
export declare const MsgUpdateTrxResponse: {
    encode(_: MsgUpdateTrxResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateTrxResponse;
    fromJSON(_: any): MsgUpdateTrxResponse;
    toJSON(_: MsgUpdateTrxResponse): unknown;
    fromPartial(_: DeepPartial<MsgUpdateTrxResponse>): MsgUpdateTrxResponse;
};
export declare const MsgDeleteTrx: {
    encode(message: MsgDeleteTrx, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteTrx;
    fromJSON(object: any): MsgDeleteTrx;
    toJSON(message: MsgDeleteTrx): unknown;
    fromPartial(object: DeepPartial<MsgDeleteTrx>): MsgDeleteTrx;
};
export declare const MsgDeleteTrxResponse: {
    encode(_: MsgDeleteTrxResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteTrxResponse;
    fromJSON(_: any): MsgDeleteTrxResponse;
    toJSON(_: MsgDeleteTrxResponse): unknown;
    fromPartial(_: DeepPartial<MsgDeleteTrxResponse>): MsgDeleteTrxResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    /** this line is used by starport scaffolding # proto/tx/rpc */
    CrossChainTransfer(request: MsgCrossChainTransfer): Promise<MsgCrossChainTransferResponse>;
    CreateTrx(request: MsgCreateTrx): Promise<MsgCreateTrxResponse>;
    UpdateTrx(request: MsgUpdateTrx): Promise<MsgUpdateTrxResponse>;
    DeleteTrx(request: MsgDeleteTrx): Promise<MsgDeleteTrxResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    CrossChainTransfer(request: MsgCrossChainTransfer): Promise<MsgCrossChainTransferResponse>;
    CreateTrx(request: MsgCreateTrx): Promise<MsgCreateTrxResponse>;
    UpdateTrx(request: MsgUpdateTrx): Promise<MsgUpdateTrxResponse>;
    DeleteTrx(request: MsgDeleteTrx): Promise<MsgDeleteTrxResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
