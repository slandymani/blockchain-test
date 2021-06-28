package blockchaintest

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/slandymani/blockchain-test/x/blockchaintest/keeper"
	"github.com/slandymani/blockchain-test/x/blockchaintest/types"
)

func handleMsgSetName(ctx sdk.Context, k keeper.Keeper, msg types.MsgSetName) (*sdk.Result, error) {
	if !msg.Owner.Equals(k.GetCreator(ctx, msg.Name)) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
	}

	k.SetName(ctx, msg.Name, msg.Value)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
