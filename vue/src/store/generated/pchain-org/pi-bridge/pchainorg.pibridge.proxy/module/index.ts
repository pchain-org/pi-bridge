// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
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

const registry = new Registry(<any>types);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
}

interface SignAndBroadcastOptions {
  fee: StdFee,
  memo?: string
}

const txClient = async (wallet: OfflineSigner, { addr: addr }: TxClientOptions = { addr: "http://localhost:26657" }) => {
  if (!wallet) throw MissingWalletError;

  const client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions = {fee: defaultFee, memo: ""}) => client.signAndBroadcast(address, msgs, fee,memo),
    msgCreateProxy: (data: MsgCreateProxy): EncodeObject => ({ typeUrl: "/pchainorg.pibridge.proxy.MsgCreateProxy", value: data }),
    msgUpdateProxy: (data: MsgUpdateProxy): EncodeObject => ({ typeUrl: "/pchainorg.pibridge.proxy.MsgUpdateProxy", value: data }),
    msgDeleteProxy: (data: MsgDeleteProxy): EncodeObject => ({ typeUrl: "/pchainorg.pibridge.proxy.MsgDeleteProxy", value: data }),
    
  };
};

interface QueryClientOptions {
  addr: string
}

const queryClient = async ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

export {
  txClient,
  queryClient,
};
