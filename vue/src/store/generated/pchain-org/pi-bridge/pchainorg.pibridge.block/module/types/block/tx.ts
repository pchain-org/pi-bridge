/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal'

export const protobufPackage = 'pchainorg.pibridge.block'

/** this line is used by starport scaffolding # proto/tx/message */
export interface MsgCreateBlock {
  creator: string
  index: string
  chainID: number
  address: string
  headers: string[]
}

export interface MsgCreateBlockResponse {}

export interface MsgUpdateBlock {
  creator: string
  index: string
  chainID: number
  address: string
  headers: string[]
}

export interface MsgUpdateBlockResponse {}

export interface MsgDeleteBlock {
  creator: string
  index: string
}

export interface MsgDeleteBlockResponse {}

const baseMsgCreateBlock: object = { creator: '', index: '', chainID: 0, address: '', headers: '' }

export const MsgCreateBlock = {
  encode(message: MsgCreateBlock, writer: Writer = Writer.create()): Writer {
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
    for (const v of message.headers) {
      writer.uint32(42).string(v!)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateBlock {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgCreateBlock } as MsgCreateBlock
    message.headers = []
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
          message.headers.push(reader.string())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgCreateBlock {
    const message = { ...baseMsgCreateBlock } as MsgCreateBlock
    message.headers = []
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
      for (const e of object.headers) {
        message.headers.push(String(e))
      }
    }
    return message
  },

  toJSON(message: MsgCreateBlock): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.index !== undefined && (obj.index = message.index)
    message.chainID !== undefined && (obj.chainID = message.chainID)
    message.address !== undefined && (obj.address = message.address)
    if (message.headers) {
      obj.headers = message.headers.map((e) => e)
    } else {
      obj.headers = []
    }
    return obj
  },

  fromPartial(object: DeepPartial<MsgCreateBlock>): MsgCreateBlock {
    const message = { ...baseMsgCreateBlock } as MsgCreateBlock
    message.headers = []
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
      for (const e of object.headers) {
        message.headers.push(e)
      }
    }
    return message
  }
}

const baseMsgCreateBlockResponse: object = {}

export const MsgCreateBlockResponse = {
  encode(_: MsgCreateBlockResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateBlockResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgCreateBlockResponse } as MsgCreateBlockResponse
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

  fromJSON(_: any): MsgCreateBlockResponse {
    const message = { ...baseMsgCreateBlockResponse } as MsgCreateBlockResponse
    return message
  },

  toJSON(_: MsgCreateBlockResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgCreateBlockResponse>): MsgCreateBlockResponse {
    const message = { ...baseMsgCreateBlockResponse } as MsgCreateBlockResponse
    return message
  }
}

const baseMsgUpdateBlock: object = { creator: '', index: '', chainID: 0, address: '', headers: '' }

export const MsgUpdateBlock = {
  encode(message: MsgUpdateBlock, writer: Writer = Writer.create()): Writer {
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
    for (const v of message.headers) {
      writer.uint32(42).string(v!)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateBlock {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgUpdateBlock } as MsgUpdateBlock
    message.headers = []
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
          message.headers.push(reader.string())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgUpdateBlock {
    const message = { ...baseMsgUpdateBlock } as MsgUpdateBlock
    message.headers = []
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
      for (const e of object.headers) {
        message.headers.push(String(e))
      }
    }
    return message
  },

  toJSON(message: MsgUpdateBlock): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.index !== undefined && (obj.index = message.index)
    message.chainID !== undefined && (obj.chainID = message.chainID)
    message.address !== undefined && (obj.address = message.address)
    if (message.headers) {
      obj.headers = message.headers.map((e) => e)
    } else {
      obj.headers = []
    }
    return obj
  },

  fromPartial(object: DeepPartial<MsgUpdateBlock>): MsgUpdateBlock {
    const message = { ...baseMsgUpdateBlock } as MsgUpdateBlock
    message.headers = []
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
      for (const e of object.headers) {
        message.headers.push(e)
      }
    }
    return message
  }
}

const baseMsgUpdateBlockResponse: object = {}

export const MsgUpdateBlockResponse = {
  encode(_: MsgUpdateBlockResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateBlockResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgUpdateBlockResponse } as MsgUpdateBlockResponse
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

  fromJSON(_: any): MsgUpdateBlockResponse {
    const message = { ...baseMsgUpdateBlockResponse } as MsgUpdateBlockResponse
    return message
  },

  toJSON(_: MsgUpdateBlockResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgUpdateBlockResponse>): MsgUpdateBlockResponse {
    const message = { ...baseMsgUpdateBlockResponse } as MsgUpdateBlockResponse
    return message
  }
}

const baseMsgDeleteBlock: object = { creator: '', index: '' }

export const MsgDeleteBlock = {
  encode(message: MsgDeleteBlock, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.index !== '') {
      writer.uint32(18).string(message.index)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteBlock {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgDeleteBlock } as MsgDeleteBlock
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

  fromJSON(object: any): MsgDeleteBlock {
    const message = { ...baseMsgDeleteBlock } as MsgDeleteBlock
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

  toJSON(message: MsgDeleteBlock): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.index !== undefined && (obj.index = message.index)
    return obj
  },

  fromPartial(object: DeepPartial<MsgDeleteBlock>): MsgDeleteBlock {
    const message = { ...baseMsgDeleteBlock } as MsgDeleteBlock
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

const baseMsgDeleteBlockResponse: object = {}

export const MsgDeleteBlockResponse = {
  encode(_: MsgDeleteBlockResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteBlockResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgDeleteBlockResponse } as MsgDeleteBlockResponse
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

  fromJSON(_: any): MsgDeleteBlockResponse {
    const message = { ...baseMsgDeleteBlockResponse } as MsgDeleteBlockResponse
    return message
  },

  toJSON(_: MsgDeleteBlockResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgDeleteBlockResponse>): MsgDeleteBlockResponse {
    const message = { ...baseMsgDeleteBlockResponse } as MsgDeleteBlockResponse
    return message
  }
}

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  CreateBlock(request: MsgCreateBlock): Promise<MsgCreateBlockResponse>
  UpdateBlock(request: MsgUpdateBlock): Promise<MsgUpdateBlockResponse>
  DeleteBlock(request: MsgDeleteBlock): Promise<MsgDeleteBlockResponse>
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc
  constructor(rpc: Rpc) {
    this.rpc = rpc
  }
  CreateBlock(request: MsgCreateBlock): Promise<MsgCreateBlockResponse> {
    const data = MsgCreateBlock.encode(request).finish()
    const promise = this.rpc.request('pchainorg.pibridge.block.Msg', 'CreateBlock', data)
    return promise.then((data) => MsgCreateBlockResponse.decode(new Reader(data)))
  }

  UpdateBlock(request: MsgUpdateBlock): Promise<MsgUpdateBlockResponse> {
    const data = MsgUpdateBlock.encode(request).finish()
    const promise = this.rpc.request('pchainorg.pibridge.block.Msg', 'UpdateBlock', data)
    return promise.then((data) => MsgUpdateBlockResponse.decode(new Reader(data)))
  }

  DeleteBlock(request: MsgDeleteBlock): Promise<MsgDeleteBlockResponse> {
    const data = MsgDeleteBlock.encode(request).finish()
    const promise = this.rpc.request('pchainorg.pibridge.block.Msg', 'DeleteBlock', data)
    return promise.then((data) => MsgDeleteBlockResponse.decode(new Reader(data)))
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
