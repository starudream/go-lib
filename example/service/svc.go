package main

import (
	"context"
	"fmt"

	"github.com/oklog/ulid/v2"

	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"

	"github.com/starudream/go-lib/core/v2/config/version"
	"github.com/starudream/go-lib/server/v2/grpc"

	"github.com/starudream/go-lib/example/v2/api/admin"
	"github.com/starudream/go-lib/example/v2/api/admin/user"
	"github.com/starudream/go-lib/example/v2/api/common"
)

type AdminUserService struct {
	admin.UnimplementedAdminUserServiceServer
}

var _ admin.AdminUserServiceServer = (*AdminUserService)(nil)

func (s *AdminUserService) Health(ctx context.Context, _ *common.Empty) (*common.Struct, error) {
	md := grpc.GetMD(ctx)
	ks := maps.Keys(md)
	slices.Sort(ks)
	for i := 0; i < len(ks); i++ {
		fmt.Printf("\t%-30s - %s\n", ks[i], md.Get(ks[i]))
	}
	return common.NewStruct(map[string]any{"version": version.GetVersionInfo().GitVersion})
}

func (s *AdminUserService) AddUser(context.Context, *user.AddUserReq) (*common.Id, error) {
	return &common.Id{Id: ulid.Make().String()}, nil
}

func (s *AdminUserService) GetUser(context.Context, *user.GetUserReq) (*user.User, error) {
	return nil, fmt.Errorf("not found")
}
