# MiniEVM - Lightweight Ethereum Virtual Machine Implementation

[![Go Report Card](https://goreportcard.com/badge/github.com/blockfuselabs/evm)](https://goreportcard.com/report/github.com/blockfuselabs/evm)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![GoDoc](https://godoc.org/github.com/blockfuselabs/evm?status.svg)](https://godoc.org/github.com/blockfuselabs/evm)

## 📖 Overview

MiniEVM is a lightweight implementation of the Ethereum Virtual Machine (EVM) written in Go. It provides a minimal, yet functional environment for executing EVM bytecode, making it ideal for educational purposes, testing, and research. While not intended for production use, it implements core EVM features including stack operations, memory management, and gas metering.

## 📗 Table of Contents

- [📖 Overview](#-overview)
- [✨ Features](#-features)
- [🏗️ Architecture](#️-architecture)
- [🚀 Installation](#-installation)
- [💻 Usage](#-usage)
- [📂 Project Structure](#-project-structure)
- [🔍 Implementation Details](#-implementation-details)
- [⛽ Gas Costs](#-gas-costs)
- [🤝 Contributing](#-contributing)
- [🛣️ Roadmap](#️-roadmap)
- [📄 License](#-license)

## ✨ Features

### Current Features
* Basic EVM execution environment
* Stack-based operation handling (push/pop)
* Memory management
* Storage operations
* Gas metering and management
* Basic arithmetic operations (ADD, SUB, MUL, DIV)
* Program counter management
* State management and execution flow control

### Planned Features
* Complete implementation of all EVM opcodes
* Contract creation and deployment
* Event logging system
* Memory expansion and gas calculation
* Advanced arithmetic operations (ADDMOD, MULMOD, EXP)
* Bitwise operations
* Jump operations with proper validation
* Exception handling mechanisms
* Contract calling mechanisms (CALL, DELEGATECALL, STATICCALL)
* State reverting capabilities
* SHA3 hashing operations
* Block information access
* Extended testing framework

## 🏗️ Architecture

The MiniEVM is built with a modular architecture consisting of several key components:

### Core Components

1. **State Management (`state.go`)**
   * Maintains execution context
   * Handles program counter
   * Manages gas consumption
   * Controls execution flow

2. **Stack (`types.go`)**
   * Implementation: `Stack` struct
   * Maximum depth: 1024 items
   * Basic operations: Push, Pop
   * Stack overflow protection

3. **Memory (`types.go`)**
   * Dynamic memory model
   * 32-byte word alignment
   * Automatic expansion
   * Bounds checking

4. **Storage (`types.go`)**
   * Key-value storage model
   * Persistent state storage
   * 32-byte keys and values

5. **Operation Codes (`opcodes.go`)**
   * Complete EVM opcode definitions
   * Grouped by functionality
   * Includes gas costs

## 🚀 Installation

```bash
# Clone the repository
git clone https://github.com/blockfuselabs/evm.git

# Change to project directory
cd evm

# Install dependencies
go mod download

# Build the project
go build ./cmd/cli
```

## 💻 Usage

### Basic Example

```go
package main

import (
    "fmt"
    "github.com/blockfuselabs/evm/types"
)

func main() {
    // Create new EVM instance with initial parameters
    evm := types.NewEvmState(
        "0x1234...", // sender address
        []byte{0x60, 0x03, 0x60, 0x02, 0x01, 0x00}, // bytecode
        1000,        // gas limit
        100,        // value
        []uint8{},  // calldata
    )

    // Execute the bytecode
    evm.Run()

    // Print results
    fmt.Printf("Remaining Gas: %d\nStack: %v\n", evm.Gas, evm.Stack.Data)
}
```

### Advanced Usage

```go
// Example of executing custom bytecode
bytecode := []byte{
    types.PUSH1, 0x03, // Push 3 onto stack
    types.PUSH1, 0x02, // Push 2 onto stack
    types.ADD,         // Add top two stack items
    types.STOP,        // Stop execution
}

evm := types.NewEvmState("sender", bytecode, 1000, 0, nil)
evm.Run()
```

## 📂 Project Structure

```
evm/
├── bin/                    # Compiled binaries
├── cmd/
│   └── cli/
│       └── main.go        # CLI entry point
└── types/
    ├── types.go           # Core type definitions
    ├── methods.go         # EVM operations implementation
    ├── opcodes.go         # Opcode definitions
    └── state.go           # State management
```

## 🔍 Implementation Details

### Stack Operations

The stack implementation follows EVM specifications:
* Maximum depth of 1024 items
* Each item is 32 bytes
* Automatic overflow checking
* Panic on stack overflow

```go
func (s *Stack) Push(data byte) {
    if len(s.Data) >= int(MaximumDepth) {
        panic("stack overflow")
    }
    s.Data = append(s.Data, data)
}
```

### Memory Management

Memory is managed in 32-byte words:
* Dynamic expansion
* Gas cost calculation for expansion
* Automatic alignment
* Bounds checking on access

### Gas Calculation

Gas costs are implemented according to the Ethereum Yellow Paper:
* Basic operations: 3 gas
* Memory expansion: dynamic cost
* Storage operations: high cost
* Complex operations: variable cost

## 🤝 Contributing

We welcome contributions! Please follow these steps:

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to your branch
5. Create a Pull Request

Please ensure your code:
* Follows Go best practices
* Includes tests
* Is properly documented
* Passes all existing tests

## 🛣️ Roadmap

### Q1 2025
- [ ] Complete implementation of arithmetic operations
- [ ] Add comprehensive testing suite
- [ ] Implement memory expansion with proper gas calculation

### Q2 2025
- [ ] Add contract creation capabilities
- [ ] Implement call operations
- [ ] Add event logging system

### Q3 2025
- [ ] Implement remaining EVM opcodes
- [ ] Add advanced debugging features
- [ ] Performance optimizations

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.