/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal'
import { Proxy } from '../proxy/proxy'
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination'

export const protobufPackage = 'pchainorg.pibridge.proxy'

/** this line is used by starport scaffolding # 3 */
export interface QueryGetProxyRequest {
  index: string
}

export interface QueryGetProxyResponse {
  Proxy: Proxy | undefined
}

export interface QueryAllProxyRequest {
  pagination: PageRequest | undefined
}

export interface QueryAllProxyResponse {
  Proxy: Proxy[]
  pagination: PageResponse | undefined
}

const baseQueryGetProxyRequest: object = { index: '' }

export const QueryGetProxyRequest = {
  encode(message: QueryGetProxyRequest, writer: Writer = Writer.create()): Writer {
    if (message.index !== '') {
      writer.uint32(10).string(message.index)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetProxyRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetProxyRequest } as QueryGetProxyRequest
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

  fromJSON(object: any): QueryGetProxyRequest {
    const message = { ...baseQueryGetProxyRequest } as QueryGetProxyRequest
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index)
    } else {
      message.index = ''
    }
    return message
  },

  toJSON(message: QueryGetProxyRequest): unknown {
    const obj: any = {}
    message.index !== undefined && (obj.index = message.index)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetProxyRequest>): QueryGetProxyRequest {
    const message = { ...baseQueryGetProxyRequest } as QueryGetProxyRequest
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index
    } else {
      message.index = ''
    }
    return message
  }
}

const baseQueryGetProxyResponse: object = {}

export const QueryGetProxyResponse = {
  encode(message: QueryGetProxyResponse, writer: Writer = Writer.create()): Writer {
    if (message.Proxy !== undefined) {
      Proxy.encode(message.Proxy, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetProxyResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetProxyResponse } as QueryGetProxyResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.Proxy = Proxy.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetProxyResponse {
    const message = { ...baseQueryGetProxyResponse } as QueryGetProxyResponse
    if (object.Proxy !== undefined && object.Proxy !== null) {
      message.Proxy = Proxy.fromJSON(object.Proxy)
    } else {
      message.Proxy = undefined
    }
    return message
  },

  toJSON(message: QueryGetProxyResponse): unknown {
    const obj: any = {}
    message.Proxy !== undefined && (obj.Proxy = message.Proxy ? Proxy.toJSON(message.Proxy) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetProxyResponse>): QueryGetProxyResponse {
    const message = { ...baseQueryGetProxyResponse } as QueryGetProxyResponse
    if (object.Proxy !== undefined && object.Proxy !== null) {
      message.Proxy = Proxy.fromPartial(object.Proxy)
    } else {
      message.Proxy = undefined
    }
    return message
  }
}

const baseQueryAllProxyRequest: object = {}

export const QueryAllProxyRequest = {
  encode(message: QueryAllProxyRequest, writer: Writer = Writer.create()): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllProxyRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllProxyRequest } as QueryAllProxyRequest
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

  fromJSON(object: any): QueryAllProxyRequest {
    const message = { ...baseQueryAllProxyRequest } as QueryAllProxyRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllProxyRequest): unknown {
    const obj: any = {}
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllProxyRequest>): QueryAllProxyRequest {
    const message = { ...baseQueryAllProxyRequest } as QueryAllProxyRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryAllProxyResponse: object = {}

export const QueryAllProxyResponse = {
  encode(message: QueryAllProxyResponse, writer: Writer = Writer.create()): Writer {
    for (const v of message.Proxy) {
      Proxy.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllProxyResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllProxyResponse } as QueryAllProxyResponse
    message.Proxy = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.Proxy.push(Proxy.decode(reader, reader.uint32()))
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

  fromJSON(object: any): QueryAllProxyResponse {
    const message = { ...baseQueryAllProxyResponse } as QueryAllProxyResponse
    message.Proxy = []
    if (object.Proxy !== undefined && object.Proxy !== null) {
      for (const e of object.Proxy) {
        message.Proxy.push(Proxy.fromJSON(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllProxyResponse): unknown {
    const obj: any = {}
    if (message.Proxy) {
      obj.Proxy = message.Proxy.map((e) => (e ? Proxy.toJSON(e) : undefined))
    } else {
      obj.Proxy = []
    }
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllProxyResponse>): QueryAllProxyResponse {
    const message = { ...baseQueryAllProxyResponse } as QueryAllProxyResponse
    message.Proxy = []
    if (object.Proxy !== undefined && object.Proxy !== null) {
      for (const e of object.Proxy) {
        message.Proxy.push(Proxy.fromPartial(e))
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
  /** Queries a proxy by index. */
  Proxy(request: QueryGetProxyRequest): Promise<QueryGetProxyResponse>
  /** Queries a list of proxy items. */
  ProxyAll(request: QueryAllProxyRequest): Promise<QueryAllProxyResponse>
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc
  constructor(rpc: Rpc) {
    this.rpc = rpc
  }
  Proxy(request: QueryGetProxyRequest): Promise<QueryGetProxyResponse> {
    const data = QueryGetProxyRequest.encode(request).finish()
    const promise = this.rpc.request('pchainorg.pibridge.proxy.Query', 'Proxy', data)
    return promise.then((data) => QueryGetProxyResponse.decode(new Reader(data)))
  }

  ProxyAll(request: QueryAllProxyRequest): Promise<QueryAllProxyResponse> {
    const data = QueryAllProxyRequest.encode(request).finish()
    const promise = this.rpc.request('pchainorg.pibridge.proxy.Query', 'ProxyAll', data)
    return promise.then((data) => QueryAllProxyResponse.decode(new Reader(data)))
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
