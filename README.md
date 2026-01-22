# Digital Asset Management System

A blockchain-based **Digital Asset Management System** built on **Hyperledger Fabric** , featuring a Go backend service and a web-based management interface.

This project demonstrates a typical **enterprise blockchain application architecture**, combining smart contracts (chaincode), off-chain business services, and a frontend management dashboard.

## Overview

This project implements a permissioned blockchain solution for managing digital assets with strong guarantees of **data integrity, traceability, and access control**.

Hyperledger Fabric is used as the underlying blockchain framework to handle on-chain asset state and transaction validation. A Go-based backend service interacts with the Fabric network via the Fabric SDK, while a web frontend provides an interface for asset management and visualization.

The system follows a clear separation of concerns:

- **Blockchain layer**: smart contracts and ledger state
- **Service layer**: business logic and blockchain interaction
- **Presentation layer**: web-based management interface


## Features

- Permissioned blockchain network based on Hyperledger Fabric
- Smart contract (chaincode) driven digital asset management
- Go backend service with Fabric SDK integration
- Web-based asset management dashboard
- Separation of on-chain and off-chain data
- Auditable and traceable asset operations


## Project Structure

```test
blockchain-testwork/
├── main.go # Application entry point
├── go.mod # Go module definition
├── go.sum # Go dependency checksums
├── config.yaml # Application and blockchain configuration
├── Data.sql # Database initialization script
│
├── chaincode/ # Hyperledger Fabric smart contracts
├── sdkInit/ # Fabric SDK initialization logic
├── ca/ # Certificate Authority related code
├── fixtures/ # Fabric network configuration and crypto materials
│
├── web/ # Web frontend pages
├── static/ # Frontend static resources (JS, CSS, plugins)
│
├── .idea/ # IDE configuration files
└── .DS_Store # macOS system file (can be ignored)
```

## Backend Service

The backend service is implemented in **Go** and acts as the core application
layer of the system. It is responsible for:

- Exposing RESTful APIs
- Loading application and blockchain configuration
- Initializing and interacting with the Fabric SDK
- Submitting and querying blockchain transactions
- Handling off-chain business logic and database operations

### Key Files

- `main.go`  
  Application entry point and service bootstrap.

- `config.yaml`  
  Configuration file containing Fabric network parameters and application
  settings.

- `Data.sql`  
  Database schema used for off-chain data storage.


## Blockchain Layer (Hyperledger Fabric)

The blockchain layer is built on **Hyperledger Fabric**, a permissioned,
enterprise-grade blockchain framework.

### Chaincode (`chaincode/`)

- Defines digital asset data models
- Implements asset creation, query, and update logic
- Enforces business rules on-chain

### Fabric SDK (`sdkInit/`)

- Initializes Fabric network connections
- Loads certificates and MSP configuration
- Handles chaincode invocation and state queries

### Certificate Authority (`ca/`)

- Manages identities for organizations, peers, and users
- Issues and validates X.509 certificates


## Fabric Network Configuration

The `fixtures/` directory contains all resources required to bootstrap and
run a Hyperledger Fabric network, including:

- Channel and organization configuration
- Cryptographic material generation rules
- Docker Compose configuration
- Channel artifacts (genesis block, channel transactions)
- MSP and TLS certificates

This directory is primarily intended for **development and testing
environments**.


## Web Frontend

The web frontend provides a management dashboard that allows users to:

- View and query digital assets
- Inspect asset-related transactions
- Interact with backend APIs

The frontend communicates only with the backend service and does not directly
access the blockchain network.


