/* eslint-disable */
import { Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'pchainorg.pibridge.trx'

export interface Trx {
  creator: string
  index: string
  chainID: number
  address: string
  crossChainMsgs: string
}

const baseTrx: object = { creator: '', index: '', chainID: 0, address: '', crossChainMsgs: '' }

export const Trx = {
  encode(message: Trx, writer: Writer = Writer.create()): Writer {
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

  decode(input: Reader | Uint8Array, length?: number): Trx {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseTrx } as Trx
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

  fromJSON(object: any): Trx {
    const message = { ...baseTrx } as Trx
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

  toJSON(message: Trx): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.index !== undefined && (obj.index = message.index)
    message.chainID !== undefined && (obj.chainID = message.chainID)
    message.address !== undefined && (obj.address = message.address)
    message.crossChainMsgs !== undefined && (obj.crossChainMsgs = message.crossChainMsgs)
    return obj
  },

  fromPartial(object: DeepPartial<Trx>): Trx {
    const message = { ...baseTrx } as Trx
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
