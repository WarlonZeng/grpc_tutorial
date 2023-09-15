package routes

import (
	"google.golang.org/grpc"

	"github.com/grpc_tutorial/api/services/account"
	pb "github.com/grpc_tutorial/proto/account"
)

func SetupRoutes(s *grpc.Server) {

	pb.RegisterAccountServiceServer(s, &account.AccountService{})

}
