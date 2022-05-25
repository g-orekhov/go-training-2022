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
		return err
	}
	defer client.Close()

	// open contract
	conn, err := api.NewApi(common.HexToAddress(os.Getenv("CONTRACT_ADRESS")), client)
	if err != nil {
		return err
	}

	// create auth transaction
	auth, err := etherium_auth.GetAccountAuth(r.ctx, client, os.Getenv("NODE_USER"))
	if err != nil {
		return err
	}

	// create user
	if _, err := conn.UsersCreate(auth, user.Name); err != nil {
		return err
	}

	// var created chan *api.ApiUserCreated
	// //TODO: Добавить контекст с таймаутом в WatchOpts{}
	// ctx, cancel := context.WithTimeout(r.ctx, time.Millisecond*20000) // 10sec timeout
	// defer cancel()
	// event, err := conn.WatchUserCreated(
	// 	&bind.WatchOpts{
	// 		Start:   nil,
	// 		Context: ctx,
	// 	},
	// 	created,
	// )
	// if err != nil {
	// 	return err
	// }
	// defer event.Unsubscribe()

	// //TODO: допилить обработку события.
	// /*
	// 	Тут в принципе не понятно, что если произойдёт другое событие?
	// 	Это ведь подписка на журнал в целом, а не привязка к транзакции.
	// 	То есть мы должны возвращать не только ID но и например наш адресс,
	// 	затем зависнуть на итераторе сравнивая адрес, и ждать наше изменение.

	// 	Чем больше смотрю сюда, тем больше вопросов.
	// 	Лучше просто так не делать и привязывать всё к адресам.
	// 	Или надо вешать ожидание сообытия в другую горутину, и по результату создавать отдельный
	// 	обьект в бд с последним айди для клента.
	// 	Иначе как быть в случае таймаутов с нодой или клиентом.
	// 	Клиент должен иметь возможность спросить как там дела.
	// 	Может TxIndex как то может помочь. Надо разбираться.
	// */
	// //INFO: HTTP не поддерживает подписки, надо юзать вебсокет.

	// // Это всё не работает :(
	// select {
	// case apiUserCreatedPointer := <-created:
	// 	user.Id = int64(apiUserCreatedPointer.Id)
	// 	return nil
	// case err = <-event.Err():
	// 	return err
	// case <-ctx.Done():
	// 	return ctx.Err()
	// }
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
	auth, err := etherium_auth.GetAccountAuth(r.ctx, client, os.Getenv("NODE_USER"))
	if err != nil {
		return err
	}

	// create user
	if _, err := conn.UsersDelete(auth, uint64(id)); err != nil {
		return err
	}
	//TODO: почему не работает require? в методе UsersDelete
	//TODO: разобраться как вернуть знаначение из не view функции.
	return nil
}
