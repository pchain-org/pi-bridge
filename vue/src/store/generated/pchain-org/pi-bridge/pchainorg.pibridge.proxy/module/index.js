// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgCreateProxy } from "./types/proxy/tx";
import { MsgUpdateProxy } from "./types/proxy/tx";
import { MsgDeleteProxy } from "./types/proxy/tx";
const types = [
    ["/pchainorg.pibridge.proxy.MsgCreateProxy", MsgCreateProxy],
    ["/pchainorg.pibridge.proxy.MsgUpdateProxy", MsgUpdateProxy],
    ["/pchainorg.pibridge.proxy.MsgDeleteProxy", MsgDeleteProxy],
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
        msgCreateProxy: (data) => ({ typeUrl: "/pchainorg.pibridge.proxy.MsgCreateProxy", value: data }),
        msgUpdateProxy: (data) => ({ typeUrl: "/pchainorg.pibridge.proxy.MsgUpdateProxy", value: data }),
        msgDeleteProxy: (data) => ({ typeUrl: "/pchainorg.pibridge.proxy.MsgDeleteProxy", value: data }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
