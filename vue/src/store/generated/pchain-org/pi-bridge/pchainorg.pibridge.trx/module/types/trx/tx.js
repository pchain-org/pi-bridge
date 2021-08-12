/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal';
export const protobufPackage = 'pchainorg.pibridge.trx';
const baseMsgCrossChainTransfer = { creator: '', asset: '', amount: 0, fromChain: '', fromAddress: '', toChain: '', toAddress: '' };
export const MsgCrossChainTransfer = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== '') {
            writer.uint32(10).string(message.creator);
        }
        if (message.asset !== '') {
            writer.uint32(18).string(message.asset);
        }
        if (message.amount !== 0) {
            writer.uint32(24).int32(message.amount);
        }
        if (message.fromChain !== '') {
            writer.uint32(34).string(message.fromChain);
        }
        if (message.fromAddress !== '') {
            writer.uint32(42).string(message.fromAddress);
        }
        if (message.toChain !== '') {
            writer.uint32(50).string(message.toChain);
        }
        if (message.toAddress !== '') {
            writer.uint32(58).string(message.toAddress);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgCrossChainTransfer };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.asset = reader.string();
                    break;
                case 3:
                    message.amount = reader.int32();
                    break;
                case 4:
                    message.fromChain = reader.string();
                    break;
                case 5:
                    message.fromAddress = reader.string();
                    break;
                case 6:
                    message.toChain = reader.string();
                    break;
                case 7:
                    message.toAddress = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgCrossChainTransfer };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = '';
        }
        if (object.asset !== undefined && object.asset !== null) {
            message.asset = String(object.asset);
        }
        else {
            message.asset = '';
        }
        if (object.amount !== undefined && object.amount !== null) {
            message.amount = Number(object.amount);
        }
        else {
            message.amount = 0;
        }
        if (object.fromChain !== undefined && object.fromChain !== null) {
            message.fromChain = String(object.fromChain);
        }
        else {
            message.fromChain = '';
        }
        if (object.fromAddress !== undefined && object.fromAddress !== null) {
            message.fromAddress = String(object.fromAddress);
        }
        else {
            message.fromAddress = '';
        }
        if (object.toChain !== undefined && object.toChain !== null) {
            message.toChain = String(object.toChain);
        }
        else {
            message.toChain = '';
        }
        if (object.toAddress !== undefined && object.toAddress !== null) {
            message.toAddress = String(object.toAddress);
        }
        else {
            message.toAddress = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.asset !== undefined && (obj.asset = message.asset);
        message.amount !== undefined && (obj.amount = message.amount);
        message.fromChain !== undefined && (obj.fromChain = message.fromChain);
        message.fromAddress !== undefined && (obj.fromAddress = message.fromAddress);
        message.toChain !== undefined && (obj.toChain = message.toChain);
        message.toAddress !== undefined && (obj.toAddress = message.toAddress);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgCrossChainTransfer };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = '';
        }
        if (object.asset !== undefined && object.asset !== null) {
            message.asset = object.asset;
        }
        else {
            message.asset = '';
        }
        if (object.amount !== undefined && object.amount !== null) {
            message.amount = object.amount;
        }
        else {
            message.amount = 0;
        }
        if (object.fromChain !== undefined && object.fromChain !== null) {
            message.fromChain = object.fromChain;
        }
        else {
            message.fromChain = '';
        }
        if (object.fromAddress !== undefined && object.fromAddress !== null) {
            message.fromAddress = object.fromAddress;
        }
        else {
            message.fromAddress = '';
        }
        if (object.toChain !== undefined && object.toChain !== null) {
            message.toChain = object.toChain;
        }
        else {
            message.toChain = '';
        }
        if (object.toAddress !== undefined && object.toAddress !== null) {
            message.toAddress = object.toAddress;
        }
        else {
            message.toAddress = '';
        }
        return message;
    }
};
const baseMsgCrossChainTransferResponse = {};
export const MsgCrossChainTransferResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgCrossChainTransferResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(_) {
        const message = { ...baseMsgCrossChainTransferResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgCrossChainTransferResponse };
        return message;
    }
};
const baseMsgCreateTrx = { creator: '', index: '', chainID: 0, address: '', crossChainMsgs: '', proof: '' };
export const MsgCreateTrx = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== '') {
            writer.uint32(10).string(message.creator);
        }
        if (message.index !== '') {
            writer.uint32(18).string(message.index);
        }
        if (message.chainID !== 0) {
            writer.uint32(24).int32(message.chainID);
        }
        if (message.address !== '') {
            writer.uint32(34).string(message.address);
        }
        if (message.crossChainMsgs !== '') {
            writer.uint32(42).string(message.crossChainMsgs);
        }
        if (message.proof !== '') {
            writer.uint32(50).string(message.proof);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgCreateTrx };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.index = reader.string();
                    break;
                case 3:
                    message.chainID = reader.int32();
                    break;
                case 4:
                    message.address = reader.string();
                    break;
                case 5:
                    message.crossChainMsgs = reader.string();
                    break;
                case 6:
                    message.proof = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgCreateTrx };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = '';
        }
        if (object.index !== undefined && object.index !== null) {
            message.index = String(object.index);
        }
        else {
            message.index = '';
        }
        if (object.chainID !== undefined && object.chainID !== null) {
            message.chainID = Number(object.chainID);
        }
        else {
            message.chainID = 0;
        }
        if (object.address !== undefined && object.address !== null) {
            message.address = String(object.address);
        }
        else {
            message.address = '';
        }
        if (object.crossChainMsgs !== undefined && object.crossChainMsgs !== null) {
            message.crossChainMsgs = String(object.crossChainMsgs);
        }
        else {
            message.crossChainMsgs = '';
        }
        if (object.proof !== undefined && object.proof !== null) {
            message.proof = String(object.proof);
        }
        else {
            message.proof = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.index !== undefined && (obj.index = message.index);
        message.chainID !== undefined && (obj.chainID = message.chainID);
        message.address !== undefined && (obj.address = message.address);
        message.crossChainMsgs !== undefined && (obj.crossChainMsgs = message.crossChainMsgs);
        message.proof !== undefined && (obj.proof = message.proof);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgCreateTrx };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = '';
        }
        if (object.index !== undefined && object.index !== null) {
            message.index = object.index;
        }
        else {
            message.index = '';
        }
        if (object.chainID !== undefined && object.chainID !== null) {
            message.chainID = object.chainID;
        }
        else {
            message.chainID = 0;
        }
        if (object.address !== undefined && object.address !== null) {
            message.address = object.address;
        }
        else {
            message.address = '';
        }
        if (object.crossChainMsgs !== undefined && object.crossChainMsgs !== null) {
            message.crossChainMsgs = object.crossChainMsgs;
        }
        else {
            message.crossChainMsgs = '';
        }
        if (object.proof !== undefined && object.proof !== null) {
            message.proof = object.proof;
        }
        else {
            message.proof = '';
        }
        return message;
    }
};
const baseMsgCreateTrxResponse = {};
export const MsgCreateTrxResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgCreateTrxResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(_) {
        const message = { ...baseMsgCreateTrxResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgCreateTrxResponse };
        return message;
    }
};
const baseMsgUpdateTrx = { creator: '', index: '', chainID: 0, address: '', crossChainMsgs: '' };
export const MsgUpdateTrx = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== '') {
            writer.uint32(10).string(message.creator);
        }
        if (message.index !== '') {
            writer.uint32(18).string(message.index);
        }
        if (message.chainID !== 0) {
            writer.uint32(24).int32(message.chainID);
        }
        if (message.address !== '') {
            writer.uint32(34).string(message.address);
        }
        if (message.crossChainMsgs !== '') {
            writer.uint32(42).string(message.crossChainMsgs);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgUpdateTrx };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.index = reader.string();
                    break;
                case 3:
                    message.chainID = reader.int32();
                    break;
                case 4:
                    message.address = reader.string();
                    break;
                case 5:
                    message.crossChainMsgs = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgUpdateTrx };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = '';
        }
        if (object.index !== undefined && object.index !== null) {
            message.index = String(object.index);
        }
        else {
            message.index = '';
        }
        if (object.chainID !== undefined && object.chainID !== null) {
            message.chainID = Number(object.chainID);
        }
        else {
            message.chainID = 0;
        }
        if (object.address !== undefined && object.address !== null) {
            message.address = String(object.address);
        }
        else {
            message.address = '';
        }
        if (object.crossChainMsgs !== undefined && object.crossChainMsgs !== null) {
            message.crossChainMsgs = String(object.crossChainMsgs);
        }
        else {
            message.crossChainMsgs = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.index !== undefined && (obj.index = message.index);
        message.chainID !== undefined && (obj.chainID = message.chainID);
        message.address !== undefined && (obj.address = message.address);
        message.crossChainMsgs !== undefined && (obj.crossChainMsgs = message.crossChainMsgs);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgUpdateTrx };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = '';
        }
        if (object.index !== undefined && object.index !== null) {
            message.index = object.index;
        }
        else {
            message.index = '';
        }
        if (object.chainID !== undefined && object.chainID !== null) {
            message.chainID = object.chainID;
        }
        else {
            message.chainID = 0;
        }
        if (object.address !== undefined && object.address !== null) {
            message.address = object.address;
        }
        else {
            message.address = '';
        }
        if (object.crossChainMsgs !== undefined && object.crossChainMsgs !== null) {
            message.crossChainMsgs = object.crossChainMsgs;
        }
        else {
            message.crossChainMsgs = '';
        }
        return message;
    }
};
const baseMsgUpdateTrxResponse = {};
export const MsgUpdateTrxResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgUpdateTrxResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(_) {
        const message = { ...baseMsgUpdateTrxResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgUpdateTrxResponse };
        return message;
    }
};
const baseMsgDeleteTrx = { creator: '', index: '' };
export const MsgDeleteTrx = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== '') {
            writer.uint32(10).string(message.creator);
        }
        if (message.index !== '') {
            writer.uint32(18).string(message.index);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgDeleteTrx };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
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
        const message = { ...baseMsgDeleteTrx };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = '';
        }
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
        message.creator !== undefined && (obj.creator = message.creator);
        message.index !== undefined && (obj.index = message.index);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgDeleteTrx };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = '';
        }
        if (object.index !== undefined && object.index !== null) {
            message.index = object.index;
        }
        else {
            message.index = '';
        }
        return message;
    }
};
const baseMsgDeleteTrxResponse = {};
export const MsgDeleteTrxResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgDeleteTrxResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(_) {
        const message = { ...baseMsgDeleteTrxResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgDeleteTrxResponse };
        return message;
    }
};
export class MsgClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    CrossChainTransfer(request) {
        const data = MsgCrossChainTransfer.encode(request).finish();
        const promise = this.rpc.request('pchainorg.pibridge.trx.Msg', 'CrossChainTransfer', data);
        return promise.then((data) => MsgCrossChainTransferResponse.decode(new Reader(data)));
    }
    CreateTrx(request) {
        const data = MsgCreateTrx.encode(request).finish();
        const promise = this.rpc.request('pchainorg.pibridge.trx.Msg', 'CreateTrx', data);
        return promise.then((data) => MsgCreateTrxResponse.decode(new Reader(data)));
    }
    UpdateTrx(request) {
        const data = MsgUpdateTrx.encode(request).finish();
        const promise = this.rpc.request('pchainorg.pibridge.trx.Msg', 'UpdateTrx', data);
        return promise.then((data) => MsgUpdateTrxResponse.decode(new Reader(data)));
    }
    DeleteTrx(request) {
        const data = MsgDeleteTrx.encode(request).finish();
        const promise = this.rpc.request('pchainorg.pibridge.trx.Msg', 'DeleteTrx', data);
        return promise.then((data) => MsgDeleteTrxResponse.decode(new Reader(data)));
    }
}
