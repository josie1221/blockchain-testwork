# Digital Asset Management System

A blockchain-based **Digital Asset Management System** built on
**Hyperledger Fabric**, featuring a Go backend service and a web-based
management interface.

This project demonstrates a typical **enterprise blockchain application
architecture**, combining smart contracts (chaincode), off-chain business
services, and a frontend management dashboard.

## Overview

This project implements a permissioned blockchain solution for managing
digital assets with strong guarantees of **data integrity, traceability,
and access control**.

Hyperledger Fabric is used as the underlying blockchain framework to handle
on-chain asset state and transaction validation. A Go-based backend service
interacts with the Fabric network via the Fabric SDK, while a web frontend
provides an interface for asset management and visualization.

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


## Architecture

