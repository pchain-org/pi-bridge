/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal';
import { Trx } from '../trx/trx';
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination';
export const protobufPackage = 'pchainorg.pibridge.trx';
const baseQueryGetTrxRequest = { index: '' };
export const QueryGetTrxRequest = {
    encode(message, writer = Writer.create()) {
        if (message.index !== '') {
            writer.uint32(10).string(message.index);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetTrxRequest };
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
        const message = { ...baseQueryGetTrxRequest };
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
        const message = { ...baseQueryGetTrxRequest };
        if (object.index !== undefined && object.index !== null) {
            message.index = object.index;
        }
        else {
            message.index = '';
        }
        return message;
    }
};
const baseQueryGetTrxResponse = {};
export const QueryGetTrxResponse = {
    encode(message, writer = Writer.create()) {
        if (message.Trx !== undefined) {
            Trx.encode(message.Trx, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetTrxResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.Trx = Trx.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetTrxResponse };
        if (object.Trx !== undefined && object.Trx !== null) {
            message.Trx = Trx.fromJSON(object.Trx);
        }
        else {
            message.Trx = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.Trx !== undefined && (obj.Trx = message.Trx ? Trx.toJSON(message.Trx) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetTrxResponse };
        if (object.Trx !== undefined && object.Trx !== null) {
            message.Trx = Trx.fromPartial(object.Trx);
        }
        else {
            message.Trx = undefined;
        }
        return message;
    }
};
const baseQueryAllTrxRequest = {};
export const QueryAllTrxRequest = {
    encode(message, writer = Writer.create()) {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllTrxRequest };
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
        const message = { ...baseQueryAllTrxRequest };
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
        const message = { ...baseQueryAllTrxRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryAllTrxResponse = {};
export const QueryAllTrxResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.Trx) {
            Trx.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllTrxResponse };
        message.Trx = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.Trx.push(Trx.decode(reader, reader.uint32()));
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
        const message = { ...baseQueryAllTrxResponse };
        message.Trx = [];
        if (object.Trx !== undefined && object.Trx !== null) {
            for (const e of object.Trx) {
                message.Trx.push(Trx.fromJSON(e));
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
        if (message.Trx) {
            obj.Trx = message.Trx.map((e) => (e ? Trx.toJSON(e) : undefined));
        }
        else {
            obj.Trx = [];
        }
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllTrxResponse };
        message.Trx = [];
        if (object.Trx !== undefined && object.Trx !== null) {
            for (const e of object.Trx) {
                message.Trx.push(Trx.fromPartial(e));
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
    Trx(request) {
        const data = QueryGetTrxRequest.encode(request).finish();
        const promise = this.rpc.request('pchainorg.pibridge.trx.Query', 'Trx', data);
        return promise.then((data) => QueryGetTrxResponse.decode(new Reader(data)));
    }
    TrxAll(request) {
        const data = QueryAllTrxRequest.encode(request).finish();
        const promise = this.rpc.request('pchainorg.pibridge.trx.Query', 'TrxAll', data);
        return promise.then((data) => QueryAllTrxResponse.decode(new Reader(data)));
    }
}
