// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgDeleteTrx } from "./types/trx/tx";
import { MsgUpdateTrx } from "./types/trx/tx";
import { MsgCrossChainTransfer } from "./types/trx/tx";
import { MsgCreateTrx } from "./types/trx/tx";


const types = [
  ["/pchainorg.pibridge.trx.MsgDeleteTrx", MsgDeleteTrx],
  ["/pchainorg.pibridge.trx.MsgUpdateTrx", MsgUpdateTrx],
  ["/pchainorg.pibridge.trx.MsgCrossChainTransfer", MsgCrossChainTransfer],
  ["/pchainorg.pibridge.trx.MsgCreateTrx", MsgCreateTrx],
  
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
    msgDeleteTrx: (data: MsgDeleteTrx): EncodeObject => ({ typeUrl: "/pchainorg.pibridge.trx.MsgDeleteTrx", value: data }),
    msgUpdateTrx: (data: MsgUpdateTrx): EncodeObject => ({ typeUrl: "/pchainorg.pibridge.trx.MsgUpdateTrx", value: data }),
    msgCrossChainTransfer: (data: MsgCrossChainTransfer): EncodeObject => ({ typeUrl: "/pchainorg.pibridge.trx.MsgCrossChainTransfer", value: data }),
    msgCreateTrx: (data: MsgCreateTrx): EncodeObject => ({ typeUrl: "/pchainorg.pibridge.trx.MsgCreateTrx", value: data }),
    
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
