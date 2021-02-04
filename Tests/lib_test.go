package libtest

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/user/project/Tests/testutils"

	"github.com/iotaledger/wasp/packages/solo"
)

func TestLib(t *testing.T) {
	// Contract name - Defined in Cargo.toml > package > name
	const contractName = "my_iota_sc"
	contractWasmFilePath := testutils.MustGetContractWasmFilePath(t, contractName)

	// Name of the SC function to be requested - Defined in lib.rs > add_call > my_sc_request
	functionName := "my_sc_request"

	env := solo.New(t, false, false)
	chainName := contractName + "Chain"
	chain := env.NewChain(nil, chainName)

	// Uploads wasm of SC and deploys it into chain
	err := chain.DeployWasmContract(nil, contractName, contractWasmFilePath)
	require.NoError(t, err)

	req := solo.NewCallParams(contractName, functionName)
	_, err = chain.PostRequest(req, nil)
	require.NoError(t, err)
}
