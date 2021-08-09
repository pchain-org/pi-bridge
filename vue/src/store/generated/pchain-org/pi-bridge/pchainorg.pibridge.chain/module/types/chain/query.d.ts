import { Reader, Writer } from 'protobufjs/minimal';
import { Chain } from '../chain/chain';
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination';
export declare const protobufPackage = "pchainorg.pibridge.chain";
/** this line is used by starport scaffolding # 3 */
export interface QueryGetChainRequest {
    index: string;
}
export interface QueryGetChainResponse {
    Chain: Chain | undefined;
}
export interface QueryAllChainRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllChainResponse {
    Chain: Chain[];
    pagination: PageResponse | undefined;
}
export declare const QueryGetChainRequest: {
    encode(message: QueryGetChainRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetChainRequest;
    fromJSON(object: any): QueryGetChainRequest;
    toJSON(message: QueryGetChainRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetChainRequest>): QueryGetChainRequest;
};
export declare const QueryGetChainResponse: {
    encode(message: QueryGetChainResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetChainResponse;
    fromJSON(object: any): QueryGetChainResponse;
    toJSON(message: QueryGetChainResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetChainResponse>): QueryGetChainResponse;
};
export declare const QueryAllChainRequest: {
    encode(message: QueryAllChainRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllChainRequest;
    fromJSON(object: any): QueryAllChainRequest;
    toJSON(message: QueryAllChainRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllChainRequest>): QueryAllChainRequest;
};
export declare const QueryAllChainResponse: {
    encode(message: QueryAllChainResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllChainResponse;
    fromJSON(object: any): QueryAllChainResponse;
    toJSON(message: QueryAllChainResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllChainResponse>): QueryAllChainResponse;
};
/** Query defines the gRPC querier service. */
export interface Query {
    /** Queries a chain by index. */
    Chain(request: QueryGetChainRequest): Promise<QueryGetChainResponse>;
    /** Queries a list of chain items. */
    ChainAll(request: QueryAllChainRequest): Promise<QueryAllChainResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    Chain(request: QueryGetChainRequest): Promise<QueryGetChainResponse>;
    ChainAll(request: QueryAllChainRequest): Promise<QueryAllChainResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
