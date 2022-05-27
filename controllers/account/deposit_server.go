package account

import (
	"context"
	"errors"
	"log"

	proto "grpc_server/proto/account"
)

type Deposito struct {
	DepositAmmount float32
}

var deposit Deposito

type DepositoService struct {
	proto.UnimplementedDepositServiceServer
}

func NewDepositoService() *DepositoService {
	return &DepositoService{}
}

func (d *DepositoService) Deposit(ctx context.Context, in *proto.DepositRequest) (*proto.DepositResponse, error) {
	if in.GetAmount() <= 0 {
		return &proto.DepositResponse{Ok: false}, errors.New("deposit value must bigger than 0")
	}
	log.Println("Received data from client:", in)
	deposit.DepositAmmount = deposit.DepositAmmount + in.GetAmount()
	ok := true
	response := proto.DepositResponse{
		Ok: ok,
	}
	return &response, nil
}

func (d *DepositoService) GetDeposit(ctx context.Context, in *proto.GetDepositRequest) (*proto.GetDepositResponse, error) {
	response := proto.GetDepositResponse{
		TotalDeposit: deposit.DepositAmmount,
	}
	return &response, nil
}
