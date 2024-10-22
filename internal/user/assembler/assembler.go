package assembler

import (
	pb "kwai/api/proto/user"
	"kwai/internal/user/domain"
)

func UserEntity2ToPB(entity *domain.UserEntity) *pb.UserDetailResp {
	return &pb.UserDetailResp{
		Name: entity.Name,
		Age:  entity.Age,
	}
}
