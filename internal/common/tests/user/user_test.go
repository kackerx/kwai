package user

import (
	"context"
	"fmt"
	"testing"

	pb "kwai/api/proto/user"
	"kwai/internal/common/client"
)

func TestUser(t *testing.T) {
	c, closeFn, err := client.NewUserClient()
	defer func() {
		closeFn()
	}()
	if err != nil {
		panic(err)
	}

	detail, err := c.UserDetail(context.Background(), &pb.UserDetailReq{Uid: "1"})
	if err != nil {
		panic(err)
	}

	fmt.Println(detail)
}
