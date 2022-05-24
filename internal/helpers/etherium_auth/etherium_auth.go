package etherium_auth

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetAccountAuth(ctx context.Context, client *ethclient.Client, accountAddress string) *bind.TransactOpts {

	privateKey, err := crypto.HexToECDSA(accountAddress)
	if err != nil {
		println(" - getAccountAuth: 1 - ")
		panic(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		println(" - getAccountAuth: 2 - ")
		panic("invalid key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	//fetch the last use nonce of account
	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		panic(err)
	}
	fmt.Println("nounce=", nonce)
	chainID, err := client.ChainID(ctx)
	if err != nil {
		println(" - getAccountAuth: 3 - ")
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		println(" - getAccountAuth: 4 - ")
		panic(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = big.NewInt(1000000)

	return auth
}
