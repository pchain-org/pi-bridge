/* eslint-disable */
import { Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'pchainorg.pibridge.chain'

export interface Chain {
  creator: string
  index: string
  sourceChainID: number
  height: number
  proof: string
  proxyAddress: string
  extra: string
  headerOrCrossChainMsg: string
}

const baseChain: object = { creator: '', index: '', sourceChainID: 0, height: 0, proof: '', proxyAddress: '', extra: '', headerOrCrossChainMsg: '' }

export const Chain = {
  encode(message: Chain, writer: Writer = Writer.create()): Writer {
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

  decode(input: Reader | Uint8Array, length?: number): Chain {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseChain } as Chain
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

  fromJSON(object: any): Chain {
    const message = { ...baseChain } as Chain
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

  toJSON(message: Chain): unknown {
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

  fromPartial(object: DeepPartial<Chain>): Chain {
    const message = { ...baseChain } as Chain
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
