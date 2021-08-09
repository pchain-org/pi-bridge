/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal'

export const protobufPackage = 'pchainorg.pibridge.proxy'

/** this line is used by starport scaffolding # proto/tx/message */
export interface MsgCreateProxy {
  creator: string
  index: string
  addressList: string
  address: string
}

export interface MsgCreateProxyResponse {}

export interface MsgUpdateProxy {
  creator: string
  index: string
  addressList: string
  address: string
}

export interface MsgUpdateProxyResponse {}

export interface MsgDeleteProxy {
  creator: string
  index: string
}

export interface MsgDeleteProxyResponse {}

const baseMsgCreateProxy: object = { creator: '', index: '', addressList: '', address: '' }

export const MsgCreateProxy = {
  encode(message: MsgCreateProxy, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.index !== '') {
      writer.uint32(18).string(message.index)
    }
    if (message.addressList !== '') {
      writer.uint32(26).string(message.addressList)
    }
    if (message.address !== '') {
      writer.uint32(34).string(message.address)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateProxy {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgCreateProxy } as MsgCreateProxy
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string()
          break
        case 2:
          message.index = reader.string()
          break
        case 3:
          message.addressList = reader.string()
          break
        case 4:
          message.address = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgCreateProxy {
    const message = { ...baseMsgCreateProxy } as MsgCreateProxy
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index)
    } else {
      message.index = ''
    }
    if (object.addressList !== undefined && object.addressList !== null) {
      message.addressList = String(object.addressList)
    } else {
      message.addressList = ''
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address)
    } else {
      message.address = ''
    }
    return message
  },

  toJSON(message: MsgCreateProxy): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.index !== undefined && (obj.index = message.index)
    message.addressList !== undefined && (obj.addressList = message.addressList)
    message.address !== undefined && (obj.address = message.address)
    return obj
  },

  fromPartial(object: DeepPartial<MsgCreateProxy>): MsgCreateProxy {
    const message = { ...baseMsgCreateProxy } as MsgCreateProxy
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator
    } else {
      message.creator = ''
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index
    } else {
      message.index = ''
    }
    if (object.addressList !== undefined && object.addressList !== null) {
      message.addressList = object.addressList
    } else {
      message.addressList = ''
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address
    } else {
      message.address = ''
    }
    return message
  }
}

const baseMsgCreateProxyResponse: object = {}

export const MsgCreateProxyResponse = {
  encode(_: MsgCreateProxyResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateProxyResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgCreateProxyResponse } as MsgCreateProxyResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(_: any): MsgCreateProxyResponse {
    const message = { ...baseMsgCreateProxyResponse } as MsgCreateProxyResponse
    return message
  },

  toJSON(_: MsgCreateProxyResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgCreateProxyResponse>): MsgCreateProxyResponse {
    const message = { ...baseMsgCreateProxyResponse } as MsgCreateProxyResponse
    return message
  }
}

const baseMsgUpdateProxy: object = { creator: '', index: '', addressList: '', address: '' }

export const MsgUpdateProxy = {
  encode(message: MsgUpdateProxy, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.index !== '') {
      writer.uint32(18).string(message.index)
    }
    if (message.addressList !== '') {
      writer.uint32(26).string(message.addressList)
    }
    if (message.address !== '') {
      writer.uint32(34).string(message.address)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateProxy {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgUpdateProxy } as MsgUpdateProxy
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string()
          break
        case 2:
          message.index = reader.string()
          break
        case 3:
          message.addressList = reader.string()
          break
        case 4:
          message.address = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgUpdateProxy {
    const message = { ...baseMsgUpdateProxy } as MsgUpdateProxy
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index)
    } else {
      message.index = ''
    }
    if (object.addressList !== undefined && object.addressList !== null) {
      message.addressList = String(object.addressList)
    } else {
      message.addressList = ''
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address)
    } else {
      message.address = ''
    }
    return message
  },

  toJSON(message: MsgUpdateProxy): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.index !== undefined && (obj.index = message.index)
    message.addressList !== undefined && (obj.addressList = message.addressList)
    message.address !== undefined && (obj.address = message.address)
    return obj
  },

  fromPartial(object: DeepPartial<MsgUpdateProxy>): MsgUpdateProxy {
    const message = { ...baseMsgUpdateProxy } as MsgUpdateProxy
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator
    } else {
      message.creator = ''
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index
    } else {
      message.index = ''
    }
    if (object.addressList !== undefined && object.addressList !== null) {
      message.addressList = object.addressList
    } else {
      message.addressList = ''
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address
    } else {
      message.address = ''
    }
    return message
  }
}

const baseMsgUpdateProxyResponse: object = {}

export const MsgUpdateProxyResponse = {
  encode(_: MsgUpdateProxyResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateProxyResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgUpdateProxyResponse } as MsgUpdateProxyResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(_: any): MsgUpdateProxyResponse {
    const message = { ...baseMsgUpdateProxyResponse } as MsgUpdateProxyResponse
    return message
  },

  toJSON(_: MsgUpdateProxyResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgUpdateProxyResponse>): MsgUpdateProxyResponse {
    const message = { ...baseMsgUpdateProxyResponse } as MsgUpdateProxyResponse
    return message
  }
}

const baseMsgDeleteProxy: object = { creator: '', index: '' }

export const MsgDeleteProxy = {
  encode(message: MsgDeleteProxy, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.index !== '') {
      writer.uint32(18).string(message.index)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteProxy {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgDeleteProxy } as MsgDeleteProxy
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string()
          break
        case 2:
          message.index = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgDeleteProxy {
    const message = { ...baseMsgDeleteProxy } as MsgDeleteProxy
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index)
    } else {
      message.index = ''
    }
    return message
  },

  toJSON(message: MsgDeleteProxy): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.index !== undefined && (obj.index = message.index)
    return obj
  },

  fromPartial(object: DeepPartial<MsgDeleteProxy>): MsgDeleteProxy {
    const message = { ...baseMsgDeleteProxy } as MsgDeleteProxy
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator
    } else {
      message.creator = ''
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index
    } else {
      message.index = ''
    }
    return message
  }
}

const baseMsgDeleteProxyResponse: object = {}

export const MsgDeleteProxyResponse = {
  encode(_: MsgDeleteProxyResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteProxyResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgDeleteProxyResponse } as MsgDeleteProxyResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(_: any): MsgDeleteProxyResponse {
    const message = { ...baseMsgDeleteProxyResponse } as MsgDeleteProxyResponse
    return message
  },

  toJSON(_: MsgDeleteProxyResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgDeleteProxyResponse>): MsgDeleteProxyResponse {
    const message = { ...baseMsgDeleteProxyResponse } as MsgDeleteProxyResponse
    return message
  }
}

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  CreateProxy(request: MsgCreateProxy): Promise<MsgCreateProxyResponse>
  UpdateProxy(request: MsgUpdateProxy): Promise<MsgUpdateProxyResponse>
  DeleteProxy(request: MsgDeleteProxy): Promise<MsgDeleteProxyResponse>
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc
  constructor(rpc: Rpc) {
    this.rpc = rpc
  }
  CreateProxy(request: MsgCreateProxy): Promise<MsgCreateProxyResponse> {
    const data = MsgCreateProxy.encode(request).finish()
    const promise = this.rpc.request('pchainorg.pibridge.proxy.Msg', 'CreateProxy', data)
    return promise.then((data) => MsgCreateProxyResponse.decode(new Reader(data)))
  }

  UpdateProxy(request: MsgUpdateProxy): Promise<MsgUpdateProxyResponse> {
    const data = MsgUpdateProxy.encode(request).finish()
    const promise = this.rpc.request('pchainorg.pibridge.proxy.Msg', 'UpdateProxy', data)
    return promise.then((data) => MsgUpdateProxyResponse.decode(new Reader(data)))
  }

  DeleteProxy(request: MsgDeleteProxy): Promise<MsgDeleteProxyResponse> {
    const data = MsgDeleteProxy.encode(request).finish()
    const promise = this.rpc.request('pchainorg.pibridge.proxy.Msg', 'DeleteProxy', data)
    return promise.then((data) => MsgDeleteProxyResponse.decode(new Reader(data)))
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
