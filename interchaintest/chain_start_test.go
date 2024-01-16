package interchaintest

import (
	"testing"

	"github.com/sedaprotocol/seda-chain/interchaintest/conformance"
	"github.com/strangelove-ventures/interchaintest/v8"
	"github.com/strangelove-ventures/interchaintest/v8/chain/cosmos"
	"github.com/stretchr/testify/require"
)

func TestChainStart(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	t.Parallel()

	numOfValidators := 2
	numOfFullNodes := 0

	chains := CreateChains(t, numOfValidators, numOfFullNodes)
	ic, ctx, _, _ := BuildAll(t, chains)

	chain := chains[0].(*cosmos.CosmosChain)

	const userFunds = int64(10_000_000_000)
	users := interchaintest.GetAndFundTestUsers(t, ctx, t.Name(), userFunds, chain)
	chainUser := users[0]

	conformance.ConformanceCosmWasm(t, ctx, chain, chainUser)

	require.NotNil(t, ic)
	require.NotNil(t, ctx)

	t.Cleanup(func() {
		_ = ic.Close()
	})
}