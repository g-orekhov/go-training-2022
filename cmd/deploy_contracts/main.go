package main

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	api "github.com/test_server/internal/contractsAPI"
	eth_methods "github.com/test_server/internal/helpers/etherium"
)

func main() {
	godotenv.Load("./cmd/deploy_contracts/.env")
	// address of etherum env
	client := eth_methods.ConnectToNode(os.Getenv("NODE_HTTP"))
	defer client.Close()

	// create transaction and transaction package for deploying smart contract
	transaction := eth_methods.GetAuthorizedTransaction(context.Background(), client, os.Getenv("NODE_USER"))

	//deploying smart contract
	deployedContractAddress, _, _, err := api.DeployApi(transaction, client) //api is redirected from api directory from our contract go file
	if err != nil {
		panic(err)
	}

	fmt.Println(deployedContractAddress.Hex()) // print deployed contract address
}
