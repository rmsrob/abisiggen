# abisiggen 🛠️⚡
[![Sanity checks](https://github.com/rmsrob/abisiggen/actions/workflows/ci.yml/badge.svg)](https://github.com/rmsrob/abisiggen/actions/workflows/ci.yml)
[![License: AGPL-3.0-only](https://img.shields.io/badge/License-AGPL--3.0--only-blue)](https://www.gnu.org/licenses/agpl-3.0)


> Smart Contract Interface Generator for Ethereum Developers

## Overview
Automatically generates type-safe Go bindings, event topics, method selectors, and documentation from smart contract ABIs. Designed for Foundry/Go Ethereum projects to prevent signature mismatches and accelerate development.

Key features:
- 🎯 Event topic hashes with full signatures
- 🔑 Method selectors with ABI encoding
- 📘 CLI-friendly lowercase method names
- 📖 Auto-generated contract documentation
- 🔄 Initialization helpers for Go code

## Features
- **Event Topic Generation** - Precise keccak256 hashes for all events
- **Method Selectors** - Pre-computed 4-byte signatures with full function signatures
- **Multi-Contract Support** - Processes all sub folders that contain an ABI
- **Documentation Generation** - Creates `_help.md` with event/view details
- **Type Safety** - Generates Go structs with proper types
- **Version Control** - Maintains consistency across contract versions

## Install dependencies
- abigen [Go Ethereum](https://geth.ethereum.org/docs/tools/abigen)
- jq [jqlang](https://jqlang.org/)
- cast [Foundry](https://book.getfoundry.sh/cast/)
- Go 1.20+ [Go](https://go.dev/dl/)

```bash
brew install jq go-ethereum foundry
```

## Usage
- Copy `abisiggen.sh` to your project.
- Copy your contract ABIs to `pkg/contracts/[ContractName]/[Contract].abi`

> [!TIP]  
> You can edit the `CONTRACT_SOURCES_DIR`:  
```bash
# Find all directories in the contract sources directory
# you can edit this to hardcode the contracts you want to process
PROJECT_ROOT="$(dirname "$(dirname "$(realpath "$0")")")"
CONTRACT_SOURCES_DIR="$PROJECT_ROOT/pkg/contracts"
```

> Follow this structure, or ensure `PROJECT_ROOT` is set correctly:
```txt
.sh/
├── abisiggen.sh        # Main generator script
└── pkg/contracts/
    └── [ContractName]/
        └── [Contract].abi      # Contract ABI
```

```bash
chmod +x sh/abisiggen.sh
```

```bash
./sh/abisiggen.sh
```

## Output
```txt
.sh/
├── abisiggen.sh        # Main generator script
└── pkg/contracts/
    └── [ContractName]/
        ├── [Contract].abi      # Contract ABI
        ├── [Contract].go       # Go bindings (abigen)
        ├── [Contract]_extra.go # Generated topics/methods
        └── [Contract]_help.md  # Generated documentation
```

> Example console:  
> See the generated output for Sablier contract in [example](/example/)

```console
sh .sh/abisiggen.sh
Processing SablierV2LockupLinear with abigen
✓ Generated SablierV2LockupLinear bindings
✓ Generated SablierV2LockupLinear_extra.go
Created SablierV2LockupLinear_help.md
Processing SablierV2LockupDynamic with abigen
✓ Generated SablierV2LockupDynamic bindings
✓ Generated SablierV2LockupDynamic_extra.go
Created SablierV2LockupDynamic_help.md
```
