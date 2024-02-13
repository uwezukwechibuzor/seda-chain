package keeper_test

import (
	"testing"

	"cosmossdk.io/math"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktestutil "github.com/cosmos/cosmos-sdk/x/bank/testutil"
	sdkstakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/sedaprotocol/seda-chain/x/vesting/types"
)

func TestClawback(t *testing.T) {
	f := initFixture(t)
	f.bankKeeper.SetSendEnabled(f.Context(), "aseda", true)
	err := banktestutil.FundAccount(f.Context(), f.bankKeeper, funderAddr, sdk.NewCoins(sdk.NewInt64Coin(bondDenom, 100000)))
	require.NoError(t, err)

	_, valAddrs, valPks := createValidators(t, f, []int64{5, 5, 5})
	require.NoError(t, err)

	testCases := []struct {
		testName                string
		funder                  sdk.AccAddress
		recipient               sdk.AccAddress
		vestingTime             int64
		timeUntilClawback       int64
		originalVesting         sdk.Coin
		delegation              sdk.Coin
		delegation2             sdk.Coin
		undelegation            sdk.Coin
		undelegation2           sdk.Coin
		expClawedUnbonded       sdk.Coins
		expClawedUnbonding      sdk.Coins
		expClawedBonded         sdk.Coins
		slashingFraction        math.LegacyDec
		recipientFinalSpendable sdk.Coins
	}{
		{
			testName:                "clawback from unbonded",
			funder:                  sdk.MustAccAddressFromBech32("seda1gujynygp0tkwzfpt0g7dv4829jwyk8f0yhp88d"),
			recipient:               testAddrs[0],
			vestingTime:             100,
			timeUntilClawback:       30,
			originalVesting:         sdk.NewInt64Coin(bondDenom, 10000),
			delegation:              sdk.NewInt64Coin(bondDenom, 0),
			undelegation:            sdk.NewInt64Coin(bondDenom, 0),
			expClawedUnbonded:       sdk.NewCoins(sdk.NewInt64Coin(bondDenom, 7000)),
			expClawedUnbonding:      zeroCoins,
			expClawedBonded:         zeroCoins,
			slashingFraction:        math.LegacyZeroDec(),
			recipientFinalSpendable: sdk.NewCoins(sdk.NewInt64Coin(bondDenom, 3000)),
		},
		{
			testName:                "clawback from unbonded and bonded",
			funder:                  sdk.MustAccAddressFromBech32("seda1gujynygp0tkwzfpt0g7dv4829jwyk8f0yhp88d"),
			recipient:               testAddrs[1],
			vestingTime:             100,
			timeUntilClawback:       30,
			originalVesting:         sdk.NewInt64Coin(bondDenom, 10000),
			delegation:              sdk.NewInt64Coin(bondDenom, 5000),
			delegation2:             sdk.NewInt64Coin(bondDenom, 0),
			undelegation:            sdk.NewInt64Coin(bondDenom, 0),
			undelegation2:           sdk.NewInt64Coin(bondDenom, 0),
			expClawedUnbonded:       sdk.NewCoins(sdk.NewInt64Coin(bondDenom, 5000)),
			expClawedUnbonding:      zeroCoins,
			expClawedBonded:         sdk.NewCoins(sdk.NewInt64Coin(bondDenom, 2000)),
			slashingFraction:        math.LegacyZeroDec(),
			recipientFinalSpendable: sdk.NewCoins(),
		},
		{
			testName:                "clawback from bonded",
			funder:                  sdk.MustAccAddressFromBech32("seda1gujynygp0tkwzfpt0g7dv4829jwyk8f0yhp88d"),
			recipient:               testAddrs[2],
			vestingTime:             100,
			timeUntilClawback:       60,
			originalVesting:         sdk.NewInt64Coin(bondDenom, 27500),
			delegation:              sdk.NewInt64Coin(bondDenom, 27500),
			delegation2:             sdk.NewInt64Coin(bondDenom, 0),
			undelegation:            sdk.NewInt64Coin(bondDenom, 0),
			undelegation2:           sdk.NewInt64Coin(bondDenom, 0),
			expClawedUnbonded:       zeroCoins,
			expClawedUnbonding:      zeroCoins,
			expClawedBonded:         sdk.NewCoins(sdk.NewInt64Coin(bondDenom, 11000)),
			slashingFraction:        math.LegacyZeroDec(),
			recipientFinalSpendable: sdk.NewCoins(),
		},
		{
			testName:                "clawback from unbonding",
			funder:                  sdk.MustAccAddressFromBech32("seda1gujynygp0tkwzfpt0g7dv4829jwyk8f0yhp88d"),
			recipient:               testAddrs[3],
			vestingTime:             50000,
			timeUntilClawback:       30000,
			originalVesting:         sdk.NewInt64Coin(bondDenom, 27500),
			delegation:              sdk.NewInt64Coin(bondDenom, 27500),
			delegation2:             sdk.NewInt64Coin(bondDenom, 0),
			undelegation:            sdk.NewInt64Coin(bondDenom, 27500),
			undelegation2:           sdk.NewInt64Coin(bondDenom, 0),
			expClawedUnbonded:       zeroCoins,
			expClawedUnbonding:      sdk.NewCoins(sdk.NewInt64Coin(bondDenom, 11000)),
			expClawedBonded:         zeroCoins,
			slashingFraction:        math.LegacyZeroDec(),
			recipientFinalSpendable: sdk.NewCoins(),
		},
		{
			testName:                "clawback from unbonded, unbonding, and bonded",
			funder:                  sdk.MustAccAddressFromBech32("seda1gujynygp0tkwzfpt0g7dv4829jwyk8f0yhp88d"),
			recipient:               testAddrs[4],
			vestingTime:             750000,
			timeUntilClawback:       600000,
			originalVesting:         sdk.NewInt64Coin(bondDenom, 13000),
			delegation:              sdk.NewInt64Coin(bondDenom, 10000),
			delegation2:             sdk.NewInt64Coin(bondDenom, 2000),
			undelegation:            sdk.NewInt64Coin(bondDenom, 400),
			undelegation2:           sdk.NewInt64Coin(bondDenom, 100),
			expClawedUnbonded:       sdk.NewCoins(sdk.NewInt64Coin(bondDenom, 1000)),
			expClawedUnbonding:      sdk.NewCoins(sdk.NewInt64Coin(bondDenom, 500)),
			expClawedBonded:         sdk.NewCoins(sdk.NewInt64Coin(bondDenom, 1100)),
			slashingFraction:        math.LegacyZeroDec(),
			recipientFinalSpendable: sdk.NewCoins(),
		},
		{
			testName:                "clawback from unbonded and bonded with slashing",
			funder:                  sdk.MustAccAddressFromBech32("seda1gujynygp0tkwzfpt0g7dv4829jwyk8f0yhp88d"),
			recipient:               testAddrs[6],
			vestingTime:             100,
			timeUntilClawback:       30,
			originalVesting:         sdk.NewInt64Coin(bondDenom, 10000),
			delegation:              sdk.NewInt64Coin(bondDenom, 5000),
			delegation2:             sdk.NewInt64Coin(bondDenom, 0),
			undelegation:            sdk.NewInt64Coin(bondDenom, 0),
			undelegation2:           sdk.NewInt64Coin(bondDenom, 0),
			expClawedUnbonded:       sdk.NewCoins(sdk.NewInt64Coin(bondDenom, 5000)),
			expClawedUnbonding:      zeroCoins,
			expClawedBonded:         sdk.NewCoins(sdk.NewInt64Coin(bondDenom, 2000)),
			slashingFraction:        math.LegacyNewDecWithPrec(5, 2),
			recipientFinalSpendable: sdk.NewCoins(),
		},
		{
			testName:                "clawback from unbonded, unbonding, and bonded with slashing",
			funder:                  sdk.MustAccAddressFromBech32("seda1gujynygp0tkwzfpt0g7dv4829jwyk8f0yhp88d"),
			recipient:               testAddrs[5],
			vestingTime:             750000,
			timeUntilClawback:       600000,
			originalVesting:         sdk.NewInt64Coin(bondDenom, 13000),
			delegation:              sdk.NewInt64Coin(bondDenom, 10000),
			delegation2:             sdk.NewInt64Coin(bondDenom, 2000),
			undelegation:            sdk.NewInt64Coin(bondDenom, 400),
			undelegation2:           sdk.NewInt64Coin(bondDenom, 100),
			expClawedUnbonded:       sdk.NewCoins(sdk.NewInt64Coin(bondDenom, 1000)),
			expClawedUnbonding:      sdk.NewCoins(sdk.NewInt64Coin(bondDenom, 180)),
			expClawedBonded:         sdk.NewCoins(sdk.NewInt64Coin(bondDenom, 1420)),
			slashingFraction:        math.LegacyNewDecWithPrec(8, 1), // 0.8
			recipientFinalSpendable: sdk.NewCoins(),
		},
	}
	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.testName, func(t *testing.T) {
			f.AddBlock()

			// 1. create clawback continuous vesting account
			createVestingMsg := &types.MsgCreateVestingAccount{
				FromAddress: tc.funder.String(),
				ToAddress:   tc.recipient.String(),
				Amount:      sdk.NewCoins(tc.originalVesting),
				EndTime:     f.Context().BlockTime().Unix() + tc.vestingTime,
			}
			_, err = f.RunMsg(createVestingMsg)
			require.NoError(t, err)

			// 2. delegate
			if tc.delegation.IsPositive() {
				delegateMsg := &sdkstakingtypes.MsgDelegate{
					DelegatorAddress: tc.recipient.String(),
					ValidatorAddress: valAddrs[0].String(),
					Amount:           tc.delegation,
				}
				_, err = f.RunMsg(delegateMsg)
				require.NoError(t, err)

				if tc.delegation2.IsPositive() {
					delegateMsg := &sdkstakingtypes.MsgDelegate{
						DelegatorAddress: tc.recipient.String(),
						ValidatorAddress: valAddrs[1].String(),
						Amount:           tc.delegation2,
					}
					_, err = f.RunMsg(delegateMsg)
					require.NoError(t, err)
				}
			}

			// 3. initiate unbonding after some time
			if tc.undelegation.IsPositive() {
				undelegateMsg := &sdkstakingtypes.MsgUndelegate{
					DelegatorAddress: tc.recipient.String(),
					ValidatorAddress: valAddrs[0].String(),
					Amount:           tc.undelegation,
				}
				_, err = f.RunMsg(undelegateMsg)
				require.NoError(t, err)

				if tc.undelegation2.IsPositive() {
					undelegateMsg := &sdkstakingtypes.MsgUndelegate{
						DelegatorAddress: tc.recipient.String(),
						ValidatorAddress: valAddrs[1].String(),
						Amount:           tc.undelegation2,
					}
					_, err = f.RunMsg(undelegateMsg)
					require.NoError(t, err)
				}
			}

			// possible slashing
			if tc.slashingFraction.IsPositive() {
				_, err = f.stakingKeeper.Slash(f.Context(), sdk.ConsAddress(valPks[0].Address()), f.Context().BlockHeight()-1, 5, tc.slashingFraction)
				require.NoError(t, err)
			}

			_, err = f.stakingKeeper.EndBlocker(f.Context())
			require.NoError(t, err)
			f.AddBlock()

			// 4. clawback after some time
			f.AddTime(tc.timeUntilClawback)

			clawbackMsg := &types.MsgClawback{
				FunderAddress:  tc.funder.String(),
				AccountAddress: tc.recipient.String(),
			}
			res, err := f.RunMsg(clawbackMsg)
			require.NoError(t, err)

			result := types.MsgClawbackResponse{}
			err = f.cdc.Unmarshal(res.Value, &result)
			require.NoError(t, err)

			require.Equal(t, tc.expClawedUnbonded, result.ClawedUnbonded)
			require.Equal(t, tc.expClawedUnbonding, result.ClawedUnbonding)
			require.Equal(t, tc.expClawedBonded, result.ClawedBonded)

			//
			recipientSpendable := f.bankKeeper.SpendableCoins(f.Context(), tc.recipient)
			require.Equal(t, tc.recipientFinalSpendable, recipientSpendable)
		})
	}
}
