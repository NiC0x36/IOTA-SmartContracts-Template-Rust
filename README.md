## IOTA smart contracts - Template for development in Rust

A simple template used to start developing your own smart contracts for ISCP (IOTA Smart Contract Protocol) in Rust and write unit tests in Go. 

#### Simple structure prepared to start with development right away
This is how the templated file structure looks like:

![View of the template on VSCode](VSCode_Rust_Template_View.png)

---

### Requirements
- [Rust](https://www.rust-lang.org/tools/install)
- [Wasm-pack](https://rustwasm.github.io/wasm-pack/installer/)
- [Go](https://golang.org/dl/) - [(Why Go)](WhyGo.md)
- Gcc (or equivalent for Windows [(TDM-GCC)](https://jmeubank.github.io/tdm-gcc/))
- [Visual Studio Code](https://code.visualstudio.com/Download) (VSCode)
  - [Rust extension](https://marketplace.visualstudio.com/items?itemName=rust-lang.rust)
  - [Better TOML](https://marketplace.visualstudio.com/items?itemName=bungcip.better-toml) *Optional nice to have 
  - [Go Extension](https://marketplace.visualstudio.com/items?itemName=golang.Go)

### Set code up!
- Use this template repository to create your own.
- Clone your git repository with:
```
git clone --recurse-submodules <your_git_repository>
```
- Open your git repository on VSCode

### Import how-to's:
- [Compile](Compile-SmartContract.md) a smart contract
- [Unit test and debug](UnitTest-and-debug-SmartContract.md) a smart contract
