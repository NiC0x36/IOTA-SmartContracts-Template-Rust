package libtest

import (
	"testing"

	"github.com/drand/drand/fs"
	"github.com/stretchr/testify/require"
)

// Ensure contract folder and wasm file exist before the test
func mustGetContractWasmFilePath(t *testing.T, contractName string) string {
	contractWasmFilePath := "../pkg/" + contractName + "_bg.wasm"
	exists, err := fs.Exists(contractWasmFilePath)
	require.NoError(t, err)
	require.True(t, exists)

	return contractWasmFilePath
}
