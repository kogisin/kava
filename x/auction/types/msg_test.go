package types

import (
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestMsgPlaceBid_ValidateBasic(t *testing.T) {
	addr := sdk.AccAddress([]byte("someName"))
	tests := []struct {
		name       string
		msg        MsgPlaceBid
		expectPass bool
	}{
		{"normal",
			NewMsgPlaceBid(0, addr, c("token", 10)),
			true},
		{"emptyAddr",
			NewMsgPlaceBid(0, sdk.AccAddress{}, c("token", 10)),
			false},
		{"negativeAmount",
			NewMsgPlaceBid(0, addr, sdk.Coin{Denom: "token", Amount: sdk.NewInt(-10)}),
			false},
		{"zeroAmount",
			NewMsgPlaceBid(0, addr, c("token", 0)),
			true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.expectPass {
				require.NoError(t, tc.msg.ValidateBasic())
			} else {
				require.Error(t, tc.msg.ValidateBasic())
			}
		})
	}
}
