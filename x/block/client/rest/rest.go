package rest

import (
	"fmt"
	"net/http"

	sdkclient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/pchain-org/pi-bridge/x/block/types"

	"github.com/gorilla/mux"
)

const (
	ContractAddr = "contract-address"
	TokenId      = "token-id"
	MethodName   = "method-name"
	ReturnType   = "return-type"
	OrgId        = "org-id"
)

// RegisterRoutes register distribution REST routes.
func RegisterRoutes(cliCtx sdkclient.Context, r *mux.Router) {
	// registerQueryRoutes(cliCtx, r)
	r.HandleFunc("/block/create", BlockCreateReqHandlerFn(cliCtx)).Methods("POST")
}

type blockCreateReq struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Creator string       `json:"creator,omitempty"`
	Index   string       `json:"index,omitempty"`
	ChainID int32        `json:"chainID,omitempty"`
	Address string       `json:"address,omitempty"`
	Headers []string     `son:"headers,omitempty"`
}

// ContractCallReqHandlerFn - http request handler to call a contract.
func BlockCreateReqHandlerFn(cliCtx sdkclient.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("1234567890")

		var req blockCreateReq
		if !rest.ReadRESTReq(w, r, cliCtx.LegacyAmino, &req) {
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		msg := types.NewMsgCreateBlock(req.Creator, req.Index, req.ChainID, req.Address, req.Headers)
		err := msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		tx.WriteGeneratedTxResponse(cliCtx, w, req.BaseReq, msg)
	}
}
