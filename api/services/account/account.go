package account

import (
	"log"

	"github.com/google/uuid"
	"github.com/grpc_tutorial/api/data"
	pb "github.com/grpc_tutorial/proto/account"
	"golang.org/x/net/context"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type AccountService struct {
	pb.UnimplementedAccountServiceServer
}

// handler vs. service
// service is server-free, no ctx context.Context
func (s *AccountService) NewAccount(ctx context.Context, amount *wrapperspb.Int32Value) (*pb.Account, error) {
	log.Println("New account: ", amount.GetValue())

	res := &pb.Account{
		AccountId: uuid.New().String(),
		Balance:   amount.GetValue(),
	}

	data.SimpleDB.NewAccount(res.GetAccountId(), res)

	return res, nil
}

func (s *AccountService) GetAccount(ctx context.Context, id *wrapperspb.StringValue) (*pb.Account, error) {
	log.Println("Get account: ", id.GetValue())

	return data.SimpleDB.GetAccount(id.GetValue())
}

func (s *AccountService) DeleteAccount(ctx context.Context, id *wrapperspb.StringValue) (*pb.Account, error) {
	log.Println("Delete account: ", id.GetValue())

	return nil, data.SimpleDB.DeleteAccount(id.GetValue())
}

func (s *AccountService) Deposit(ctx context.Context, account *pb.Account) (*pb.Account, error) {
	log.Println("Deposit: ", account.GetAccountId(), account.GetBalance())

	return data.SimpleDB.UpdateAccount(account.GetAccountId(), account.GetBalance())
}

func (s *AccountService) Withdraw(ctx context.Context, account *pb.Account) (*pb.Account, error) {
	log.Println("Withdraw: ", account.GetAccountId(), account.GetBalance())

	return data.SimpleDB.UpdateAccount(account.GetAccountId(), -account.GetBalance())
}
