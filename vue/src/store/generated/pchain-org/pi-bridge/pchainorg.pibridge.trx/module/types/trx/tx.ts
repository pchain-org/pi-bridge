/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal'

export const protobufPackage = 'pchainorg.pibridge.trx'

/** this line is used by starport scaffolding # proto/tx/message */
export interface MsgCrossChainTransfer {
  creator: string
  asset: string
  amount: number
  fromChain: string
  fromAddress: string
  toChain: string
  toAddress: string
}

export interface MsgCrossChainTransferResponse {}

export interface MsgCreateTrx {
  creator: string
  index: string
  chainID: number
  address: string
  crossChainMsgs: string
  proof: string
}

export interface MsgCreateTrxResponse {}

export interface MsgUpdateTrx {
  creator: string
  index: string
  chainID: number
  address: string
  crossChainMsgs: string
}

export interface MsgUpdateTrxResponse {}

export interface MsgDeleteTrx {
  creator: string
  index: string
}

export interface MsgDeleteTrxResponse {}

const baseMsgCrossChainTransfer: object = { creator: '', asset: '', amount: 0, fromChain: '', fromAddress: '', toChain: '', toAddress: '' }

export const MsgCrossChainTransfer = {
  encode(message: MsgCrossChainTransfer, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.asset !== '') {
      writer.uint32(18).string(message.asset)
    }
    if (message.amount !== 0) {
      writer.uint32(24).int32(message.amount)
    }
    if (message.fromChain !== '') {
      writer.uint32(34).string(message.fromChain)
    }
    if (message.fromAddress !== '') {
      writer.uint32(42).string(message.fromAddress)
    }
    if (message.toChain !== '') {
      writer.uint32(50).string(message.toChain)
    }
    if (message.toAddress !== '') {
      writer.uint32(58).string(message.toAddress)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCrossChainTransfer {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgCrossChainTransfer } as MsgCrossChainTransfer
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string()
          break
        case 2:
          message.asset = reader.string()
          break
        case 3:
          message.amount = reader.int32()
          break
        case 4:
          message.fromChain = reader.string()
          break
        case 5:
          message.fromAddress = reader.string()
          break
        case 6:
          message.toChain = reader.string()
          break
        case 7:
          message.toAddress = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgCrossChainTransfer {
    const message = { ...baseMsgCrossChainTransfer } as MsgCrossChainTransfer
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
    }
    if (object.asset !== undefined && object.asset !== null) {
      message.asset = String(object.asset)
    } else {
      message.asset = ''
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = Number(object.amount)
    } else {
      message.amount = 0
    }
    if (object.fromChain !== undefined && object.fromChain !== null) {
      message.fromChain = String(object.fromChain)
    } else {
      message.fromChain = ''
    }
    if (object.fromAddress !== undefined && object.fromAddress !== null) {
      message.fromAddress = String(object.fromAddress)
    } else {
      message.fromAddress = ''
    }
    if (object.toChain !== undefined && object.toChain !== null) {
      message.toChain = String(object.toChain)
    } else {
      message.toChain = ''
    }
    if (object.toAddress !== undefined && object.toAddress !== null) {
      message.toAddress = String(object.toAddress)
    } else {
      message.toAddress = ''
    }
    return message
  },

  toJSON(message: MsgCrossChainTransfer): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.asset !== undefined && (obj.asset = message.asset)
    message.amount !== undefined && (obj.amount = message.amount)
    message.fromChain !== undefined && (obj.fromChain = message.fromChain)
    message.fromAddress !== undefined && (obj.fromAddress = message.fromAddress)
    message.toChain !== undefined && (obj.toChain = message.toChain)
    message.toAddress !== undefined && (obj.toAddress = message.toAddress)
    return obj
  },

  fromPartial(object: DeepPartial<MsgCrossChainTransfer>): MsgCrossChainTransfer {
    const message = { ...baseMsgCrossChainTransfer } as MsgCrossChainTransfer
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator
    } else {
      message.creator = ''
    }
    if (object.asset !== undefined && object.asset !== null) {
      message.asset = object.asset
    } else {
      message.asset = ''
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = object.amount
    } else {
      message.amount = 0
    }
    if (object.fromChain !== undefined && object.fromChain !== null) {
      message.fromChain = object.fromChain
    } else {
      message.fromChain = ''
    }
    if (object.fromAddress !== undefined && object.fromAddress !== null) {
      message.fromAddress = object.fromAddress
    } else {
      message.fromAddress = ''
    }
    if (object.toChain !== undefined && object.toChain !== null) {
      message.toChain = object.toChain
    } else {
      message.toChain = ''
    }
    if (object.toAddress !== undefined && object.toAddress !== null) {
      message.toAddress = object.toAddress
    } else {
      message.toAddress = ''
    }
    return message
  }
}

const baseMsgCrossChainTransferResponse: object = {}

export const MsgCrossChainTransferResponse = {
  encode(_: MsgCrossChainTransferResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCrossChainTransferResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgCrossChainTransferResponse } as MsgCrossChainTransferResponse
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

  fromJSON(_: any): MsgCrossChainTransferResponse {
    const message = { ...baseMsgCrossChainTransferResponse } as MsgCrossChainTransferResponse
    return message
  },

  toJSON(_: MsgCrossChainTransferResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgCrossChainTransferResponse>): MsgCrossChainTransferResponse {
    const message = { ...baseMsgCrossChainTransferResponse } as MsgCrossChainTransferResponse
    return message
  }
}

const baseMsgCreateTrx: object = { creator: '', index: '', chainID: 0, address: '', crossChainMsgs: '', proof: '' }

export const MsgCreateTrx = {
  encode(message: MsgCreateTrx, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.index !== '') {
      writer.uint32(18).string(message.index)
    }
    if (message.chainID !== 0) {
      writer.uint32(24).int32(message.chainID)
    }
    if (message.address !== '') {
      writer.uint32(34).string(message.address)
    }
    if (message.crossChainMsgs !== '') {
      writer.uint32(42).string(message.crossChainMsgs)
    }
    if (message.proof !== '') {
      writer.uint32(50).string(message.proof)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateTrx {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgCreateTrx } as MsgCreateTrx
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
          message.chainID = reader.int32()
          break
        case 4:
          message.address = reader.string()
          break
        case 5:
          message.crossChainMsgs = reader.string()
          break
        case 6:
          message.proof = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgCreateTrx {
    const message = { ...baseMsgCreateTrx } as MsgCreateTrx
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
    if (object.chainID !== undefined && object.chainID !== null) {
      message.chainID = Number(object.chainID)
    } else {
      message.chainID = 0
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address)
    } else {
      message.address = ''
    }
    if (object.crossChainMsgs !== undefined && object.crossChainMsgs !== null) {
      message.crossChainMsgs = String(object.crossChainMsgs)
    } else {
      message.crossChainMsgs = ''
    }
    if (object.proof !== undefined && object.proof !== null) {
      message.proof = String(object.proof)
    } else {
      message.proof = ''
    }
    return message
  },

  toJSON(message: MsgCreateTrx): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.index !== undefined && (obj.index = message.index)
    message.chainID !== undefined && (obj.chainID = message.chainID)
    message.address !== undefined && (obj.address = message.address)
    message.crossChainMsgs !== undefined && (obj.crossChainMsgs = message.crossChainMsgs)
    message.proof !== undefined && (obj.proof = message.proof)
    return obj
  },

  fromPartial(object: DeepPartial<MsgCreateTrx>): MsgCreateTrx {
    const message = { ...baseMsgCreateTrx } as MsgCreateTrx
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
    if (object.chainID !== undefined && object.chainID !== null) {
      message.chainID = object.chainID
    } else {
      message.chainID = 0
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address
    } else {
      message.address = ''
    }
    if (object.crossChainMsgs !== undefined && object.crossChainMsgs !== null) {
      message.crossChainMsgs = object.crossChainMsgs
    } else {
      message.crossChainMsgs = ''
    }
    if (object.proof !== undefined && object.proof !== null) {
      message.proof = object.proof
    } else {
      message.proof = ''
    }
    return message
  }
}

const baseMsgCreateTrxResponse: object = {}

export const MsgCreateTrxResponse = {
  encode(_: MsgCreateTrxResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateTrxResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgCreateTrxResponse } as MsgCreateTrxResponse
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

  fromJSON(_: any): MsgCreateTrxResponse {
    const message = { ...baseMsgCreateTrxResponse } as MsgCreateTrxResponse
    return message
  },

  toJSON(_: MsgCreateTrxResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgCreateTrxResponse>): MsgCreateTrxResponse {
    const message = { ...baseMsgCreateTrxResponse } as MsgCreateTrxResponse
    return message
  }
}

const baseMsgUpdateTrx: object = { creator: '', index: '', chainID: 0, address: '', crossChainMsgs: '' }

export const MsgUpdateTrx = {
  encode(message: MsgUpdateTrx, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.index !== '') {
      writer.uint32(18).string(message.index)
    }
    if (message.chainID !== 0) {
      writer.uint32(24).int32(message.chainID)
    }
    if (message.address !== '') {
      writer.uint32(34).string(message.address)
    }
    if (message.crossChainMsgs !== '') {
      writer.uint32(42).string(message.crossChainMsgs)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateTrx {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgUpdateTrx } as MsgUpdateTrx
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
          message.chainID = reader.int32()
          break
        case 4:
          message.address = reader.string()
          break
        case 5:
          message.crossChainMsgs = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgUpdateTrx {
    const message = { ...baseMsgUpdateTrx } as MsgUpdateTrx
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
    if (object.chainID !== undefined && object.chainID !== null) {
      message.chainID = Number(object.chainID)
    } else {
      message.chainID = 0
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address)
    } else {
      message.address = ''
    }
    if (object.crossChainMsgs !== undefined && object.crossChainMsgs !== null) {
      message.crossChainMsgs = String(object.crossChainMsgs)
    } else {
      message.crossChainMsgs = ''
    }
    return message
  },

  toJSON(message: MsgUpdateTrx): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.index !== undefined && (obj.index = message.index)
    message.chainID !== undefined && (obj.chainID = message.chainID)
    message.address !== undefined && (obj.address = message.address)
    message.crossChainMsgs !== undefined && (obj.crossChainMsgs = message.crossChainMsgs)
    return obj
  },

  fromPartial(object: DeepPartial<MsgUpdateTrx>): MsgUpdateTrx {
    const message = { ...baseMsgUpdateTrx } as MsgUpdateTrx
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
    if (object.chainID !== undefined && object.chainID !== null) {
      message.chainID = object.chainID
    } else {
      message.chainID = 0
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address
    } else {
      message.address = ''
    }
    if (object.crossChainMsgs !== undefined && object.crossChainMsgs !== null) {
      message.crossChainMsgs = object.crossChainMsgs
    } else {
      message.crossChainMsgs = ''
    }
    return message
  }
}

const baseMsgUpdateTrxResponse: object = {}

export const MsgUpdateTrxResponse = {
  encode(_: MsgUpdateTrxResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateTrxResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgUpdateTrxResponse } as MsgUpdateTrxResponse
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

  fromJSON(_: any): MsgUpdateTrxResponse {
    const message = { ...baseMsgUpdateTrxResponse } as MsgUpdateTrxResponse
    return message
  },

  toJSON(_: MsgUpdateTrxResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgUpdateTrxResponse>): MsgUpdateTrxResponse {
    const message = { ...baseMsgUpdateTrxResponse } as MsgUpdateTrxResponse
    return message
  }
}

const baseMsgDeleteTrx: object = { creator: '', index: '' }

export const MsgDeleteTrx = {
  encode(message: MsgDeleteTrx, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.index !== '') {
      writer.uint32(18).string(message.index)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteTrx {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgDeleteTrx } as MsgDeleteTrx
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

  fromJSON(object: any): MsgDeleteTrx {
    const message = { ...baseMsgDeleteTrx } as MsgDeleteTrx
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

  toJSON(message: MsgDeleteTrx): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.index !== undefined && (obj.index = message.index)
    return obj
  },

  fromPartial(object: DeepPartial<MsgDeleteTrx>): MsgDeleteTrx {
    const message = { ...baseMsgDeleteTrx } as MsgDeleteTrx
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

const baseMsgDeleteTrxResponse: object = {}

export const MsgDeleteTrxResponse = {
  encode(_: MsgDeleteTrxResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteTrxResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgDeleteTrxResponse } as MsgDeleteTrxResponse
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

  fromJSON(_: any): MsgDeleteTrxResponse {
    const message = { ...baseMsgDeleteTrxResponse } as MsgDeleteTrxResponse
    return message
  },

  toJSON(_: MsgDeleteTrxResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgDeleteTrxResponse>): MsgDeleteTrxResponse {
    const message = { ...baseMsgDeleteTrxResponse } as MsgDeleteTrxResponse
    return message
  }
}

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  CrossChainTransfer(request: MsgCrossChainTransfer): Promise<MsgCrossChainTransferResponse>
  CreateTrx(request: MsgCreateTrx): Promise<MsgCreateTrxResponse>
  UpdateTrx(request: MsgUpdateTrx): Promise<MsgUpdateTrxResponse>
  DeleteTrx(request: MsgDeleteTrx): Promise<MsgDeleteTrxResponse>
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc
  constructor(rpc: Rpc) {
    this.rpc = rpc
  }
  CrossChainTransfer(request: MsgCrossChainTransfer): Promise<MsgCrossChainTransferResponse> {
    const data = MsgCrossChainTransfer.encode(request).finish()
    const promise = this.rpc.request('pchainorg.pibridge.trx.Msg', 'CrossChainTransfer', data)
    return promise.then((data) => MsgCrossChainTransferResponse.decode(new Reader(data)))
  }

  CreateTrx(request: MsgCreateTrx): Promise<MsgCreateTrxResponse> {
    const data = MsgCreateTrx.encode(request).finish()
    const promise = this.rpc.request('pchainorg.pibridge.trx.Msg', 'CreateTrx', data)
    return promise.then((data) => MsgCreateTrxResponse.decode(new Reader(data)))
  }

  UpdateTrx(request: MsgUpdateTrx): Promise<MsgUpdateTrxResponse> {
    const data = MsgUpdateTrx.encode(request).finish()
    const promise = this.rpc.request('pchainorg.pibridge.trx.Msg', 'UpdateTrx', data)
    return promise.then((data) => MsgUpdateTrxResponse.decode(new Reader(data)))
  }

  DeleteTrx(request: MsgDeleteTrx): Promise<MsgDeleteTrxResponse> {
    const data = MsgDeleteTrx.encode(request).finish()
    const promise = this.rpc.request('pchainorg.pibridge.trx.Msg', 'DeleteTrx', data)
    return promise.then((data) => MsgDeleteTrxResponse.decode(new Reader(data)))
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
