// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgDeleteNode } from "./types/node/tx";
import { MsgCreateNode } from "./types/node/tx";
import { MsgUpdateNode } from "./types/node/tx";
const types = [
    ["/pchainorg.pibridge.node.MsgDeleteNode", MsgDeleteNode],
    ["/pchainorg.pibridge.node.MsgCreateNode", MsgCreateNode],
    ["/pchainorg.pibridge.node.MsgUpdateNode", MsgUpdateNode],
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
        msgDeleteNode: (data) => ({ typeUrl: "/pchainorg.pibridge.node.MsgDeleteNode", value: data }),
        msgCreateNode: (data) => ({ typeUrl: "/pchainorg.pibridge.node.MsgCreateNode", value: data }),
        msgUpdateNode: (data) => ({ typeUrl: "/pchainorg.pibridge.node.MsgUpdateNode", value: data }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
