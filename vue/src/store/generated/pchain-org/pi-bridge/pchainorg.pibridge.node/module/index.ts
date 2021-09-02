// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgCreateNode } from "./types/node/tx";
import { MsgUpdateNode } from "./types/node/tx";
import { MsgDeleteNode } from "./types/node/tx";


const types = [
  ["/pchainorg.pibridge.node.MsgCreateNode", MsgCreateNode],
  ["/pchainorg.pibridge.node.MsgUpdateNode", MsgUpdateNode],
  ["/pchainorg.pibridge.node.MsgDeleteNode", MsgDeleteNode],
  
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
    msgCreateNode: (data: MsgCreateNode): EncodeObject => ({ typeUrl: "/pchainorg.pibridge.node.MsgCreateNode", value: data }),
    msgUpdateNode: (data: MsgUpdateNode): EncodeObject => ({ typeUrl: "/pchainorg.pibridge.node.MsgUpdateNode", value: data }),
    msgDeleteNode: (data: MsgDeleteNode): EncodeObject => ({ typeUrl: "/pchainorg.pibridge.node.MsgDeleteNode", value: data }),
    
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
