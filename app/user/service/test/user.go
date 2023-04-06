package main

import (
	"context"
	v12 "github.com/869413421/micro-chat/api/user/service/v1"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	transgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	"log"
)

func main() {
	callGRPC()
}

func callGRPC() {
	conn, err := transgrpc.DialInsecure(
		context.Background(),
		transgrpc.WithEndpoint("127.0.0.1:9000"),
		transgrpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := v12.NewUserClient(conn)
	reply, err := client.CreateUser(context.Background(), &v12.CreateUserRequest{
		Email:    "13528685024@163.com",
		Password: "admin",
		Name:     "清水泥沙",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("[grpc] SayHello %+v\n", reply.Id)

	if errors.IsBadRequest(err) {
		log.Printf("[grpc] SayHello error is invalid argument: %v\n", err)
	}
}
