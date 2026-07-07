<div align="center">

# Guard VPN

**A lightweight VPN built with Go that provides secure communication between clients and servers using encrypted tunnels, virtual network interfaces, and UDP-based packet forwarding.**

[![Go](https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=white)](https://go.dev/)
[![Linux](https://img.shields.io/badge/Linux-FCC624?logo=linux&logoColor=black)](https://kernel.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

</div>

---

# Overview

Guard VPN is a lightweight Virtual Private Network (VPN) solution developed in Go that enables secure communication between distributed systems over the internet.

It creates encrypted tunnels between clients and servers using UDP sockets and virtual network interfaces (TUN devices), allowing private network communication as if the devices were connected to the same local network.

Guard VPN is designed for learning and experimentation with networking concepts such as:

- VPN architecture
- TUN interfaces
- UDP communication
- Packet forwarding
- Routings
- Secure remote connectivity

---

# Features

| Feature | Description |
|---|---|
| 🌐 **Virtual Network Interface** | Uses Linux TUN devices to create virtual private networks |
| ⚡ **UDP-Based Communication** | Fast packet transmission using UDP sockets |
| 🔄 **Packet Forwarding** | Transfers IP packets between clients and the VPN server |
| 🖥️ **Cross-Network Connectivity** | Connect devices across different networks securely |
| 📚 **Educational Project** | Helps understand VPN internals and Linux networking |

---

# Architecture

```text
+-------------+       UDP Tunnel       +-------------+
| VPN Client  | <--------------------> | VPN Server  |
|   tun0      |                        |   tun0      |
+-------------+                        +-------------+
       |                                      |
       +---------- Virtual Private Network ---+
```

---

# Tech Stack

| Layer | Technology |
|---|---|
| **Language** | Go (Golang) |
| **Operating System** | Linux |
| **Networking** | UDP Sockets |
| **Virtual Interfaces** | TUN/TAP |
| **System Programming** | Linux Networking APIs |


---

# Getting Started

## Prerequisites

- Go 1.22+
- Linux Operating System
- Git

---

# Installation

## 1. Clone Repository

```bash
git clone https://github.com/johngithiyon/Guard.git
```

```bash
cd Guard
```

---

## 2. Install Dependencies

```bash
go mod tidy
```

---

## 3. Build the Server

```bash
go build -o guard-server ./cmd/server
```

---

## 4. Build the Client

```bash
go build -o guard-client ./cmd/client
```

---

# Running the VPN

## Start the Server

```bash
sudo ./guard-server
```

---

## Start the Client

```bash
sudo ./guard-client
```


---

# Learning Objectives

This project demonstrates:

- VPN implementation from scratch
- Linux networking internals
- TUN/TAP interfaces
- UDP socket programming
- Routing and packet forwarding
- Secure communication over public networks

---

# Future Improvements

- [ ] Encryption using AES or ChaCha20
- [ ] Authentication mechanism
- [ ] Cross-platform support

---

# License

This project is licensed under the MIT License — see the [LICENSE](LICENSE) file for details.

---

<div align="center">

**Built with ❤️ using Go and Linux Networking**

</div>