package blockchaintest

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/slandymani/blockchain-test/x/blockchaintest/keeper"
	"github.com/slandymani/blockchain-test/x/blockchaintest/types"
)

func handleMsgCreateWhois(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateWhois) (*sdk.Result, error) {
	k.CreateWhois(ctx, msg)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
