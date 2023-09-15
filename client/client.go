package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/grpc_tutorial/client/account"
	pb "github.com/grpc_tutorial/proto/account"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:8080", "the address to connect to")
)

func main() {

	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewAccountServiceClient(conn)

	account_made, err := account.NewAccount(c, 1995)
	if err != nil {
		log.Fatalf("%v", err)
	}

	fmt.Println("New account", account_made)

	account_get, err := account.GetAccount(c, account_made.GetAccountId())
	if err != nil {
		log.Fatalf("%v", err)
	}

	fmt.Println("Get account", account_get)

	account_deposit_changed, err := account.Deposit(
		c,
		&pb.Account{
			AccountId: account_made.GetAccountId(),
			Balance:   500,
		})
	if err != nil {
		log.Fatalf("%v", err)
	}

	fmt.Println("Deposit", account_deposit_changed)

	account_withdraw_changed, err := account.Withdraw(
		c,
		&pb.Account{
			AccountId: account_made.GetAccountId(),
			Balance:   300,
		})
	if err != nil {
		log.Fatalf("%v", err)
	}

	fmt.Println("Withdraw", account_withdraw_changed)

}
