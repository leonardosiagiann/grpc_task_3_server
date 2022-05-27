package account

import (
	"context"
	"log"
	"net"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"

	account "grpc_server/proto/account"
)

func dial() func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(1024 * 1024)

	server := grpc.NewServer()

	account.RegisterDepositServiceServer(server, &DepositoService{})

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}

}

func TestDepositService_Deposit(t *testing.T) {
	test := []struct {
		name   string
		amount float32
		res    *account.DepositResponse
		errMsg string
	}{
		{
			"Invalid deposit amount (negative)",
			-1.1,
			&account.DepositResponse{Ok: false},
			"deposit value must bigger than 0",
		},
		{
			"Valid Request",
			1.1,
			&account.DepositResponse{Ok: true},
			"",
		},
	}

	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dial()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := account.NewDepositServiceClient(conn)

	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			request := &account.DepositRequest{Amount: tc.amount}

			response, err := client.Deposit(ctx, request)

			if response.GetOk() != tc.res.GetOk() {
				t.Error("response expected", tc.res.GetOk(), "received", response.GetOk())
			}

			if err != nil {
				if er, _ := status.FromError(err); er.Message() != tc.errMsg {
					t.Error("error code:expected", tc.errMsg, "received", err)
				}
			}
		})
	}

}

func TestDepositService_GetDeposit(t *testing.T) {
	test := struct {
		name string
		res  *account.GetDepositResponse
		err  error
	}{
		"Test Get Total Deposit",
		&account.GetDepositResponse{TotalDeposit: 1.1},
		nil,
	}

	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dial()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := account.NewDepositServiceClient(conn)

	t.Run(test.name, func(t *testing.T) {
		request := &account.GetDepositRequest{}

		response, err := client.GetDeposit(ctx, request)

		if response.TotalDeposit != test.res.TotalDeposit {
			t.Error("response expected", test.res, "received", response)
		}

		if err != nil {
			t.Error("response expected", test.err, "received", err)
		}
	})

}
