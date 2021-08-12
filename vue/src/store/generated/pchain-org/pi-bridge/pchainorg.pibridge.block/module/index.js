// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgUpdateBlock } from "./types/block/tx";
import { MsgDeleteBlock } from "./types/block/tx";
import { MsgCreateBlock } from "./types/block/tx";
const types = [
    ["/pchainorg.pibridge.block.MsgUpdateBlock", MsgUpdateBlock],
    ["/pchainorg.pibridge.block.MsgDeleteBlock", MsgDeleteBlock],
    ["/pchainorg.pibridge.block.MsgCreateBlock", MsgCreateBlock],
];
export const MissingWalletError = new Error("wallet is required");
const registry = new Registry(types);
const defaultFee = {
    amount: [],
    gas: "200000",
};
const txClient = async (wallet, { addr: addr } = { addr: "http://localhost:26657" }) => {
    if (!wallet)
        throw MissingWalletError;
    const client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
    const { address } = (await wallet.getAccounts())[0];
    return {
        signAndBroadcast: (msgs, { fee, memo } = { fee: defaultFee, memo: "" }) => client.signAndBroadcast(address, msgs, fee, memo),
        msgUpdateBlock: (data) => ({ typeUrl: "/pchainorg.pibridge.block.MsgUpdateBlock", value: data }),
        msgDeleteBlock: (data) => ({ typeUrl: "/pchainorg.pibridge.block.MsgDeleteBlock", value: data }),
        msgCreateBlock: (data) => ({ typeUrl: "/pchainorg.pibridge.block.MsgCreateBlock", value: data }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
