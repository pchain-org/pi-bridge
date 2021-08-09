/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal'
import { Chain } from '../chain/chain'
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination'

export const protobufPackage = 'pchainorg.pibridge.chain'

/** this line is used by starport scaffolding # 3 */
export interface QueryGetChainRequest {
  index: string
}

export interface QueryGetChainResponse {
  Chain: Chain | undefined
}

export interface QueryAllChainRequest {
  pagination: PageRequest | undefined
}

export interface QueryAllChainResponse {
  Chain: Chain[]
  pagination: PageResponse | undefined
}

const baseQueryGetChainRequest: object = { index: '' }

export const QueryGetChainRequest = {
  encode(message: QueryGetChainRequest, writer: Writer = Writer.create()): Writer {
    if (message.index !== '') {
      writer.uint32(10).string(message.index)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetChainRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetChainRequest } as QueryGetChainRequest
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

  fromJSON(object: any): QueryGetChainRequest {
    const message = { ...baseQueryGetChainRequest } as QueryGetChainRequest
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index)
    } else {
      message.index = ''
    }
    return message
  },

  toJSON(message: QueryGetChainRequest): unknown {
    const obj: any = {}
    message.index !== undefined && (obj.index = message.index)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetChainRequest>): QueryGetChainRequest {
    const message = { ...baseQueryGetChainRequest } as QueryGetChainRequest
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index
    } else {
      message.index = ''
    }
    return message
  }
}

const baseQueryGetChainResponse: object = {}

export const QueryGetChainResponse = {
  encode(message: QueryGetChainResponse, writer: Writer = Writer.create()): Writer {
    if (message.Chain !== undefined) {
      Chain.encode(message.Chain, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetChainResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetChainResponse } as QueryGetChainResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.Chain = Chain.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetChainResponse {
    const message = { ...baseQueryGetChainResponse } as QueryGetChainResponse
    if (object.Chain !== undefined && object.Chain !== null) {
      message.Chain = Chain.fromJSON(object.Chain)
    } else {
      message.Chain = undefined
    }
    return message
  },

  toJSON(message: QueryGetChainResponse): unknown {
    const obj: any = {}
    message.Chain !== undefined && (obj.Chain = message.Chain ? Chain.toJSON(message.Chain) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetChainResponse>): QueryGetChainResponse {
    const message = { ...baseQueryGetChainResponse } as QueryGetChainResponse
    if (object.Chain !== undefined && object.Chain !== null) {
      message.Chain = Chain.fromPartial(object.Chain)
    } else {
      message.Chain = undefined
    }
    return message
  }
}

const baseQueryAllChainRequest: object = {}

export const QueryAllChainRequest = {
  encode(message: QueryAllChainRequest, writer: Writer = Writer.create()): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllChainRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllChainRequest } as QueryAllChainRequest
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

  fromJSON(object: any): QueryAllChainRequest {
    const message = { ...baseQueryAllChainRequest } as QueryAllChainRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllChainRequest): unknown {
    const obj: any = {}
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllChainRequest>): QueryAllChainRequest {
    const message = { ...baseQueryAllChainRequest } as QueryAllChainRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryAllChainResponse: object = {}

export const QueryAllChainResponse = {
  encode(message: QueryAllChainResponse, writer: Writer = Writer.create()): Writer {
    for (const v of message.Chain) {
      Chain.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllChainResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllChainResponse } as QueryAllChainResponse
    message.Chain = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.Chain.push(Chain.decode(reader, reader.uint32()))
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

  fromJSON(object: any): QueryAllChainResponse {
    const message = { ...baseQueryAllChainResponse } as QueryAllChainResponse
    message.Chain = []
    if (object.Chain !== undefined && object.Chain !== null) {
      for (const e of object.Chain) {
        message.Chain.push(Chain.fromJSON(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllChainResponse): unknown {
    const obj: any = {}
    if (message.Chain) {
      obj.Chain = message.Chain.map((e) => (e ? Chain.toJSON(e) : undefined))
    } else {
      obj.Chain = []
    }
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllChainResponse>): QueryAllChainResponse {
    const message = { ...baseQueryAllChainResponse } as QueryAllChainResponse
    message.Chain = []
    if (object.Chain !== undefined && object.Chain !== null) {
      for (const e of object.Chain) {
        message.Chain.push(Chain.fromPartial(e))
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
  /** Queries a chain by index. */
  Chain(request: QueryGetChainRequest): Promise<QueryGetChainResponse>
  /** Queries a list of chain items. */
  ChainAll(request: QueryAllChainRequest): Promise<QueryAllChainResponse>
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc
  constructor(rpc: Rpc) {
    this.rpc = rpc
  }
  Chain(request: QueryGetChainRequest): Promise<QueryGetChainResponse> {
    const data = QueryGetChainRequest.encode(request).finish()
    const promise = this.rpc.request('pchainorg.pibridge.chain.Query', 'Chain', data)
    return promise.then((data) => QueryGetChainResponse.decode(new Reader(data)))
  }

  ChainAll(request: QueryAllChainRequest): Promise<QueryAllChainResponse> {
    const data = QueryAllChainRequest.encode(request).finish()
    const promise = this.rpc.request('pchainorg.pibridge.chain.Query', 'ChainAll', data)
    return promise.then((data) => QueryAllChainResponse.decode(new Reader(data)))
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
