package user

import (
	"context"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	api "github.com/test_server/internal/contractsAPI"
	eth_methods "github.com/test_server/internal/helpers/etherium"
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

func (r *repository) FindAll() (users []User, err error) {
	users = make([]User, 0, UsersCount)
	ctx, cancel := context.WithTimeout(r.ctx, time.Millisecond*10000) // 10sec timeout
	defer cancel()
	defer eth_methods.TrasferPanicToError(&err)
	// connect to node
	client := eth_methods.ConnectToNode(os.Getenv("NODE_HTTP"))
	defer client.Close()
	// get contract api
	contractApi := eth_methods.GetContractApi(client, os.Getenv("CONTRACT_ADRESS"))
	// get user
	ret, err := contractApi.UsersGetAll(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}
	// transform returned objects to #User structure
	for i := 0; i < len(ret); i++ {
		users = append(users, User{Id: int64(ret[i].Id), Name: ret[i].Name})
	}
	return users, nil
}

func (r *repository) FindOne(id int64) (user *User, err error) {
	ctx, cancel := context.WithTimeout(r.ctx, time.Millisecond*10000) // 10sec timeout
	defer cancel()
	defer eth_methods.TrasferPanicToError(&err)
	// connect to node
	client := eth_methods.ConnectToNode(os.Getenv("NODE_HTTP"))
	defer client.Close()
	// get contract api
	contractApi := eth_methods.GetContractApi(client, os.Getenv("CONTRACT_ADRESS"))
	// get user
	ret, err := contractApi.UsersGetOne(&bind.CallOpts{Context: ctx}, uint64(id))
	if err != nil {
		return nil, err
	}
	user = &User{Id: int64(ret.Id), Name: ret.Name}
	return user, err
}

func (r *repository) Create(user *User) (err error) {
	ctx, cancel := context.WithTimeout(r.ctx, time.Millisecond*10000) // 10sec timeout
	defer cancel()
	defer eth_methods.TrasferPanicToError(&err)
	// connect to node
	client := eth_methods.ConnectToNode(os.Getenv("NODE_HTTP"))
	defer client.Close()
	// get contract api
	contractApi := eth_methods.GetContractApi(client, os.Getenv("CONTRACT_ADRESS"))
	// create auth transaction
	transaction := eth_methods.GetAuthorizedTransaction(ctx, client, os.Getenv("NODE_USER"))

	//Subscribe to #UserCreated event
	created := make(chan *api.ApiUserCreated)
	event, err := contractApi.WatchUserCreated(
		&bind.WatchOpts{
			Start:   nil,
			Context: ctx,
		},
		created,
	)
	if err != nil {
		return err
	}
	defer event.Unsubscribe()

	// create user
	if _, err := contractApi.UsersCreate(transaction, user.Name); err != nil {
		return err
	}

	//TODO: допилить обработку события.
	/*
		Тут в принципе не понятно, что если произойдёт другое событие?
		Это ведь подписка на журнал в целом, а не привязка к транзакции.
		То есть мы должны возвращать не только ID но и например наш адресс,
		затем зависнуть на итераторе сравнивая адрес, и ждать именно наше изменение.

		Чем больше смотрю сюда, тем больше вопросов.
		Лучше просто так не делать и привязывать всё к адресам.
		Или надо вешать ожидание сообытия в другую горутину, и по результату создавать отдельный
		обьект в бд с последним айди для клента.
		Иначе как быть в случае таймаутов с нодой или клиентом.
		Клиент должен иметь возможность спросить как там дела.
		Надо разбираться.
	*/
	//INFO: HTTP не поддерживает подписки, надо юзать вебсокет.

	select {
	case apiUserCreatedPointer := <-created:
		user.Id = int64(apiUserCreatedPointer.Id)
		return nil
	case err = <-event.Err():
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (r *repository) Delete(id int64) (err error) {
	ctx, cancel := context.WithTimeout(r.ctx, time.Millisecond*10000) // 10sec timeout
	defer cancel()
	defer eth_methods.TrasferPanicToError(&err)
	// connect to node
	client := eth_methods.ConnectToNode(os.Getenv("NODE_HTTP"))
	defer client.Close()
	// get contract api
	contractApi := eth_methods.GetContractApi(client, os.Getenv("CONTRACT_ADRESS"))
	// create auth transaction
	transaction := eth_methods.GetAuthorizedTransaction(ctx, client, os.Getenv("NODE_USER"))
	// delete user
	if _, err := contractApi.UsersDelete(transaction, uint64(id)); err != nil {
		return err
	}
	return nil
}
