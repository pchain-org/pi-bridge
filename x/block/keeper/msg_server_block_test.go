package keeper

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/pchain-org/pi-bridge/x/block/types"
)

func TestBlockMsgServerCreate(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	srv := NewMsgServerImpl(*keeper)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		idx := fmt.Sprintf("%d", i)
		expected := &types.MsgCreateBlock{Creator: creator, Index: idx}
		_, err := srv.CreateBlock(wctx, expected)
		require.NoError(t, err)
		rst, found := keeper.GetBlock(ctx, expected.Index)
		require.True(t, found)
		assert.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestBlockMsgServerUpdate(t *testing.T) {
	creator := "A"
	index := "any"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateBlock
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateBlock{Creator: creator, Index: index},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateBlock{Creator: "B", Index: index},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgUpdateBlock{Creator: creator, Index: "missing"},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			keeper, ctx := setupKeeper(t)
			srv := NewMsgServerImpl(*keeper)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateBlock{Creator: creator, Index: index}
			_, err := srv.CreateBlock(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateBlock(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := keeper.GetBlock(ctx, expected.Index)
				require.True(t, found)
				assert.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestBlockMsgServerDelete(t *testing.T) {
	creator := "A"
	index := "any"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteBlock
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteBlock{Creator: creator, Index: index},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteBlock{Creator: "B", Index: index},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteBlock{Creator: creator, Index: "missing"},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			keeper, ctx := setupKeeper(t)
			srv := NewMsgServerImpl(*keeper)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateBlock(wctx, &types.MsgCreateBlock{Creator: creator, Index: index})
			require.NoError(t, err)
			_, err = srv.DeleteBlock(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := keeper.GetBlock(ctx, tc.request.Index)
				require.False(t, found)
			}
		})
	}
}
