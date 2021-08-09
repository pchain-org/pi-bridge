/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal';
import { Chain } from '../chain/chain';
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination';
export const protobufPackage = 'pchainorg.pibridge.chain';
const baseQueryGetChainRequest = { index: '' };
export const QueryGetChainRequest = {
    encode(message, writer = Writer.create()) {
        if (message.index !== '') {
            writer.uint32(10).string(message.index);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetChainRequest };
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
        const message = { ...baseQueryGetChainRequest };
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
        const message = { ...baseQueryGetChainRequest };
        if (object.index !== undefined && object.index !== null) {
            message.index = object.index;
        }
        else {
            message.index = '';
        }
        return message;
    }
};
const baseQueryGetChainResponse = {};
export const QueryGetChainResponse = {
    encode(message, writer = Writer.create()) {
        if (message.Chain !== undefined) {
            Chain.encode(message.Chain, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetChainResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.Chain = Chain.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetChainResponse };
        if (object.Chain !== undefined && object.Chain !== null) {
            message.Chain = Chain.fromJSON(object.Chain);
        }
        else {
            message.Chain = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.Chain !== undefined && (obj.Chain = message.Chain ? Chain.toJSON(message.Chain) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetChainResponse };
        if (object.Chain !== undefined && object.Chain !== null) {
            message.Chain = Chain.fromPartial(object.Chain);
        }
        else {
            message.Chain = undefined;
        }
        return message;
    }
};
const baseQueryAllChainRequest = {};
export const QueryAllChainRequest = {
    encode(message, writer = Writer.create()) {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllChainRequest };
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
        const message = { ...baseQueryAllChainRequest };
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
        const message = { ...baseQueryAllChainRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryAllChainResponse = {};
export const QueryAllChainResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.Chain) {
            Chain.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllChainResponse };
        message.Chain = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.Chain.push(Chain.decode(reader, reader.uint32()));
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
        const message = { ...baseQueryAllChainResponse };
        message.Chain = [];
        if (object.Chain !== undefined && object.Chain !== null) {
            for (const e of object.Chain) {
                message.Chain.push(Chain.fromJSON(e));
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
        if (message.Chain) {
            obj.Chain = message.Chain.map((e) => (e ? Chain.toJSON(e) : undefined));
        }
        else {
            obj.Chain = [];
        }
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllChainResponse };
        message.Chain = [];
        if (object.Chain !== undefined && object.Chain !== null) {
            for (const e of object.Chain) {
                message.Chain.push(Chain.fromPartial(e));
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
    Chain(request) {
        const data = QueryGetChainRequest.encode(request).finish();
        const promise = this.rpc.request('pchainorg.pibridge.chain.Query', 'Chain', data);
        return promise.then((data) => QueryGetChainResponse.decode(new Reader(data)));
    }
    ChainAll(request) {
        const data = QueryAllChainRequest.encode(request).finish();
        const promise = this.rpc.request('pchainorg.pibridge.chain.Query', 'ChainAll', data);
        return promise.then((data) => QueryAllChainResponse.decode(new Reader(data)));
    }
}
