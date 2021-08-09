/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal'

export const protobufPackage = 'pchainorg.pibridge.node'

/** this line is used by starport scaffolding # proto/tx/message */
export interface MsgCreateNode {
  creator: string
  index: string
  peerPubkey: string
  address: string
}

export interface MsgCreateNodeResponse {}

export interface MsgUpdateNode {
  creator: string
  index: string
  peerPubkey: string
  address: string
}

export interface MsgUpdateNodeResponse {}

export interface MsgDeleteNode {
  creator: string
  index: string
}

export interface MsgDeleteNodeResponse {}

const baseMsgCreateNode: object = { creator: '', index: '', peerPubkey: '', address: '' }

export const MsgCreateNode = {
  encode(message: MsgCreateNode, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.index !== '') {
      writer.uint32(18).string(message.index)
    }
    if (message.peerPubkey !== '') {
      writer.uint32(26).string(message.peerPubkey)
    }
    if (message.address !== '') {
      writer.uint32(34).string(message.address)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateNode {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgCreateNode } as MsgCreateNode
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
          message.peerPubkey = reader.string()
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

  fromJSON(object: any): MsgCreateNode {
    const message = { ...baseMsgCreateNode } as MsgCreateNode
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
    if (object.peerPubkey !== undefined && object.peerPubkey !== null) {
      message.peerPubkey = String(object.peerPubkey)
    } else {
      message.peerPubkey = ''
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address)
    } else {
      message.address = ''
    }
    return message
  },

  toJSON(message: MsgCreateNode): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.index !== undefined && (obj.index = message.index)
    message.peerPubkey !== undefined && (obj.peerPubkey = message.peerPubkey)
    message.address !== undefined && (obj.address = message.address)
    return obj
  },

  fromPartial(object: DeepPartial<MsgCreateNode>): MsgCreateNode {
    const message = { ...baseMsgCreateNode } as MsgCreateNode
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
    if (object.peerPubkey !== undefined && object.peerPubkey !== null) {
      message.peerPubkey = object.peerPubkey
    } else {
      message.peerPubkey = ''
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address
    } else {
      message.address = ''
    }
    return message
  }
}

const baseMsgCreateNodeResponse: object = {}

export const MsgCreateNodeResponse = {
  encode(_: MsgCreateNodeResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateNodeResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgCreateNodeResponse } as MsgCreateNodeResponse
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

  fromJSON(_: any): MsgCreateNodeResponse {
    const message = { ...baseMsgCreateNodeResponse } as MsgCreateNodeResponse
    return message
  },

  toJSON(_: MsgCreateNodeResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgCreateNodeResponse>): MsgCreateNodeResponse {
    const message = { ...baseMsgCreateNodeResponse } as MsgCreateNodeResponse
    return message
  }
}

const baseMsgUpdateNode: object = { creator: '', index: '', peerPubkey: '', address: '' }

export const MsgUpdateNode = {
  encode(message: MsgUpdateNode, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.index !== '') {
      writer.uint32(18).string(message.index)
    }
    if (message.peerPubkey !== '') {
      writer.uint32(26).string(message.peerPubkey)
    }
    if (message.address !== '') {
      writer.uint32(34).string(message.address)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateNode {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgUpdateNode } as MsgUpdateNode
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
          message.peerPubkey = reader.string()
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

  fromJSON(object: any): MsgUpdateNode {
    const message = { ...baseMsgUpdateNode } as MsgUpdateNode
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
    if (object.peerPubkey !== undefined && object.peerPubkey !== null) {
      message.peerPubkey = String(object.peerPubkey)
    } else {
      message.peerPubkey = ''
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address)
    } else {
      message.address = ''
    }
    return message
  },

  toJSON(message: MsgUpdateNode): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.index !== undefined && (obj.index = message.index)
    message.peerPubkey !== undefined && (obj.peerPubkey = message.peerPubkey)
    message.address !== undefined && (obj.address = message.address)
    return obj
  },

  fromPartial(object: DeepPartial<MsgUpdateNode>): MsgUpdateNode {
    const message = { ...baseMsgUpdateNode } as MsgUpdateNode
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
    if (object.peerPubkey !== undefined && object.peerPubkey !== null) {
      message.peerPubkey = object.peerPubkey
    } else {
      message.peerPubkey = ''
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address
    } else {
      message.address = ''
    }
    return message
  }
}

const baseMsgUpdateNodeResponse: object = {}

export const MsgUpdateNodeResponse = {
  encode(_: MsgUpdateNodeResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateNodeResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgUpdateNodeResponse } as MsgUpdateNodeResponse
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

  fromJSON(_: any): MsgUpdateNodeResponse {
    const message = { ...baseMsgUpdateNodeResponse } as MsgUpdateNodeResponse
    return message
  },

  toJSON(_: MsgUpdateNodeResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgUpdateNodeResponse>): MsgUpdateNodeResponse {
    const message = { ...baseMsgUpdateNodeResponse } as MsgUpdateNodeResponse
    return message
  }
}

const baseMsgDeleteNode: object = { creator: '', index: '' }

export const MsgDeleteNode = {
  encode(message: MsgDeleteNode, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.index !== '') {
      writer.uint32(18).string(message.index)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteNode {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgDeleteNode } as MsgDeleteNode
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

  fromJSON(object: any): MsgDeleteNode {
    const message = { ...baseMsgDeleteNode } as MsgDeleteNode
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

  toJSON(message: MsgDeleteNode): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.index !== undefined && (obj.index = message.index)
    return obj
  },

  fromPartial(object: DeepPartial<MsgDeleteNode>): MsgDeleteNode {
    const message = { ...baseMsgDeleteNode } as MsgDeleteNode
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

const baseMsgDeleteNodeResponse: object = {}

export const MsgDeleteNodeResponse = {
  encode(_: MsgDeleteNodeResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteNodeResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgDeleteNodeResponse } as MsgDeleteNodeResponse
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

  fromJSON(_: any): MsgDeleteNodeResponse {
    const message = { ...baseMsgDeleteNodeResponse } as MsgDeleteNodeResponse
    return message
  },

  toJSON(_: MsgDeleteNodeResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgDeleteNodeResponse>): MsgDeleteNodeResponse {
    const message = { ...baseMsgDeleteNodeResponse } as MsgDeleteNodeResponse
    return message
  }
}

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  CreateNode(request: MsgCreateNode): Promise<MsgCreateNodeResponse>
  UpdateNode(request: MsgUpdateNode): Promise<MsgUpdateNodeResponse>
  DeleteNode(request: MsgDeleteNode): Promise<MsgDeleteNodeResponse>
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc
  constructor(rpc: Rpc) {
    this.rpc = rpc
  }
  CreateNode(request: MsgCreateNode): Promise<MsgCreateNodeResponse> {
    const data = MsgCreateNode.encode(request).finish()
    const promise = this.rpc.request('pchainorg.pibridge.node.Msg', 'CreateNode', data)
    return promise.then((data) => MsgCreateNodeResponse.decode(new Reader(data)))
  }

  UpdateNode(request: MsgUpdateNode): Promise<MsgUpdateNodeResponse> {
    const data = MsgUpdateNode.encode(request).finish()
    const promise = this.rpc.request('pchainorg.pibridge.node.Msg', 'UpdateNode', data)
    return promise.then((data) => MsgUpdateNodeResponse.decode(new Reader(data)))
  }

  DeleteNode(request: MsgDeleteNode): Promise<MsgDeleteNodeResponse> {
    const data = MsgDeleteNode.encode(request).finish()
    const promise = this.rpc.request('pchainorg.pibridge.node.Msg', 'DeleteNode', data)
    return promise.then((data) => MsgDeleteNodeResponse.decode(new Reader(data)))
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
