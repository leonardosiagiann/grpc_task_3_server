# grpc_task_3_server

## Usage

### Run Server

1. Clone [Server](https://github.com/leonardosiagiann/grpc_task_3_server) and [Client](https://github.com/leonardosiagiann/grpc_task_3_client)

2. Run Makefile (Client & Server)
```bash
make grpc
make proto
```

3. Run server
```bash
go run ./server/main.go
```

4. Server ready to retrieve data from client


### Unit Test
Run unit test for server
```bash
go test -v ./controllers/account/

=== RUN   TestDepositServiceClient_Deposit
=== RUN   TestDepositServiceClient_Deposit/Invalid_Request_With_Invalid_Amount
=== RUN   TestDepositServiceClient_Deposit/Valid_Request_With_Valid_Amount
--- PASS: TestDepositServiceClient_Deposit (0.00s)
    --- PASS: TestDepositServiceClient_Deposit/Invalid_Request_With_Invalid_Amount (0.00s)
    --- PASS: TestDepositServiceClient_Deposit/Valid_Request_With_Valid_Amount (0.00s)
=== RUN   TestDepositServiceClient_GetDeposit
=== RUN   TestDepositServiceClient_GetDeposit/Valid_Test_Get_Deposit
--- PASS: TestDepositServiceClient_GetDeposit (0.00s)
    --- PASS: TestDepositServiceClient_GetDeposit/Valid_Test_Get_Deposit (0.00s)
PASS
ok  	grpc_client/controllers/account	(cached)
```
