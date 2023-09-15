package account

import (
	"context"
	"time"

	pb "github.com/grpc_tutorial/proto/account"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func NewAccount(c pb.AccountServiceClient, balance int32) (*pb.Account, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	return c.NewAccount(ctx, &wrapperspb.Int32Value{Value: balance})
}

func GetAccount(c pb.AccountServiceClient, id string) (*pb.Account, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	return c.GetAccount(ctx, &wrapperspb.StringValue{Value: id})
}

func DeleteAccount(c pb.AccountServiceClient, id string) (*pb.Account, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	return c.DeleteAccount(ctx, &wrapperspb.StringValue{Value: id})
}

func Deposit(c pb.AccountServiceClient, account *pb.Account) (*pb.Account, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	return c.Deposit(ctx, account)
}

func Withdraw(c pb.AccountServiceClient, account *pb.Account) (*pb.Account, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	return c.Withdraw(ctx, account)
}
