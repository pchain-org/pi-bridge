import { Reader, Writer } from 'protobufjs/minimal';
import { Node } from '../node/node';
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination';
export declare const protobufPackage = "pchainorg.pibridge.node";
/** this line is used by starport scaffolding # 3 */
export interface QueryGetNodeRequest {
    index: string;
}
export interface QueryGetNodeResponse {
    Node: Node | undefined;
}
export interface QueryAllNodeRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllNodeResponse {
    Node: Node[];
    pagination: PageResponse | undefined;
}
export declare const QueryGetNodeRequest: {
    encode(message: QueryGetNodeRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetNodeRequest;
    fromJSON(object: any): QueryGetNodeRequest;
    toJSON(message: QueryGetNodeRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetNodeRequest>): QueryGetNodeRequest;
};
export declare const QueryGetNodeResponse: {
    encode(message: QueryGetNodeResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetNodeResponse;
    fromJSON(object: any): QueryGetNodeResponse;
    toJSON(message: QueryGetNodeResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetNodeResponse>): QueryGetNodeResponse;
};
export declare const QueryAllNodeRequest: {
    encode(message: QueryAllNodeRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllNodeRequest;
    fromJSON(object: any): QueryAllNodeRequest;
    toJSON(message: QueryAllNodeRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllNodeRequest>): QueryAllNodeRequest;
};
export declare const QueryAllNodeResponse: {
    encode(message: QueryAllNodeResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllNodeResponse;
    fromJSON(object: any): QueryAllNodeResponse;
    toJSON(message: QueryAllNodeResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllNodeResponse>): QueryAllNodeResponse;
};
/** Query defines the gRPC querier service. */
export interface Query {
    /** Queries a node by index. */
    Node(request: QueryGetNodeRequest): Promise<QueryGetNodeResponse>;
    /** Queries a list of node items. */
    NodeAll(request: QueryAllNodeRequest): Promise<QueryAllNodeResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    Node(request: QueryGetNodeRequest): Promise<QueryGetNodeResponse>;
    NodeAll(request: QueryAllNodeRequest): Promise<QueryAllNodeResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
