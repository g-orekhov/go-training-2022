package main

import (
	"context"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	api "github.com/test_server/internal/contractsAPI"
	"github.com/test_server/internal/helpers/etherium_auth"
)

func main() {
	godotenv.Load("./cmd/deploy_contracts/.env")
	// address of etherum env
	client, err := ethclient.Dial(os.Getenv("NODE_HTTP"))
	if err != nil {
		println(" - Main: 1 - ")
		panic(err)
	}
	defer client.Close()

	// create auth and transaction package for deploying smart contract
	auth := etherium_auth.GetAccountAuth(context.Background(), client, os.Getenv("NODE_USER"))

	//deploying smart contract
	deployedContractAddress, _, _, err := api.DeployApi(auth, client) //api is redirected from api directory from our contract go file
	if err != nil {
		println(" - Main: 2 - ")
		panic(err)
	}

	fmt.Println(deployedContractAddress.Hex()) // print deployed contract address
}
