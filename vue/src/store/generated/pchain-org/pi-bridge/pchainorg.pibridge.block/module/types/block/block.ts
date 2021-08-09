/* eslint-disable */
import { Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'pchainorg.pibridge.block'

export interface Block {
  creator: string
  index: string
  chainID: number
  address: string
  headers: string
}

const baseBlock: object = { creator: '', index: '', chainID: 0, address: '', headers: '' }

export const Block = {
  encode(message: Block, writer: Writer = Writer.create()): Writer {
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
    if (message.headers !== '') {
      writer.uint32(42).string(message.headers)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): Block {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseBlock } as Block
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
          message.headers = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): Block {
    const message = { ...baseBlock } as Block
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
    if (object.headers !== undefined && object.headers !== null) {
      message.headers = String(object.headers)
    } else {
      message.headers = ''
    }
    return message
  },

  toJSON(message: Block): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.index !== undefined && (obj.index = message.index)
    message.chainID !== undefined && (obj.chainID = message.chainID)
    message.address !== undefined && (obj.address = message.address)
    message.headers !== undefined && (obj.headers = message.headers)
    return obj
  },

  fromPartial(object: DeepPartial<Block>): Block {
    const message = { ...baseBlock } as Block
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
    if (object.headers !== undefined && object.headers !== null) {
      message.headers = object.headers
    } else {
      message.headers = ''
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
