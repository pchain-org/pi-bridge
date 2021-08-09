/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal';
import { Node } from '../node/node';
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination';
export const protobufPackage = 'pchainorg.pibridge.node';
const baseQueryGetNodeRequest = { index: '' };
export const QueryGetNodeRequest = {
    encode(message, writer = Writer.create()) {
        if (message.index !== '') {
            writer.uint32(10).string(message.index);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetNodeRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.index = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetNodeRequest };
        if (object.index !== undefined && object.index !== null) {
            message.index = String(object.index);
        }
        else {
            message.index = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.index !== undefined && (obj.index = message.index);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetNodeRequest };
        if (object.index !== undefined && object.index !== null) {
            message.index = object.index;
        }
        else {
            message.index = '';
        }
        return message;
    }
};
const baseQueryGetNodeResponse = {};
export const QueryGetNodeResponse = {
    encode(message, writer = Writer.create()) {
        if (message.Node !== undefined) {
            Node.encode(message.Node, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetNodeResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.Node = Node.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetNodeResponse };
        if (object.Node !== undefined && object.Node !== null) {
            message.Node = Node.fromJSON(object.Node);
        }
        else {
            message.Node = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.Node !== undefined && (obj.Node = message.Node ? Node.toJSON(message.Node) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetNodeResponse };
        if (object.Node !== undefined && object.Node !== null) {
            message.Node = Node.fromPartial(object.Node);
        }
        else {
            message.Node = undefined;
        }
        return message;
    }
};
const baseQueryAllNodeRequest = {};
export const QueryAllNodeRequest = {
    encode(message, writer = Writer.create()) {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllNodeRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.pagination = PageRequest.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryAllNodeRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllNodeRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryAllNodeResponse = {};
export const QueryAllNodeResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.Node) {
            Node.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllNodeResponse };
        message.Node = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.Node.push(Node.decode(reader, reader.uint32()));
                    break;
                case 2:
                    message.pagination = PageResponse.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryAllNodeResponse };
        message.Node = [];
        if (object.Node !== undefined && object.Node !== null) {
            for (const e of object.Node) {
                message.Node.push(Node.fromJSON(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.Node) {
            obj.Node = message.Node.map((e) => (e ? Node.toJSON(e) : undefined));
        }
        else {
            obj.Node = [];
        }
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllNodeResponse };
        message.Node = [];
        if (object.Node !== undefined && object.Node !== null) {
            for (const e of object.Node) {
                message.Node.push(Node.fromPartial(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
export class QueryClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    Node(request) {
        const data = QueryGetNodeRequest.encode(request).finish();
        const promise = this.rpc.request('pchainorg.pibridge.node.Query', 'Node', data);
        return promise.then((data) => QueryGetNodeResponse.decode(new Reader(data)));
    }
    NodeAll(request) {
        const data = QueryAllNodeRequest.encode(request).finish();
        const promise = this.rpc.request('pchainorg.pibridge.node.Query', 'NodeAll', data);
        return promise.then((data) => QueryAllNodeResponse.decode(new Reader(data)));
    }
}
