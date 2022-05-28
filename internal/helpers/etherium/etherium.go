package etherium_methods

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	api "github.com/test_server/internal/contractsAPI"
)

func GetAuthorizedTransaction(ctx context.Context, client *ethclient.Client, accountAddress string) *bind.TransactOpts {
	//TODO: Got this method from google. Make refactoring.
	privateKey, err := crypto.HexToECDSA(accountAddress)
	if err != nil {
		panic(err.Error())
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("invalid key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	//fetch the last use nonce of account
	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("nounce=", nonce)
	chainID, err := client.ChainID(ctx)
	if err != nil {
		panic(err.Error())
	}

	transaction, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		panic(err.Error())
	}
	transaction.Nonce = big.NewInt(int64(nonce))
	transaction.Value = big.NewInt(0)      // in wei
	transaction.GasLimit = uint64(3000000) // in units
	transaction.GasPrice = big.NewInt(1000000)

	return transaction
}

func TrasferPanicToError(err *error) {
	if recovery := recover(); recovery != nil {
		if errRecovery, ok := recovery.(error); ok {
			(*err) = errRecovery
		}
		if strRecovery, ok := recovery.(string); ok {
			(*err) = errors.New(strRecovery)
		}
		(*err) = errors.New(fmt.Sprint("Panic in TrasferPanicToError:", recovery))
	}
}

func ConnectToNode(nodeAdress string) *ethclient.Client {
	client, err := ethclient.Dial(nodeAdress)
	if err != nil {
		panic(err.Error())
	}
	return client
}

func GetContractApi(client *ethclient.Client, contractAdress string) *api.Api {
	conn, err := api.NewApi(common.HexToAddress(contractAdress), client)
	if err != nil {
		panic(err.Error())
	}
	return conn
}
