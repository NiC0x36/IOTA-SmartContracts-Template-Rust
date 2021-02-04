#### Why is Go required?
Go is used by [Solo](https://github.com/iotaledger/wasp/tree/develop/packages/solo) to simulate the behavior of Wasp nodes. Unit tests for smart contracts are written in Go so Solo is acessible. The only other option would be to deploy the SCs under development to Wasp nodes, without the chance to test it locally.
