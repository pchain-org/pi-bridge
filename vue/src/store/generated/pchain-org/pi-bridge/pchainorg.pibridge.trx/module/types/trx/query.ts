/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal'
import { Trx } from '../trx/trx'
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination'

export const protobufPackage = 'pchainorg.pibridge.trx'

/** this line is used by starport scaffolding # 3 */
export interface QueryGetTrxRequest {
  index: string
}

export interface QueryGetTrxResponse {
  Trx: Trx | undefined
}

export interface QueryAllTrxRequest {
  pagination: PageRequest | undefined
}

export interface QueryAllTrxResponse {
  Trx: Trx[]
  pagination: PageResponse | undefined
}

const baseQueryGetTrxRequest: object = { index: '' }

export const QueryGetTrxRequest = {
  encode(message: QueryGetTrxRequest, writer: Writer = Writer.create()): Writer {
    if (message.index !== '') {
      writer.uint32(10).string(message.index)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetTrxRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetTrxRequest } as QueryGetTrxRequest
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

  fromJSON(object: any): QueryGetTrxRequest {
    const message = { ...baseQueryGetTrxRequest } as QueryGetTrxRequest
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index)
    } else {
      message.index = ''
    }
    return message
  },

  toJSON(message: QueryGetTrxRequest): unknown {
    const obj: any = {}
    message.index !== undefined && (obj.index = message.index)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetTrxRequest>): QueryGetTrxRequest {
    const message = { ...baseQueryGetTrxRequest } as QueryGetTrxRequest
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index
    } else {
      message.index = ''
    }
    return message
  }
}

const baseQueryGetTrxResponse: object = {}

export const QueryGetTrxResponse = {
  encode(message: QueryGetTrxResponse, writer: Writer = Writer.create()): Writer {
    if (message.Trx !== undefined) {
      Trx.encode(message.Trx, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetTrxResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetTrxResponse } as QueryGetTrxResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.Trx = Trx.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetTrxResponse {
    const message = { ...baseQueryGetTrxResponse } as QueryGetTrxResponse
    if (object.Trx !== undefined && object.Trx !== null) {
      message.Trx = Trx.fromJSON(object.Trx)
    } else {
      message.Trx = undefined
    }
    return message
  },

  toJSON(message: QueryGetTrxResponse): unknown {
    const obj: any = {}
    message.Trx !== undefined && (obj.Trx = message.Trx ? Trx.toJSON(message.Trx) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetTrxResponse>): QueryGetTrxResponse {
    const message = { ...baseQueryGetTrxResponse } as QueryGetTrxResponse
    if (object.Trx !== undefined && object.Trx !== null) {
      message.Trx = Trx.fromPartial(object.Trx)
    } else {
      message.Trx = undefined
    }
    return message
  }
}

const baseQueryAllTrxRequest: object = {}

export const QueryAllTrxRequest = {
  encode(message: QueryAllTrxRequest, writer: Writer = Writer.create()): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllTrxRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllTrxRequest } as QueryAllTrxRequest
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

  fromJSON(object: any): QueryAllTrxRequest {
    const message = { ...baseQueryAllTrxRequest } as QueryAllTrxRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllTrxRequest): unknown {
    const obj: any = {}
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllTrxRequest>): QueryAllTrxRequest {
    const message = { ...baseQueryAllTrxRequest } as QueryAllTrxRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryAllTrxResponse: object = {}

export const QueryAllTrxResponse = {
  encode(message: QueryAllTrxResponse, writer: Writer = Writer.create()): Writer {
    for (const v of message.Trx) {
      Trx.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllTrxResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllTrxResponse } as QueryAllTrxResponse
    message.Trx = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.Trx.push(Trx.decode(reader, reader.uint32()))
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

  fromJSON(object: any): QueryAllTrxResponse {
    const message = { ...baseQueryAllTrxResponse } as QueryAllTrxResponse
    message.Trx = []
    if (object.Trx !== undefined && object.Trx !== null) {
      for (const e of object.Trx) {
        message.Trx.push(Trx.fromJSON(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllTrxResponse): unknown {
    const obj: any = {}
    if (message.Trx) {
      obj.Trx = message.Trx.map((e) => (e ? Trx.toJSON(e) : undefined))
    } else {
      obj.Trx = []
    }
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllTrxResponse>): QueryAllTrxResponse {
    const message = { ...baseQueryAllTrxResponse } as QueryAllTrxResponse
    message.Trx = []
    if (object.Trx !== undefined && object.Trx !== null) {
      for (const e of object.Trx) {
        message.Trx.push(Trx.fromPartial(e))
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
  /** Queries a trx by index. */
  Trx(request: QueryGetTrxRequest): Promise<QueryGetTrxResponse>
  /** Queries a list of trx items. */
  TrxAll(request: QueryAllTrxRequest): Promise<QueryAllTrxResponse>
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc
  constructor(rpc: Rpc) {
    this.rpc = rpc
  }
  Trx(request: QueryGetTrxRequest): Promise<QueryGetTrxResponse> {
    const data = QueryGetTrxRequest.encode(request).finish()
    const promise = this.rpc.request('pchainorg.pibridge.trx.Query', 'Trx', data)
    return promise.then((data) => QueryGetTrxResponse.decode(new Reader(data)))
  }

  TrxAll(request: QueryAllTrxRequest): Promise<QueryAllTrxResponse> {
    const data = QueryAllTrxRequest.encode(request).finish()
    const promise = this.rpc.request('pchainorg.pibridge.trx.Query', 'TrxAll', data)
    return promise.then((data) => QueryAllTrxResponse.decode(new Reader(data)))
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
