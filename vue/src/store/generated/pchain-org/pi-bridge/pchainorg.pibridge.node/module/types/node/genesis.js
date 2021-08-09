/* eslint-disable */
import { Node } from '../node/node';
import { Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'pchainorg.pibridge.node';
const baseGenesisState = {};
export const GenesisState = {
    encode(message, writer = Writer.create()) {
        for (const v of message.nodeList) {
            Node.encode(v, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseGenesisState };
        message.nodeList = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.nodeList.push(Node.decode(reader, reader.uint32()));
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
        message.nodeList = [];
        if (object.nodeList !== undefined && object.nodeList !== null) {
            for (const e of object.nodeList) {
                message.nodeList.push(Node.fromJSON(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.nodeList) {
            obj.nodeList = message.nodeList.map((e) => (e ? Node.toJSON(e) : undefined));
        }
        else {
            obj.nodeList = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseGenesisState };
        message.nodeList = [];
        if (object.nodeList !== undefined && object.nodeList !== null) {
            for (const e of object.nodeList) {
                message.nodeList.push(Node.fromPartial(e));
            }
        }
        return message;
    }
};
