package user

import (
	"context"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	api "github.com/test_server/internal/contractsAPI"
	"github.com/test_server/internal/helpers/etherium_auth"
)

type Repository interface {
	FindAll() ([]User, error)
	FindOne(int64) (*User, error)
	Create(*User) error
	Delete(int64) error
}

const UsersCount int64 = 10

type repository struct {
	// Some internal data
	ctx context.Context
}

func NewRepository(ctx context.Context) Repository {
	return &repository{ctx: ctx}
}

func (r *repository) FindAll() ([]User, error) {
	users := make([]User, 0, UsersCount)

	// connect to node
	client, err := ethclient.Dial(os.Getenv("NODE_HTTP"))
	if err != nil {
		fmt.Println("- Dial error")
		return nil, err
	}
	defer client.Close()

	// open contract
	conn, err := api.NewApi(common.HexToAddress(os.Getenv("CONTRACT_ADRESS")), client)
	if err != nil {
		fmt.Println("- Connection to contract error")
		return nil, err
	}

	// create user
	ret, err := conn.UsersGetAll(&bind.CallOpts{})
	if err != nil {
		fmt.Println("- Get all")
		return nil, err
	}

	for i := 0; i < len(ret); i++ {
		users = append(users, User{Id: int64(ret[i].Id), Name: ret[i].Name})
	}

	return users, nil
}

func (r *repository) FindOne(id int64) (*User, error) {
	// connect to node
	client, err := ethclient.Dial(os.Getenv("NODE_HTTP"))
	if err != nil {
		fmt.Println("- Dial error")
		return nil, err
	}
	defer client.Close()

	// open contract
	conn, err := api.NewApi(common.HexToAddress(os.Getenv("CONTRACT_ADRESS")), client)
	if err != nil {
		fmt.Println("- Connection to contract error")
		return nil, err
	}

	// create user
	ret, err := conn.UsersGetOne(&bind.CallOpts{}, uint64(id))
	if err != nil {
		return nil, err
	}

	return &User{Id: int64(ret.Id), Name: ret.Name}, err
}

func (r *repository) Create(user *User) error {
	// connect to node
	client, err := ethclient.Dial(os.Getenv("NODE_HTTP"))
	if err != nil {
		fmt.Println("- Dial error")
		return err
	}
	defer client.Close()

	// open contract
	conn, err := api.NewApi(common.HexToAddress(os.Getenv("CONTRACT_ADRESS")), client)
	if err != nil {
		fmt.Println("- Connection to contract error")
		return err
	}

	// create auth
	auth := etherium_auth.GetAccountAuth(r.ctx, client, os.Getenv("NODE_USER"))

	// create user
	if _, err := conn.UsersCreate(auth, user.Name); err != nil {
		return err
	}
	//TODO: разобраться как вернуть знаначение из не view функции.
	return nil
}

func (r *repository) Delete(id int64) error {
	// connect to node
	client, err := ethclient.Dial(os.Getenv("NODE_HTTP"))
	if err != nil {
		fmt.Println("- Dial error")
		return err
	}
	defer client.Close()

	// open contract
	conn, err := api.NewApi(common.HexToAddress(os.Getenv("CONTRACT_ADRESS")), client)
	if err != nil {
		fmt.Println("- Connection to contract error")
		return err
	}

	// create auth
	auth := etherium_auth.GetAccountAuth(r.ctx, client, os.Getenv("NODE_USER"))

	// create user
	if _, err := conn.UsersDelete(auth, uint64(id)); err != nil {
		return err
	}
	//TODO: почему не работает require? в методе UsersDelete
	//TODO: разобраться как вернуть знаначение из не view функции.
	return nil
}
