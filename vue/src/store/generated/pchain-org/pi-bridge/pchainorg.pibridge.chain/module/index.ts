// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgCreateChain } from "./types/chain/tx";
import { MsgUpdateChain } from "./types/chain/tx";
import { MsgDeleteChain } from "./types/chain/tx";


const types = [
  ["/pchainorg.pibridge.chain.MsgCreateChain", MsgCreateChain],
  ["/pchainorg.pibridge.chain.MsgUpdateChain", MsgUpdateChain],
  ["/pchainorg.pibridge.chain.MsgDeleteChain", MsgDeleteChain],
  
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
    msgCreateChain: (data: MsgCreateChain): EncodeObject => ({ typeUrl: "/pchainorg.pibridge.chain.MsgCreateChain", value: data }),
    msgUpdateChain: (data: MsgUpdateChain): EncodeObject => ({ typeUrl: "/pchainorg.pibridge.chain.MsgUpdateChain", value: data }),
    msgDeleteChain: (data: MsgDeleteChain): EncodeObject => ({ typeUrl: "/pchainorg.pibridge.chain.MsgDeleteChain", value: data }),
    
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
