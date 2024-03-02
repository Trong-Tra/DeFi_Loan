package types

import (
	"strconv"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRequestLoan{}

func NewMsgRequestLoan(creator string, amount string, fee string, collateral string, deadline string) *MsgRequestLoan {
	return &MsgRequestLoan{
		Creator:    creator,
		Amount:     amount,
		Fee:        fee,
		Collateral: collateral,
		Deadline:   deadline,
	}
}

func (msg *MsgRequestLoan) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid creator address (%s)", err)
	}
	amount, _ := sdk.ParseCoinNormalized(msg.Amount)
	if !amount.IsValid() {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "Amount is not a valid Coins object")
	}
	if amount.IsZero() {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "Amount is empty")
	}
	fee, _ := sdk.ParseCoinNormalized(msg.Fee)
	if !fee.IsValid() {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "Fee is not a valid Coins object")
	}
	deadline, err := strconv.ParseInt(msg.Deadline, 10, 64)
	if err != nil {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "Deadline is not an integer")
	}
	if deadline <= 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "Deadline need to be a positive integer")
	}
	collateral, _ := sdk.ParseCoinNormalized(msg.Collateral)
	if !collateral.IsValid() {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "Collateral is not a valid Coiins object")
	}
	if collateral.IsZero() {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "collateral is empty")
	}
	return nil
}
