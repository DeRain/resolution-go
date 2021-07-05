![Test](https://github.com/unstoppabledomains/resolution-go/workflows/Test/badge.svg?branch=master)
![Lint](https://github.com/unstoppabledomains/resolution-go/workflows/Lint/badge.svg?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/unstoppabledomains/resolution-go)](https://goreportcard.com/report/github.com/unstoppabledomains/resolution-go)
[![GoDoc](https://godoc.org/github.com/unstoppabledomains/resolution-go?status.svg)](https://pkg.go.dev/github.com/unstoppabledomains/resolution-go)
[![Unstoppable Domains Documentation](https://img.shields.io/badge/docs-unstoppabledomains.com-blue)](https://docs.unstoppabledomains.com/)
[![Get help on Discord](https://img.shields.io/badge/Get%20help%20on-Discord-blueviolet)](https://discord.gg/b6ZVxSZ9Hn)

# resolution-go

resolution-go is a library for interacting with blockchain domain names. It can be used to retrieve [payment addresses](https://unstoppabledomains.com/features#Add-Crypto-Addresses), IPFS hashes for [decentralized websites](https://unstoppabledomains.com/features#Build-Website), DNS records and other [records types](https://docs.unstoppabledomains.com/domain-registry-essentials/records-reference)

resolution-go is primarily built and maintained by [Unstoppable Domains](https://unstoppabledomains.com/).

resolution-go decentralized domains across two zones:

- Unstoppable Name Service (UNS)
  - `.crypto`
  - `.wallet`
  - `.coin`
  - `.bitcoin`
  - `.x`
  - `.888`
  - `.nft`
  - `.dao`
  - `.blockchain`
- Zilliqa Name Service (ZNS)
  - `.zil`

# Usage

```go
package main

import (
  "fmt"
  "github.com/Zilliqa/gozilliqa-sdk/provider"
  "github.com/ethereum/go-ethereum/ethclient"
  "github.com/unstoppabledomains/resolution-go"
  "github.com/unstoppabledomains/resolution-go/namingservice"
)

func main() {
  // Resolve .crypto
  uns, _ := resolution.NewUnsBuilder().Build()
  ethAddress, _ := uns.Addr("brad.crypto", "ETH")
  fmt.Println("ETH address for brad.crypto is", ethAddress)

  // Resolve .zil
  zns, _ := resolution.NewZnsBuilder().Build()
  btcAddress, _ := zns.Addr("brad.zil", "BTC")
  fmt.Println("BTC address for brad.zil is", btcAddress)

  // Detect domain naming service
  namingServices := map[string]resolution.NamingService{namingservice.UNS: uns, namingservice.ZNS: zns}
  domainToDetect := "ryan.crypto"
  namingServiceName, _ := resolution.DetectNamingService(domainToDetect)
  if namingServices[namingServiceName] != nil {
    resolvedAddress, _ := namingServices[namingServiceName].Addr(domainToDetect, "ETH")
    fmt.Println("ETH address for", domainToDetect, "is", resolvedAddress)
  }

  // Set custom Ethereum endpoint for UNS service
  ethContractBackend, _ := ethclient.Dial("https://mainnet.infura.io/v3/c5da69dfac9c4d9d96dd232580d4124e")
  unsWithCustomBackend, _ := resolution.NewUnsBuilder().SetContractBackend(ethContractBackend).Build()
  allUnsRecords, _ := unsWithCustomBackend.AllRecords("beresnev.crypto")
  fmt.Println("Records for beresnev.crypto", allUnsRecords)

  // Set custom Zilliqa endpoint for ZNS service
  zilliqaProvider := provider.NewProvider("https://api.zilliqa.com")
  znsWithCustomProvider, _ := resolution.NewZnsBuilder().SetProvider(zilliqaProvider).Build()
  allZnsRecords, _ := znsWithCustomProvider.AllRecords("brad.zil")
  fmt.Println("Records for brad.zil", allZnsRecords)
}
```

# Network support

Library supports Ethereum mainnet and Zilliqa mainnet only.

# Contributions

Contributions to this library are more than welcome. The easiest way to contribute is through GitHub issues and pull requests.

# Free advertising for integrated apps

Once your app has a working Unstoppable Domains integration, [register it here](https://unstoppabledomains.com/app-submission). Registered apps appear on the Unstoppable Domains [homepage](https://unstoppabledomains.com/) and [Applications](https://unstoppabledomains.com/apps) page — putting your app in front of tens of thousands of potential customers per day.

Also, every week we select a newly-integrated app to feature in the Unstoppable Update newsletter. This newsletter is delivered to straight into the inbox of ~100,000 crypto fanatics — all of whom could be new customers to grow your business.

# Get help

[Join our discord community](https://discord.com/invite/b6ZVxSZ9Hn) and ask questions.
