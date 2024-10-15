package main

import (
	"PowerLedgerGo/infrastructure/contract/artifacts"
	"context"
	"encoding/hex"
	"fmt"
	"github.com/FISCO-BCOS/go-sdk/v3/client"
	"github.com/ethereum/go-ethereum/common"
	"log"
)

func main() {
	privateKey, _ := hex.DecodeString("6a8c9ee93251719eb7f3542fb22b44c362f0cc072551a81872ca39c88f7f449f")

	config := &client.Config{
		IsSMCrypto:  false,
		GroupID:     "group0",
		PrivateKey:  privateKey,
		Host:        "127.0.0.1",
		Port:        20200,
		TLSCaFile:   "./infrastructure/config/ca.crt",
		TLSKeyFile:  "./infrastructure/config/sdk.key",
		TLSCertFile: "./infrastructure/config/sdk.crt"}

	contractInstance, err := client.DialContext(context.Background(), config)
	if err != nil {
		log.Fatal(err)
	}

	// 部署HomomorphicEncryption合约
	HomomorphicEncryptionContract, receipt, _, err := artifacts.DeployHomomorphicEncryption(contractInstance.GetTransactOpts(), contractInstance)
	if err != nil {
		log.Fatalln("部署HomomorphicEncryption合约失败: " + err.Error())
	}
	fmt.Println("HomomorphicEncryption contract address: ", HomomorphicEncryptionContract.Hex())
	fmt.Println("deploy HomomorphicEncryption contract transaction hash: ", receipt.TransactionHash)
	fmt.Println()

	AdvancedVirtualPowerPlantDAOContract, receipt, _, err := artifacts.DeployAdvancedVirtualPowerPlantDAO(
		contractInstance.GetTransactOpts(),
		contractInstance,
		HomomorphicEncryptionContract,
		common.HexToAddress("0x8f4847e1C1AB77E026AdA8d9467F0fB5A74b0e50"))
	if err != nil {
		log.Fatalln("部署AdvancedVirtualPowerPlantDAO合约失败: " + err.Error())
	}
	fmt.Println("AdvancedVirtualPowerPlantDAO contract address: ", AdvancedVirtualPowerPlantDAOContract.Hex())
	fmt.Println("deploy AdvancedVirtualPowerPlantDAO contract transaction hash: ", receipt.TransactionHash)
}
