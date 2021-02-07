package codesamples

import (
	"testing"

	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/balance"
	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasp/packages/testutil"
	"github.com/stretchr/testify/require"
)

func Test_CreateChain_chainCreatorSpecified(t *testing.T) {
	env := solo.New(t, false, false)

	// Create address with dummy tokens in it.
	const initialWalletFunds = testutil.RequestFundsAmount
	chainOriginatorsWalletKeyPair := env.NewSignatureSchemeWithFunds()
	require.NotNil(t, chainOriginatorsWalletKeyPair)
	// Chain originators must have at least 1 iota token to create a chain
	require.GreaterOrEqual(t, initialWalletFunds, 1)

	// Wallet addresses
	chainOriginatorsWalletAddress := chainOriginatorsWalletKeyPair.Address()
	require.NotNil(t, chainOriginatorsWalletAddress)

	// Wallet balance in value tangle
	env.AssertAddressBalance(chainOriginatorsWalletAddress, balance.ColorIOTA, initialWalletFunds)

	// Create a chain where chainCreatorsWalletKeyPair is the owner.
	chain := env.NewChain(chainOriginatorsWalletKeyPair, "myChain")
	require.NotNil(t, chain)
	require.NotEqual(t, chain.ChainColor, balance.ColorIOTA)

	// AgentID of the wallet (also, chain.OriginatorAgentID)
	chainOriginatorsAgentID := coretypes.NewAgentIDFromAddress(chainOriginatorsWalletAddress)

	// IMPORTANT: When a chain is created, 1 IOTA is sent from the originator's account in the value tangle their account in the chain
	chain.AssertAccountBalance(chainOriginatorsAgentID, balance.ColorIOTA, 1)
}

// Sample of how to create chain without specifying a chainOriginator.
// A dummy chain originator is created in the background (by NewChain).
func Test_CreateChain_NoChainCreatorSpecified(t *testing.T) {
	env := solo.New(t, false, false)

	// Create a chain where chainCreatorsWalletKeyPair is the owner.
	chain := env.NewChain(nil, "myChain")
	require.NotNil(t, chain)
	require.NotEqual(t, chain.ChainColor, balance.ColorIOTA)

	// IMPORTANT: When a chain is created, 1 IOTA is sent from the originator's account in the value tangle their account in the chain
	chain.AssertAccountBalance(chain.OriginatorAgentID, balance.ColorIOTA, 1)
}
