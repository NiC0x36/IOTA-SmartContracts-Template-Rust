package codesamples

import (
	"testing"

	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/balance"
	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasp/packages/testutil"
	"github.com/iotaledger/wasp/packages/vm/core/accounts"
	"github.com/stretchr/testify/require"
)

/// This is a sample of how to send tokens from the value tangle to a chain, keeping the ownership of the tokens on that chain
func Test_SendTokensToChain(t *testing.T) {
	env := solo.New(t, false, false)

	// Generates key pairs for sender and receiver wallets and provides both with dummy funds.
	// The amount is defined in Wasp (constant testutil.RequestFundsAmount) and WaspConn plug-in (constant utxodb.RequestFundsAmount)
	senderWalletKeyPair := env.NewSignatureSchemeWithFunds()
	require.NotNil(t, senderWalletKeyPair)

	// Wallet addresses
	senderWalletAddress := senderWalletKeyPair.Address()
	require.NotNil(t, senderWalletAddress)

	// Wallet balance in value tangle -> Before transfer
	senderIotaBalance := env.GetAddressBalance(senderWalletAddress, balance.ColorIOTA) // Balance in value tangle
	env.AssertAddressBalance(senderWalletAddress, balance.ColorIOTA, testutil.RequestFundsAmount)

	// Wallet AgentID
	senderAgentID := coretypes.NewAgentIDFromAddress(senderWalletAddress)
	require.NotNil(t, senderAgentID)

	// Wallet balance in chain -> Before transfer
	chain := env.NewChain(senderWalletKeyPair, "myChain")
	chain.AssertAccountBalance(senderAgentID, balance.ColorIOTA, 0)

	// Transfer from value tangle to the chain
	const transferValueIotas = 1000
	transferRequest := solo.NewCallParams(accounts.Name, accounts.FuncDeposit).WithTransfer(balance.ColorIOTA, transferValueIotas)
	chain.PostRequest(transferRequest, senderWalletKeyPair)

	// Wallet balances -> After transfer to chain
	senderIotaBalance = env.GetAddressBalance(senderWalletAddress, balance.ColorIOTA)                                  // Balance in value tangle
	env.AssertAddressBalance(senderWalletAddress, balance.ColorIOTA, testutil.RequestFundsAmount-transferValueIotas-1) // -1 for technical reasons
	senderAgentIotaBalance := chain.GetAccountBalance(senderAgentID).Balance(balance.ColorIOTA)
	chain.AssertAccountBalance(senderAgentID, balance.ColorIOTA, transferValueIotas+1)

	t.Logf("%s", chain.DumpAccounts())
	t.Logf("Value tangle wallet %s after transfer - Balance: %d IOTA(s)", senderWalletAddress, senderIotaBalance)
	t.Logf("Chain agent %s after transfer - Balance: %d IOTA(s)", senderAgentID, senderAgentIotaBalance)
}

// TODO Finish writing this
func Test_SendAndReceiveTokens(t *testing.T) {
	env := solo.New(t, false, false)

	// Generates key pairs for sender and receiver wallets and provides both with dummy funds.
	// The amount is defined in Wasp (constant testutil.RequestFundsAmount) and WaspConn plug-in (constant utxodb.RequestFundsAmount)
	senderWalletKeyPair := env.NewSignatureSchemeWithFunds()
	receivedWalletKeyPair := env.NewSignatureSchemeWithFunds()
	require.NotNil(t, senderWalletKeyPair)
	require.NotNil(t, receivedWalletKeyPair)

	// Wallet addresses
	senderWalletAddress := senderWalletKeyPair.Address()
	receiverWalletAddress := receivedWalletKeyPair.Address()
	require.NotNil(t, senderWalletAddress)
	require.NotNil(t, receiverWalletAddress)

	// Wallet balances -> Before transfer
	senderIotaBalance := env.GetAddressBalance(senderWalletAddress, balance.ColorIOTA)
	receiverIotaBalance := env.GetAddressBalance(receiverWalletAddress, balance.ColorIOTA)
	require.Equal(t, testutil.RequestFundsAmount, senderIotaBalance)
	require.Equal(t, testutil.RequestFundsAmount, receiverIotaBalance)
	env.AssertAddressBalance(senderWalletAddress, balance.ColorIOTA, testutil.RequestFundsAmount)
	env.AssertAddressBalance(receiverWalletAddress, balance.ColorIOTA, testutil.RequestFundsAmount)

	// Creating chain to transact in
	chain := env.NewChain(nil, "myChain")

	// Transfer within the chain
	const transferValueIotas = 1000
	transferRequest := solo.NewCallParams(accounts.Name, accounts.FuncDeposit).WithTransfer(balance.ColorIOTA, transferValueIotas)
	chain.PostRequest(transferRequest, senderWalletKeyPair)

	// Wallet balances -> After transfer
	senderIotaBalance = env.GetAddressBalance(senderWalletAddress, balance.ColorIOTA)
	receiverIotaBalance = env.GetAddressBalance(receiverWalletAddress, balance.ColorIOTA)
	require.Equal(t, testutil.RequestFundsAmount, senderIotaBalance)
	require.Equal(t, testutil.RequestFundsAmount, receiverIotaBalance)
	env.AssertAddressBalance(senderWalletAddress, balance.ColorIOTA, testutil.RequestFundsAmount-transferValueIotas)
	env.AssertAddressBalance(receiverWalletAddress, balance.ColorIOTA, testutil.RequestFundsAmount+transferValueIotas)

	t.Logf("Sender wallet address after transfer: %s - Balance: %d IOTA(s)", senderWalletAddress, senderIotaBalance)
	t.Logf("Receiver wallet address after transfer: %s - Balance: %d IOTA(s)", receiverWalletAddress, receiverIotaBalance)
}
