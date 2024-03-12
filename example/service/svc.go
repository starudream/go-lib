package main

import (
	"context"
	"fmt"

	"github.com/oklog/ulid/v2"

	"github.com/starudream/go-lib/core/v2/config/version"

	"github.com/starudream/go-lib/example/v2/api/admin"
	"github.com/starudream/go-lib/example/v2/api/admin/user"
	"github.com/starudream/go-lib/example/v2/api/common"
)

type AdminUserService struct {
	admin.UnimplementedAdminUserServiceServer
}

var _ admin.AdminUserServiceServer = (*AdminUserService)(nil)

func (s *AdminUserService) Health(context.Context, *common.Empty) (*common.Struct, error) {
	vi := version.GetVersionInfo()
	return common.NewStruct(map[string]any{
		"version": vi.GitVersion,
	})
}

func (s *AdminUserService) AddUser(context.Context, *user.AddUserReq) (*common.Id, error) {
	return &common.Id{Id: ulid.Make().String()}, nil
}

func (s *AdminUserService) GetUser(context.Context, *user.GetUserReq) (*user.User, error) {
	return nil, fmt.Errorf("not found")
}
