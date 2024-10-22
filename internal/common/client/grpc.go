package client

import (
	"google.golang.org/grpc"

	"kwai/api/proto/user"
)

func NewUserClient() (client user.UserClient, close func() error, err error) {
	conn, err := grpc.Dial(":8085", grpc.WithInsecure())
	if err != nil {
		return nil, nil, err
	}

	client = user.NewUserClient(conn)
	close = conn.Close
	return
}
