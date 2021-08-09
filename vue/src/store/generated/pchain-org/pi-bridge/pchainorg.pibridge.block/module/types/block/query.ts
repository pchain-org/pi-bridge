/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal'
import { Block } from '../block/block'
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination'

export const protobufPackage = 'pchainorg.pibridge.block'

/** this line is used by starport scaffolding # 3 */
export interface QueryGetBlockRequest {
  index: string
}

export interface QueryGetBlockResponse {
  Block: Block | undefined
}

export interface QueryAllBlockRequest {
  pagination: PageRequest | undefined
}

export interface QueryAllBlockResponse {
  Block: Block[]
  pagination: PageResponse | undefined
}

const baseQueryGetBlockRequest: object = { index: '' }

export const QueryGetBlockRequest = {
  encode(message: QueryGetBlockRequest, writer: Writer = Writer.create()): Writer {
    if (message.index !== '') {
      writer.uint32(10).string(message.index)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetBlockRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetBlockRequest } as QueryGetBlockRequest
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

  fromJSON(object: any): QueryGetBlockRequest {
    const message = { ...baseQueryGetBlockRequest } as QueryGetBlockRequest
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index)
    } else {
      message.index = ''
    }
    return message
  },

  toJSON(message: QueryGetBlockRequest): unknown {
    const obj: any = {}
    message.index !== undefined && (obj.index = message.index)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetBlockRequest>): QueryGetBlockRequest {
    const message = { ...baseQueryGetBlockRequest } as QueryGetBlockRequest
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index
    } else {
      message.index = ''
    }
    return message
  }
}

const baseQueryGetBlockResponse: object = {}

export const QueryGetBlockResponse = {
  encode(message: QueryGetBlockResponse, writer: Writer = Writer.create()): Writer {
    if (message.Block !== undefined) {
      Block.encode(message.Block, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetBlockResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetBlockResponse } as QueryGetBlockResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.Block = Block.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetBlockResponse {
    const message = { ...baseQueryGetBlockResponse } as QueryGetBlockResponse
    if (object.Block !== undefined && object.Block !== null) {
      message.Block = Block.fromJSON(object.Block)
    } else {
      message.Block = undefined
    }
    return message
  },

  toJSON(message: QueryGetBlockResponse): unknown {
    const obj: any = {}
    message.Block !== undefined && (obj.Block = message.Block ? Block.toJSON(message.Block) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetBlockResponse>): QueryGetBlockResponse {
    const message = { ...baseQueryGetBlockResponse } as QueryGetBlockResponse
    if (object.Block !== undefined && object.Block !== null) {
      message.Block = Block.fromPartial(object.Block)
    } else {
      message.Block = undefined
    }
    return message
  }
}

const baseQueryAllBlockRequest: object = {}

export const QueryAllBlockRequest = {
  encode(message: QueryAllBlockRequest, writer: Writer = Writer.create()): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllBlockRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllBlockRequest } as QueryAllBlockRequest
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

  fromJSON(object: any): QueryAllBlockRequest {
    const message = { ...baseQueryAllBlockRequest } as QueryAllBlockRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllBlockRequest): unknown {
    const obj: any = {}
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllBlockRequest>): QueryAllBlockRequest {
    const message = { ...baseQueryAllBlockRequest } as QueryAllBlockRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryAllBlockResponse: object = {}

export const QueryAllBlockResponse = {
  encode(message: QueryAllBlockResponse, writer: Writer = Writer.create()): Writer {
    for (const v of message.Block) {
      Block.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllBlockResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllBlockResponse } as QueryAllBlockResponse
    message.Block = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.Block.push(Block.decode(reader, reader.uint32()))
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

  fromJSON(object: any): QueryAllBlockResponse {
    const message = { ...baseQueryAllBlockResponse } as QueryAllBlockResponse
    message.Block = []
    if (object.Block !== undefined && object.Block !== null) {
      for (const e of object.Block) {
        message.Block.push(Block.fromJSON(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllBlockResponse): unknown {
    const obj: any = {}
    if (message.Block) {
      obj.Block = message.Block.map((e) => (e ? Block.toJSON(e) : undefined))
    } else {
      obj.Block = []
    }
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllBlockResponse>): QueryAllBlockResponse {
    const message = { ...baseQueryAllBlockResponse } as QueryAllBlockResponse
    message.Block = []
    if (object.Block !== undefined && object.Block !== null) {
      for (const e of object.Block) {
        message.Block.push(Block.fromPartial(e))
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
  /** Queries a block by index. */
  Block(request: QueryGetBlockRequest): Promise<QueryGetBlockResponse>
  /** Queries a list of block items. */
  BlockAll(request: QueryAllBlockRequest): Promise<QueryAllBlockResponse>
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc
  constructor(rpc: Rpc) {
    this.rpc = rpc
  }
  Block(request: QueryGetBlockRequest): Promise<QueryGetBlockResponse> {
    const data = QueryGetBlockRequest.encode(request).finish()
    const promise = this.rpc.request('pchainorg.pibridge.block.Query', 'Block', data)
    return promise.then((data) => QueryGetBlockResponse.decode(new Reader(data)))
  }

  BlockAll(request: QueryAllBlockRequest): Promise<QueryAllBlockResponse> {
    const data = QueryAllBlockRequest.encode(request).finish()
    const promise = this.rpc.request('pchainorg.pibridge.block.Query', 'BlockAll', data)
    return promise.then((data) => QueryAllBlockResponse.decode(new Reader(data)))
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
