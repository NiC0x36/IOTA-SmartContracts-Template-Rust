package codesamples

import (
	"testing"

	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/balance"
	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/kv/codec"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasp/packages/testutil"
	"github.com/iotaledger/wasp/packages/vm/core/accounts"
	"github.com/stretchr/testify/require"
)

/// This is a sample of how to send tokens from the value tangle to a chain, keeping the ownership of the tokens on that chain
func Test_SendTokensToChain(t *testing.T) {
	env := solo.New(t, false, false)

	const initialWalletFunds = testutil.RequestFundsAmount // used to fund address in NewSignatureSchemeWithFunds
	const transferValueIotas = 1000
	require.GreaterOrEqual(t, initialWalletFunds, transferValueIotas)

	// Generates key pairs for sender and receiver wallets and provides both with dummy funds.
	senderWalletKeyPair := env.NewSignatureSchemeWithFunds()
	senderWalletAddress := senderWalletKeyPair.Address()
	senderAgentID := coretypes.NewAgentIDFromAddress(senderWalletAddress)
	require.NotNil(t, senderWalletKeyPair)
	require.NotNil(t, senderWalletAddress)
	require.NotNil(t, senderAgentID)

	// Wallet balance in value tangle -> Before transfer
	env.AssertAddressBalance(senderWalletAddress, balance.ColorIOTA, initialWalletFunds)

	// Generate a dummy chain NOT beloging to sender
	chain := env.NewChain(nil, "myChain")

	// Wallet balance in chain -> Before transfer
	chain.AssertAccountBalance(senderAgentID, balance.ColorIOTA, 0)

	// Transfer from value tangle to the chain
	transferRequest := solo.NewCallParams(accounts.Name, accounts.FuncDeposit).WithTransfer(balance.ColorIOTA, transferValueIotas)
	chain.PostRequest(transferRequest, senderWalletKeyPair)

	// Wallet balances -> After transfer to chain
	env.AssertAddressBalance(senderWalletAddress, balance.ColorIOTA, initialWalletFunds-transferValueIotas-1) // -1 ???
	chain.AssertAccountBalance(senderAgentID, balance.ColorIOTA, transferValueIotas+1)                        // +1 ???
}

func Test_SendAndReceiveTokens(t *testing.T) {
	env := solo.New(t, false, false)

	const initialSenderWalletFunds = testutil.RequestFundsAmount // used to fund address in NewSignatureSchemeWithFunds
	const transferValueIotas = 1000
	require.GreaterOrEqual(t, initialSenderWalletFunds, transferValueIotas)

	// Generates key pairs for sender wallet and provides it with dummy funds.
	senderWalletKeyPair := env.NewSignatureSchemeWithFunds()
	senderWalletAddress := senderWalletKeyPair.Address()
	senderWalletAgentID := coretypes.NewAgentIDFromAddress(senderWalletAddress)
	senderIotaBalance := env.GetAddressBalance(senderWalletAddress, balance.ColorIOTA)
	require.NotNil(t, senderWalletKeyPair)
	require.NotNil(t, senderWalletAddress)
	env.AssertAddressBalance(senderWalletAddress, balance.ColorIOTA, initialSenderWalletFunds)

	// Generates key pairs for sender wallet.
	receivedWalletKeyPair := env.NewSignatureScheme()
	receiverWalletAddress := receivedWalletKeyPair.Address()
	receiverWalletAgentID := coretypes.NewAgentIDFromAddress(receiverWalletAddress)
	receiverIotaBalance := env.GetAddressBalance(receiverWalletAddress, balance.ColorIOTA)
	require.NotNil(t, receivedWalletKeyPair)
	require.NotNil(t, receiverWalletAddress)
	require.NotNil(t, receiverWalletAgentID)
	env.AssertAddressBalance(receiverWalletAddress, balance.ColorIOTA, 0)

	// Generate a dummy chain NEITHER belonging to sender NOR receiver
	chain := env.NewChain(nil, "myChain")

	// Transfer within the chain
	params := codec.MakeDict(map[string]interface{}{
		accounts.ParamAgentID: codec.EncodeAgentID(receiverWalletAgentID),
	})
	transferRequest := solo.NewCallParamsFromDic(accounts.Name, accounts.FuncDeposit, params).WithTransfer(balance.ColorIOTA, transferValueIotas)
	chain.PostRequest(transferRequest, senderWalletKeyPair)

	// Wallet balances -> After transfer
	env.AssertAddressBalance(senderWalletAddress, balance.ColorIOTA, initialSenderWalletFunds-transferValueIotas-1) // -1 ???
	chain.AssertAccountBalance(senderWalletAgentID, balance.ColorIOTA, +1)                                          // +1 ???

	env.AssertAddressBalance(receiverWalletAddress, balance.ColorIOTA, 0)
	chain.AssertAccountBalance(receiverWalletAgentID, balance.ColorIOTA, 0+transferValueIotas)

	t.Logf("Sender wallet address after transfer: %s - Balance: %d IOTA(s)", senderWalletAddress, senderIotaBalance)
	t.Logf("Receiver wallet address after transfer: %s - Balance: %d IOTA(s)", receiverWalletAddress, receiverIotaBalance)
}
