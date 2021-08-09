import { Reader, Writer } from 'protobufjs/minimal';
import { Block } from '../block/block';
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination';
export declare const protobufPackage = "pchainorg.pibridge.block";
/** this line is used by starport scaffolding # 3 */
export interface QueryGetBlockRequest {
    index: string;
}
export interface QueryGetBlockResponse {
    Block: Block | undefined;
}
export interface QueryAllBlockRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllBlockResponse {
    Block: Block[];
    pagination: PageResponse | undefined;
}
export declare const QueryGetBlockRequest: {
    encode(message: QueryGetBlockRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetBlockRequest;
    fromJSON(object: any): QueryGetBlockRequest;
    toJSON(message: QueryGetBlockRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetBlockRequest>): QueryGetBlockRequest;
};
export declare const QueryGetBlockResponse: {
    encode(message: QueryGetBlockResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetBlockResponse;
    fromJSON(object: any): QueryGetBlockResponse;
    toJSON(message: QueryGetBlockResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetBlockResponse>): QueryGetBlockResponse;
};
export declare const QueryAllBlockRequest: {
    encode(message: QueryAllBlockRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllBlockRequest;
    fromJSON(object: any): QueryAllBlockRequest;
    toJSON(message: QueryAllBlockRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllBlockRequest>): QueryAllBlockRequest;
};
export declare const QueryAllBlockResponse: {
    encode(message: QueryAllBlockResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllBlockResponse;
    fromJSON(object: any): QueryAllBlockResponse;
    toJSON(message: QueryAllBlockResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllBlockResponse>): QueryAllBlockResponse;
};
/** Query defines the gRPC querier service. */
export interface Query {
    /** Queries a block by index. */
    Block(request: QueryGetBlockRequest): Promise<QueryGetBlockResponse>;
    /** Queries a list of block items. */
    BlockAll(request: QueryAllBlockRequest): Promise<QueryAllBlockResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    Block(request: QueryGetBlockRequest): Promise<QueryGetBlockResponse>;
    BlockAll(request: QueryAllBlockRequest): Promise<QueryAllBlockResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
