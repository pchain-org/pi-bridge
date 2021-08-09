/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal'

export const protobufPackage = 'pchainorg.pibridge.chain'

/** this line is used by starport scaffolding # proto/tx/message */
export interface MsgCreateChain {
  creator: string
  index: string
  sourceChainID: number
  height: number
  proof: string
  proxyAddress: string
  extra: string
  headerOrCrossChainMsg: string
}

export interface MsgCreateChainResponse {}

export interface MsgUpdateChain {
  creator: string
  index: string
  sourceChainID: number
  height: number
  proof: string
  proxyAddress: string
  extra: string
  headerOrCrossChainMsg: string
}

export interface MsgUpdateChainResponse {}

export interface MsgDeleteChain {
  creator: string
  index: string
}

export interface MsgDeleteChainResponse {}

const baseMsgCreateChain: object = { creator: '', index: '', sourceChainID: 0, height: 0, proof: '', proxyAddress: '', extra: '', headerOrCrossChainMsg: '' }

export const MsgCreateChain = {
  encode(message: MsgCreateChain, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.index !== '') {
      writer.uint32(18).string(message.index)
    }
    if (message.sourceChainID !== 0) {
      writer.uint32(24).int32(message.sourceChainID)
    }
    if (message.height !== 0) {
      writer.uint32(32).int32(message.height)
    }
    if (message.proof !== '') {
      writer.uint32(42).string(message.proof)
    }
    if (message.proxyAddress !== '') {
      writer.uint32(50).string(message.proxyAddress)
    }
    if (message.extra !== '') {
      writer.uint32(58).string(message.extra)
    }
    if (message.headerOrCrossChainMsg !== '') {
      writer.uint32(66).string(message.headerOrCrossChainMsg)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateChain {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgCreateChain } as MsgCreateChain
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
          message.sourceChainID = reader.int32()
          break
        case 4:
          message.height = reader.int32()
          break
        case 5:
          message.proof = reader.string()
          break
        case 6:
          message.proxyAddress = reader.string()
          break
        case 7:
          message.extra = reader.string()
          break
        case 8:
          message.headerOrCrossChainMsg = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgCreateChain {
    const message = { ...baseMsgCreateChain } as MsgCreateChain
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
    if (object.sourceChainID !== undefined && object.sourceChainID !== null) {
      message.sourceChainID = Number(object.sourceChainID)
    } else {
      message.sourceChainID = 0
    }
    if (object.height !== undefined && object.height !== null) {
      message.height = Number(object.height)
    } else {
      message.height = 0
    }
    if (object.proof !== undefined && object.proof !== null) {
      message.proof = String(object.proof)
    } else {
      message.proof = ''
    }
    if (object.proxyAddress !== undefined && object.proxyAddress !== null) {
      message.proxyAddress = String(object.proxyAddress)
    } else {
      message.proxyAddress = ''
    }
    if (object.extra !== undefined && object.extra !== null) {
      message.extra = String(object.extra)
    } else {
      message.extra = ''
    }
    if (object.headerOrCrossChainMsg !== undefined && object.headerOrCrossChainMsg !== null) {
      message.headerOrCrossChainMsg = String(object.headerOrCrossChainMsg)
    } else {
      message.headerOrCrossChainMsg = ''
    }
    return message
  },

  toJSON(message: MsgCreateChain): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.index !== undefined && (obj.index = message.index)
    message.sourceChainID !== undefined && (obj.sourceChainID = message.sourceChainID)
    message.height !== undefined && (obj.height = message.height)
    message.proof !== undefined && (obj.proof = message.proof)
    message.proxyAddress !== undefined && (obj.proxyAddress = message.proxyAddress)
    message.extra !== undefined && (obj.extra = message.extra)
    message.headerOrCrossChainMsg !== undefined && (obj.headerOrCrossChainMsg = message.headerOrCrossChainMsg)
    return obj
  },

  fromPartial(object: DeepPartial<MsgCreateChain>): MsgCreateChain {
    const message = { ...baseMsgCreateChain } as MsgCreateChain
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
    if (object.sourceChainID !== undefined && object.sourceChainID !== null) {
      message.sourceChainID = object.sourceChainID
    } else {
      message.sourceChainID = 0
    }
    if (object.height !== undefined && object.height !== null) {
      message.height = object.height
    } else {
      message.height = 0
    }
    if (object.proof !== undefined && object.proof !== null) {
      message.proof = object.proof
    } else {
      message.proof = ''
    }
    if (object.proxyAddress !== undefined && object.proxyAddress !== null) {
      message.proxyAddress = object.proxyAddress
    } else {
      message.proxyAddress = ''
    }
    if (object.extra !== undefined && object.extra !== null) {
      message.extra = object.extra
    } else {
      message.extra = ''
    }
    if (object.headerOrCrossChainMsg !== undefined && object.headerOrCrossChainMsg !== null) {
      message.headerOrCrossChainMsg = object.headerOrCrossChainMsg
    } else {
      message.headerOrCrossChainMsg = ''
    }
    return message
  }
}

const baseMsgCreateChainResponse: object = {}

export const MsgCreateChainResponse = {
  encode(_: MsgCreateChainResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateChainResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgCreateChainResponse } as MsgCreateChainResponse
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

  fromJSON(_: any): MsgCreateChainResponse {
    const message = { ...baseMsgCreateChainResponse } as MsgCreateChainResponse
    return message
  },

  toJSON(_: MsgCreateChainResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgCreateChainResponse>): MsgCreateChainResponse {
    const message = { ...baseMsgCreateChainResponse } as MsgCreateChainResponse
    return message
  }
}

const baseMsgUpdateChain: object = { creator: '', index: '', sourceChainID: 0, height: 0, proof: '', proxyAddress: '', extra: '', headerOrCrossChainMsg: '' }

export const MsgUpdateChain = {
  encode(message: MsgUpdateChain, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.index !== '') {
      writer.uint32(18).string(message.index)
    }
    if (message.sourceChainID !== 0) {
      writer.uint32(24).int32(message.sourceChainID)
    }
    if (message.height !== 0) {
      writer.uint32(32).int32(message.height)
    }
    if (message.proof !== '') {
      writer.uint32(42).string(message.proof)
    }
    if (message.proxyAddress !== '') {
      writer.uint32(50).string(message.proxyAddress)
    }
    if (message.extra !== '') {
      writer.uint32(58).string(message.extra)
    }
    if (message.headerOrCrossChainMsg !== '') {
      writer.uint32(66).string(message.headerOrCrossChainMsg)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateChain {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgUpdateChain } as MsgUpdateChain
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
          message.sourceChainID = reader.int32()
          break
        case 4:
          message.height = reader.int32()
          break
        case 5:
          message.proof = reader.string()
          break
        case 6:
          message.proxyAddress = reader.string()
          break
        case 7:
          message.extra = reader.string()
          break
        case 8:
          message.headerOrCrossChainMsg = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgUpdateChain {
    const message = { ...baseMsgUpdateChain } as MsgUpdateChain
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
    if (object.sourceChainID !== undefined && object.sourceChainID !== null) {
      message.sourceChainID = Number(object.sourceChainID)
    } else {
      message.sourceChainID = 0
    }
    if (object.height !== undefined && object.height !== null) {
      message.height = Number(object.height)
    } else {
      message.height = 0
    }
    if (object.proof !== undefined && object.proof !== null) {
      message.proof = String(object.proof)
    } else {
      message.proof = ''
    }
    if (object.proxyAddress !== undefined && object.proxyAddress !== null) {
      message.proxyAddress = String(object.proxyAddress)
    } else {
      message.proxyAddress = ''
    }
    if (object.extra !== undefined && object.extra !== null) {
      message.extra = String(object.extra)
    } else {
      message.extra = ''
    }
    if (object.headerOrCrossChainMsg !== undefined && object.headerOrCrossChainMsg !== null) {
      message.headerOrCrossChainMsg = String(object.headerOrCrossChainMsg)
    } else {
      message.headerOrCrossChainMsg = ''
    }
    return message
  },

  toJSON(message: MsgUpdateChain): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.index !== undefined && (obj.index = message.index)
    message.sourceChainID !== undefined && (obj.sourceChainID = message.sourceChainID)
    message.height !== undefined && (obj.height = message.height)
    message.proof !== undefined && (obj.proof = message.proof)
    message.proxyAddress !== undefined && (obj.proxyAddress = message.proxyAddress)
    message.extra !== undefined && (obj.extra = message.extra)
    message.headerOrCrossChainMsg !== undefined && (obj.headerOrCrossChainMsg = message.headerOrCrossChainMsg)
    return obj
  },

  fromPartial(object: DeepPartial<MsgUpdateChain>): MsgUpdateChain {
    const message = { ...baseMsgUpdateChain } as MsgUpdateChain
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
    if (object.sourceChainID !== undefined && object.sourceChainID !== null) {
      message.sourceChainID = object.sourceChainID
    } else {
      message.sourceChainID = 0
    }
    if (object.height !== undefined && object.height !== null) {
      message.height = object.height
    } else {
      message.height = 0
    }
    if (object.proof !== undefined && object.proof !== null) {
      message.proof = object.proof
    } else {
      message.proof = ''
    }
    if (object.proxyAddress !== undefined && object.proxyAddress !== null) {
      message.proxyAddress = object.proxyAddress
    } else {
      message.proxyAddress = ''
    }
    if (object.extra !== undefined && object.extra !== null) {
      message.extra = object.extra
    } else {
      message.extra = ''
    }
    if (object.headerOrCrossChainMsg !== undefined && object.headerOrCrossChainMsg !== null) {
      message.headerOrCrossChainMsg = object.headerOrCrossChainMsg
    } else {
      message.headerOrCrossChainMsg = ''
    }
    return message
  }
}

const baseMsgUpdateChainResponse: object = {}

export const MsgUpdateChainResponse = {
  encode(_: MsgUpdateChainResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateChainResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgUpdateChainResponse } as MsgUpdateChainResponse
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

  fromJSON(_: any): MsgUpdateChainResponse {
    const message = { ...baseMsgUpdateChainResponse } as MsgUpdateChainResponse
    return message
  },

  toJSON(_: MsgUpdateChainResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgUpdateChainResponse>): MsgUpdateChainResponse {
    const message = { ...baseMsgUpdateChainResponse } as MsgUpdateChainResponse
    return message
  }
}

const baseMsgDeleteChain: object = { creator: '', index: '' }

export const MsgDeleteChain = {
  encode(message: MsgDeleteChain, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.index !== '') {
      writer.uint32(18).string(message.index)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteChain {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgDeleteChain } as MsgDeleteChain
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

  fromJSON(object: any): MsgDeleteChain {
    const message = { ...baseMsgDeleteChain } as MsgDeleteChain
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

  toJSON(message: MsgDeleteChain): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.index !== undefined && (obj.index = message.index)
    return obj
  },

  fromPartial(object: DeepPartial<MsgDeleteChain>): MsgDeleteChain {
    const message = { ...baseMsgDeleteChain } as MsgDeleteChain
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

const baseMsgDeleteChainResponse: object = {}

export const MsgDeleteChainResponse = {
  encode(_: MsgDeleteChainResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteChainResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgDeleteChainResponse } as MsgDeleteChainResponse
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

  fromJSON(_: any): MsgDeleteChainResponse {
    const message = { ...baseMsgDeleteChainResponse } as MsgDeleteChainResponse
    return message
  },

  toJSON(_: MsgDeleteChainResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgDeleteChainResponse>): MsgDeleteChainResponse {
    const message = { ...baseMsgDeleteChainResponse } as MsgDeleteChainResponse
    return message
  }
}

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  CreateChain(request: MsgCreateChain): Promise<MsgCreateChainResponse>
  UpdateChain(request: MsgUpdateChain): Promise<MsgUpdateChainResponse>
  DeleteChain(request: MsgDeleteChain): Promise<MsgDeleteChainResponse>
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc
  constructor(rpc: Rpc) {
    this.rpc = rpc
  }
  CreateChain(request: MsgCreateChain): Promise<MsgCreateChainResponse> {
    const data = MsgCreateChain.encode(request).finish()
    const promise = this.rpc.request('pchainorg.pibridge.chain.Msg', 'CreateChain', data)
    return promise.then((data) => MsgCreateChainResponse.decode(new Reader(data)))
  }

  UpdateChain(request: MsgUpdateChain): Promise<MsgUpdateChainResponse> {
    const data = MsgUpdateChain.encode(request).finish()
    const promise = this.rpc.request('pchainorg.pibridge.chain.Msg', 'UpdateChain', data)
    return promise.then((data) => MsgUpdateChainResponse.decode(new Reader(data)))
  }

  DeleteChain(request: MsgDeleteChain): Promise<MsgDeleteChainResponse> {
    const data = MsgDeleteChain.encode(request).finish()
    const promise = this.rpc.request('pchainorg.pibridge.chain.Msg', 'DeleteChain', data)
    return promise.then((data) => MsgDeleteChainResponse.decode(new Reader(data)))
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
