import { Reader, Writer } from 'protobufjs/minimal';
import { Trx } from '../trx/trx';
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination';
export declare const protobufPackage = "pchainorg.pibridge.trx";
/** this line is used by starport scaffolding # 3 */
export interface QueryGetTrxRequest {
    index: string;
}
export interface QueryGetTrxResponse {
    Trx: Trx | undefined;
}
export interface QueryAllTrxRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllTrxResponse {
    Trx: Trx[];
    pagination: PageResponse | undefined;
}
export declare const QueryGetTrxRequest: {
    encode(message: QueryGetTrxRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetTrxRequest;
    fromJSON(object: any): QueryGetTrxRequest;
    toJSON(message: QueryGetTrxRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetTrxRequest>): QueryGetTrxRequest;
};
export declare const QueryGetTrxResponse: {
    encode(message: QueryGetTrxResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetTrxResponse;
    fromJSON(object: any): QueryGetTrxResponse;
    toJSON(message: QueryGetTrxResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetTrxResponse>): QueryGetTrxResponse;
};
export declare const QueryAllTrxRequest: {
    encode(message: QueryAllTrxRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllTrxRequest;
    fromJSON(object: any): QueryAllTrxRequest;
    toJSON(message: QueryAllTrxRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllTrxRequest>): QueryAllTrxRequest;
};
export declare const QueryAllTrxResponse: {
    encode(message: QueryAllTrxResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllTrxResponse;
    fromJSON(object: any): QueryAllTrxResponse;
    toJSON(message: QueryAllTrxResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllTrxResponse>): QueryAllTrxResponse;
};
/** Query defines the gRPC querier service. */
export interface Query {
    /** Queries a trx by index. */
    Trx(request: QueryGetTrxRequest): Promise<QueryGetTrxResponse>;
    /** Queries a list of trx items. */
    TrxAll(request: QueryAllTrxRequest): Promise<QueryAllTrxResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    Trx(request: QueryGetTrxRequest): Promise<QueryGetTrxResponse>;
    TrxAll(request: QueryAllTrxRequest): Promise<QueryAllTrxResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
