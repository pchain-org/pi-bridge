import { StdFee } from "@cosmjs/launchpad";
import { OfflineSigner, EncodeObject } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgDeleteTrx } from "./types/trx/tx";
import { MsgUpdateTrx } from "./types/trx/tx";
import { MsgCrossChainTransfer } from "./types/trx/tx";
import { MsgCreateTrx } from "./types/trx/tx";
export declare const MissingWalletError: Error;
interface TxClientOptions {
    addr: string;
}
interface SignAndBroadcastOptions {
    fee: StdFee;
    memo?: string;
}
declare const txClient: (wallet: OfflineSigner, { addr: addr }?: TxClientOptions) => Promise<{
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }?: SignAndBroadcastOptions) => Promise<import("@cosmjs/stargate").BroadcastTxResponse>;
    msgDeleteTrx: (data: MsgDeleteTrx) => EncodeObject;
    msgUpdateTrx: (data: MsgUpdateTrx) => EncodeObject;
    msgCrossChainTransfer: (data: MsgCrossChainTransfer) => EncodeObject;
    msgCreateTrx: (data: MsgCreateTrx) => EncodeObject;
}>;
interface QueryClientOptions {
    addr: string;
}
declare const queryClient: ({ addr: addr }?: QueryClientOptions) => Promise<Api<unknown>>;
export { txClient, queryClient, };
