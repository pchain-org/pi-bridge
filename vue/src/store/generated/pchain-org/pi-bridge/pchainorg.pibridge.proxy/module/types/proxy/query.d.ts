import { Reader, Writer } from 'protobufjs/minimal';
import { Proxy } from '../proxy/proxy';
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination';
export declare const protobufPackage = "pchainorg.pibridge.proxy";
/** this line is used by starport scaffolding # 3 */
export interface QueryGetProxyRequest {
    index: string;
}
export interface QueryGetProxyResponse {
    Proxy: Proxy | undefined;
}
export interface QueryAllProxyRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllProxyResponse {
    Proxy: Proxy[];
    pagination: PageResponse | undefined;
}
export declare const QueryGetProxyRequest: {
    encode(message: QueryGetProxyRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetProxyRequest;
    fromJSON(object: any): QueryGetProxyRequest;
    toJSON(message: QueryGetProxyRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetProxyRequest>): QueryGetProxyRequest;
};
export declare const QueryGetProxyResponse: {
    encode(message: QueryGetProxyResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetProxyResponse;
    fromJSON(object: any): QueryGetProxyResponse;
    toJSON(message: QueryGetProxyResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetProxyResponse>): QueryGetProxyResponse;
};
export declare const QueryAllProxyRequest: {
    encode(message: QueryAllProxyRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllProxyRequest;
    fromJSON(object: any): QueryAllProxyRequest;
    toJSON(message: QueryAllProxyRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllProxyRequest>): QueryAllProxyRequest;
};
export declare const QueryAllProxyResponse: {
    encode(message: QueryAllProxyResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllProxyResponse;
    fromJSON(object: any): QueryAllProxyResponse;
    toJSON(message: QueryAllProxyResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllProxyResponse>): QueryAllProxyResponse;
};
/** Query defines the gRPC querier service. */
export interface Query {
    /** Queries a proxy by index. */
    Proxy(request: QueryGetProxyRequest): Promise<QueryGetProxyResponse>;
    /** Queries a list of proxy items. */
    ProxyAll(request: QueryAllProxyRequest): Promise<QueryAllProxyResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    Proxy(request: QueryGetProxyRequest): Promise<QueryGetProxyResponse>;
    ProxyAll(request: QueryAllProxyRequest): Promise<QueryAllProxyResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
