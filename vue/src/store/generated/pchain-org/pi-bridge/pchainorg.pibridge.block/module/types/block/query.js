/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal';
import { Block } from '../block/block';
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination';
export const protobufPackage = 'pchainorg.pibridge.block';
const baseQueryGetBlockRequest = { index: '' };
export const QueryGetBlockRequest = {
    encode(message, writer = Writer.create()) {
        if (message.index !== '') {
            writer.uint32(10).string(message.index);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetBlockRequest };
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
        const message = { ...baseQueryGetBlockRequest };
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
        const message = { ...baseQueryGetBlockRequest };
        if (object.index !== undefined && object.index !== null) {
            message.index = object.index;
        }
        else {
            message.index = '';
        }
        return message;
    }
};
const baseQueryGetBlockResponse = {};
export const QueryGetBlockResponse = {
    encode(message, writer = Writer.create()) {
        if (message.Block !== undefined) {
            Block.encode(message.Block, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetBlockResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.Block = Block.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetBlockResponse };
        if (object.Block !== undefined && object.Block !== null) {
            message.Block = Block.fromJSON(object.Block);
        }
        else {
            message.Block = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.Block !== undefined && (obj.Block = message.Block ? Block.toJSON(message.Block) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetBlockResponse };
        if (object.Block !== undefined && object.Block !== null) {
            message.Block = Block.fromPartial(object.Block);
        }
        else {
            message.Block = undefined;
        }
        return message;
    }
};
const baseQueryAllBlockRequest = {};
export const QueryAllBlockRequest = {
    encode(message, writer = Writer.create()) {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllBlockRequest };
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
        const message = { ...baseQueryAllBlockRequest };
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
        const message = { ...baseQueryAllBlockRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryAllBlockResponse = {};
export const QueryAllBlockResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.Block) {
            Block.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllBlockResponse };
        message.Block = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.Block.push(Block.decode(reader, reader.uint32()));
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
        const message = { ...baseQueryAllBlockResponse };
        message.Block = [];
        if (object.Block !== undefined && object.Block !== null) {
            for (const e of object.Block) {
                message.Block.push(Block.fromJSON(e));
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
        if (message.Block) {
            obj.Block = message.Block.map((e) => (e ? Block.toJSON(e) : undefined));
        }
        else {
            obj.Block = [];
        }
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllBlockResponse };
        message.Block = [];
        if (object.Block !== undefined && object.Block !== null) {
            for (const e of object.Block) {
                message.Block.push(Block.fromPartial(e));
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
    Block(request) {
        const data = QueryGetBlockRequest.encode(request).finish();
        const promise = this.rpc.request('pchainorg.pibridge.block.Query', 'Block', data);
        return promise.then((data) => QueryGetBlockResponse.decode(new Reader(data)));
    }
    BlockAll(request) {
        const data = QueryAllBlockRequest.encode(request).finish();
        const promise = this.rpc.request('pchainorg.pibridge.block.Query', 'BlockAll', data);
        return promise.then((data) => QueryAllBlockResponse.decode(new Reader(data)));
    }
}
