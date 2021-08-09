/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal'
import { Node } from '../node/node'
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination'

export const protobufPackage = 'pchainorg.pibridge.node'

/** this line is used by starport scaffolding # 3 */
export interface QueryGetNodeRequest {
  index: string
}

export interface QueryGetNodeResponse {
  Node: Node | undefined
}

export interface QueryAllNodeRequest {
  pagination: PageRequest | undefined
}

export interface QueryAllNodeResponse {
  Node: Node[]
  pagination: PageResponse | undefined
}

const baseQueryGetNodeRequest: object = { index: '' }

export const QueryGetNodeRequest = {
  encode(message: QueryGetNodeRequest, writer: Writer = Writer.create()): Writer {
    if (message.index !== '') {
      writer.uint32(10).string(message.index)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetNodeRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetNodeRequest } as QueryGetNodeRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetNodeRequest {
    const message = { ...baseQueryGetNodeRequest } as QueryGetNodeRequest
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index)
    } else {
      message.index = ''
    }
    return message
  },

  toJSON(message: QueryGetNodeRequest): unknown {
    const obj: any = {}
    message.index !== undefined && (obj.index = message.index)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetNodeRequest>): QueryGetNodeRequest {
    const message = { ...baseQueryGetNodeRequest } as QueryGetNodeRequest
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index
    } else {
      message.index = ''
    }
    return message
  }
}

const baseQueryGetNodeResponse: object = {}

export const QueryGetNodeResponse = {
  encode(message: QueryGetNodeResponse, writer: Writer = Writer.create()): Writer {
    if (message.Node !== undefined) {
      Node.encode(message.Node, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetNodeResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetNodeResponse } as QueryGetNodeResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.Node = Node.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetNodeResponse {
    const message = { ...baseQueryGetNodeResponse } as QueryGetNodeResponse
    if (object.Node !== undefined && object.Node !== null) {
      message.Node = Node.fromJSON(object.Node)
    } else {
      message.Node = undefined
    }
    return message
  },

  toJSON(message: QueryGetNodeResponse): unknown {
    const obj: any = {}
    message.Node !== undefined && (obj.Node = message.Node ? Node.toJSON(message.Node) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetNodeResponse>): QueryGetNodeResponse {
    const message = { ...baseQueryGetNodeResponse } as QueryGetNodeResponse
    if (object.Node !== undefined && object.Node !== null) {
      message.Node = Node.fromPartial(object.Node)
    } else {
      message.Node = undefined
    }
    return message
  }
}

const baseQueryAllNodeRequest: object = {}

export const QueryAllNodeRequest = {
  encode(message: QueryAllNodeRequest, writer: Writer = Writer.create()): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllNodeRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllNodeRequest } as QueryAllNodeRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryAllNodeRequest {
    const message = { ...baseQueryAllNodeRequest } as QueryAllNodeRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllNodeRequest): unknown {
    const obj: any = {}
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllNodeRequest>): QueryAllNodeRequest {
    const message = { ...baseQueryAllNodeRequest } as QueryAllNodeRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryAllNodeResponse: object = {}

export const QueryAllNodeResponse = {
  encode(message: QueryAllNodeResponse, writer: Writer = Writer.create()): Writer {
    for (const v of message.Node) {
      Node.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllNodeResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllNodeResponse } as QueryAllNodeResponse
    message.Node = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.Node.push(Node.decode(reader, reader.uint32()))
          break
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryAllNodeResponse {
    const message = { ...baseQueryAllNodeResponse } as QueryAllNodeResponse
    message.Node = []
    if (object.Node !== undefined && object.Node !== null) {
      for (const e of object.Node) {
        message.Node.push(Node.fromJSON(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllNodeResponse): unknown {
    const obj: any = {}
    if (message.Node) {
      obj.Node = message.Node.map((e) => (e ? Node.toJSON(e) : undefined))
    } else {
      obj.Node = []
    }
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllNodeResponse>): QueryAllNodeResponse {
    const message = { ...baseQueryAllNodeResponse } as QueryAllNodeResponse
    message.Node = []
    if (object.Node !== undefined && object.Node !== null) {
      for (const e of object.Node) {
        message.Node.push(Node.fromPartial(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

/** Query defines the gRPC querier service. */
export interface Query {
  /** Queries a node by index. */
  Node(request: QueryGetNodeRequest): Promise<QueryGetNodeResponse>
  /** Queries a list of node items. */
  NodeAll(request: QueryAllNodeRequest): Promise<QueryAllNodeResponse>
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc
  constructor(rpc: Rpc) {
    this.rpc = rpc
  }
  Node(request: QueryGetNodeRequest): Promise<QueryGetNodeResponse> {
    const data = QueryGetNodeRequest.encode(request).finish()
    const promise = this.rpc.request('pchainorg.pibridge.node.Query', 'Node', data)
    return promise.then((data) => QueryGetNodeResponse.decode(new Reader(data)))
  }

  NodeAll(request: QueryAllNodeRequest): Promise<QueryAllNodeResponse> {
    const data = QueryAllNodeRequest.encode(request).finish()
    const promise = this.rpc.request('pchainorg.pibridge.node.Query', 'NodeAll', data)
    return promise.then((data) => QueryAllNodeResponse.decode(new Reader(data)))
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>
}

type Builtin = Date | Function | Uint8Array | string | number | undefined
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>
