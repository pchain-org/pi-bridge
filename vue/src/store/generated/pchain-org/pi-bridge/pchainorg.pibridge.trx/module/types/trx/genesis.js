/* eslint-disable */
import { Trx } from '../trx/trx';
import { Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'pchainorg.pibridge.trx';
const baseGenesisState = {};
export const GenesisState = {
    encode(message, writer = Writer.create()) {
        for (const v of message.trxList) {
            Trx.encode(v, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseGenesisState };
        message.trxList = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.trxList.push(Trx.decode(reader, reader.uint32()));
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
        message.trxList = [];
        if (object.trxList !== undefined && object.trxList !== null) {
            for (const e of object.trxList) {
                message.trxList.push(Trx.fromJSON(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.trxList) {
            obj.trxList = message.trxList.map((e) => (e ? Trx.toJSON(e) : undefined));
        }
        else {
            obj.trxList = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseGenesisState };
        message.trxList = [];
        if (object.trxList !== undefined && object.trxList !== null) {
            for (const e of object.trxList) {
                message.trxList.push(Trx.fromPartial(e));
            }
        }
        return message;
    }
};
