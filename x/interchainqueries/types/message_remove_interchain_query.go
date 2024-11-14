package types

import (
	"strings"

	"cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRemoveInterchainQuery{}

func NewMsgRemoveInterchainQuery(sender string, queryID uint64) MsgRemoveInterchainQuery {
	return MsgRemoveInterchainQuery{
		QueryId: queryID,
		Sender:  sender,
	}
}

func (msg MsgRemoveInterchainQuery) Route() string {
	return RouterKey
}

func (msg MsgRemoveInterchainQuery) Type() string {
	return "remove-interchain-query"
}

func (msg MsgRemoveInterchainQuery) Validate() error {
	if msg.GetQueryId() == 0 {
		return errors.Wrap(ErrInvalidQueryID, "query_id cannot be empty or equal to 0")
	}

	if strings.TrimSpace(msg.Sender) == "" {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, "missing sender address")
	}

	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "failed to parse address: %s", msg.Sender)
	}

	return nil
}

func (msg MsgRemoveInterchainQuery) GetSignBytes() []byte {
	return ModuleCdc.MustMarshalJSON(&msg)
}

func (msg MsgRemoveInterchainQuery) GetSigners() []sdk.AccAddress {
	senderAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil { // should never happen as valid basic rejects invalid addresses
		panic(err.Error())
	}
	return []sdk.AccAddress{senderAddr}
}
