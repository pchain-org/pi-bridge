// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgDeleteChain } from "./types/chain/tx";
import { MsgCreateChain } from "./types/chain/tx";
import { MsgUpdateChain } from "./types/chain/tx";
const types = [
    ["/pchainorg.pibridge.chain.MsgDeleteChain", MsgDeleteChain],
    ["/pchainorg.pibridge.chain.MsgCreateChain", MsgCreateChain],
    ["/pchainorg.pibridge.chain.MsgUpdateChain", MsgUpdateChain],
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
        msgDeleteChain: (data) => ({ typeUrl: "/pchainorg.pibridge.chain.MsgDeleteChain", value: data }),
        msgCreateChain: (data) => ({ typeUrl: "/pchainorg.pibridge.chain.MsgCreateChain", value: data }),
        msgUpdateChain: (data) => ({ typeUrl: "/pchainorg.pibridge.chain.MsgUpdateChain", value: data }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
