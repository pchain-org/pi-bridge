/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal';
import { Proxy } from '../proxy/proxy';
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination';
export const protobufPackage = 'pchainorg.pibridge.proxy';
const baseQueryGetProxyRequest = { index: '' };
export const QueryGetProxyRequest = {
    encode(message, writer = Writer.create()) {
        if (message.index !== '') {
            writer.uint32(10).string(message.index);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetProxyRequest };
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
        const message = { ...baseQueryGetProxyRequest };
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
        const message = { ...baseQueryGetProxyRequest };
        if (object.index !== undefined && object.index !== null) {
            message.index = object.index;
        }
        else {
            message.index = '';
        }
        return message;
    }
};
const baseQueryGetProxyResponse = {};
export const QueryGetProxyResponse = {
    encode(message, writer = Writer.create()) {
        if (message.Proxy !== undefined) {
            Proxy.encode(message.Proxy, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetProxyResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.Proxy = Proxy.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetProxyResponse };
        if (object.Proxy !== undefined && object.Proxy !== null) {
            message.Proxy = Proxy.fromJSON(object.Proxy);
        }
        else {
            message.Proxy = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.Proxy !== undefined && (obj.Proxy = message.Proxy ? Proxy.toJSON(message.Proxy) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetProxyResponse };
        if (object.Proxy !== undefined && object.Proxy !== null) {
            message.Proxy = Proxy.fromPartial(object.Proxy);
        }
        else {
            message.Proxy = undefined;
        }
        return message;
    }
};
const baseQueryAllProxyRequest = {};
export const QueryAllProxyRequest = {
    encode(message, writer = Writer.create()) {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllProxyRequest };
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
        const message = { ...baseQueryAllProxyRequest };
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
        const message = { ...baseQueryAllProxyRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryAllProxyResponse = {};
export const QueryAllProxyResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.Proxy) {
            Proxy.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllProxyResponse };
        message.Proxy = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.Proxy.push(Proxy.decode(reader, reader.uint32()));
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
        const message = { ...baseQueryAllProxyResponse };
        message.Proxy = [];
        if (object.Proxy !== undefined && object.Proxy !== null) {
            for (const e of object.Proxy) {
                message.Proxy.push(Proxy.fromJSON(e));
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
        if (message.Proxy) {
            obj.Proxy = message.Proxy.map((e) => (e ? Proxy.toJSON(e) : undefined));
        }
        else {
            obj.Proxy = [];
        }
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllProxyResponse };
        message.Proxy = [];
        if (object.Proxy !== undefined && object.Proxy !== null) {
            for (const e of object.Proxy) {
                message.Proxy.push(Proxy.fromPartial(e));
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
    Proxy(request) {
        const data = QueryGetProxyRequest.encode(request).finish();
        const promise = this.rpc.request('pchainorg.pibridge.proxy.Query', 'Proxy', data);
        return promise.then((data) => QueryGetProxyResponse.decode(new Reader(data)));
    }
    ProxyAll(request) {
        const data = QueryAllProxyRequest.encode(request).finish();
        const promise = this.rpc.request('pchainorg.pibridge.proxy.Query', 'ProxyAll', data);
        return promise.then((data) => QueryAllProxyResponse.decode(new Reader(data)));
    }
}
