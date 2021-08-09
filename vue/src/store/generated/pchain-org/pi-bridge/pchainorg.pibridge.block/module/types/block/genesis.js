/* eslint-disable */
import { Block } from '../block/block';
import { Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'pchainorg.pibridge.block';
const baseGenesisState = {};
export const GenesisState = {
    encode(message, writer = Writer.create()) {
        for (const v of message.blockList) {
            Block.encode(v, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseGenesisState };
        message.blockList = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.blockList.push(Block.decode(reader, reader.uint32()));
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
        message.blockList = [];
        if (object.blockList !== undefined && object.blockList !== null) {
            for (const e of object.blockList) {
                message.blockList.push(Block.fromJSON(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.blockList) {
            obj.blockList = message.blockList.map((e) => (e ? Block.toJSON(e) : undefined));
        }
        else {
            obj.blockList = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseGenesisState };
        message.blockList = [];
        if (object.blockList !== undefined && object.blockList !== null) {
            for (const e of object.blockList) {
                message.blockList.push(Block.fromPartial(e));
            }
        }
        return message;
    }
};
