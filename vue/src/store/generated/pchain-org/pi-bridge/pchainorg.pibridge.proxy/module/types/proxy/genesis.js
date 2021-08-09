/* eslint-disable */
import { Proxy } from '../proxy/proxy';
import { Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'pchainorg.pibridge.proxy';
const baseGenesisState = {};
export const GenesisState = {
    encode(message, writer = Writer.create()) {
        for (const v of message.proxyList) {
            Proxy.encode(v, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseGenesisState };
        message.proxyList = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.proxyList.push(Proxy.decode(reader, reader.uint32()));
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseGenesisState };
        message.proxyList = [];
        if (object.proxyList !== undefined && object.proxyList !== null) {
            for (const e of object.proxyList) {
                message.proxyList.push(Proxy.fromJSON(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.proxyList) {
            obj.proxyList = message.proxyList.map((e) => (e ? Proxy.toJSON(e) : undefined));
        }
        else {
            obj.proxyList = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseGenesisState };
        message.proxyList = [];
        if (object.proxyList !== undefined && object.proxyList !== null) {
            for (const e of object.proxyList) {
                message.proxyList.push(Proxy.fromPartial(e));
            }
        }
        return message;
    }
};
