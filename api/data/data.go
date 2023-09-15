package data

import (
	"errors"
	"fmt"
	"sync"

	pb "github.com/grpc_tutorial/proto/account"
)

type SimpleDatabase struct {
	AccountData map[string]*pb.Account
	mu          sync.Mutex
}

func NewSimpleDatabase() *SimpleDatabase {
	return &SimpleDatabase{
		AccountData: map[string]*pb.Account{},
		mu:          sync.Mutex{},
	}
}

// Account CRUD
func (db *SimpleDatabase) NewAccount(id string, account *pb.Account) {
	db.mu.Lock()
	db.AccountData[id] = account
	db.mu.Unlock()
}

func (db *SimpleDatabase) GetAccount(id string) (*pb.Account, error) {
	db.mu.Lock()
	val, err := db.AccountData[id]
	db.mu.Unlock()

	fmt.Println(id)
	fmt.Println(val)
	fmt.Println(db.AccountData)

	if !err {
		return nil, errors.New("account not found")
	}

	return val, nil
}

func (db *SimpleDatabase) DeleteAccount(id string) error {
	_, err := db.GetAccount(id)
	if err != nil {
		return err
	}

	delete(db.AccountData, id)

	return nil
}

func (db *SimpleDatabase) UpdateAccount(id string, value int32) (*pb.Account, error) {
	val, err := db.GetAccount(id)
	if err != nil {
		return nil, err
	}

	db.mu.Lock()
	val.Balance += value
	db.mu.Unlock()

	return val, nil
}

var SimpleDB = NewSimpleDatabase()
