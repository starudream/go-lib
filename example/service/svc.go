package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/oklog/ulid/v2"

	"github.com/starudream/go-lib/core/v2/config/version"
	"github.com/starudream/go-lib/server/v2/grpc"
	"github.com/starudream/go-lib/server/v2/otel"

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
	for k, vs := range md {
		fmt.Printf("%s -> %s\n", k, strings.Join(vs, ","))
	}
	fmt.Println("spanId", otel.SpanID(ctx))
	fmt.Println("traceId", otel.TraceID(ctx))
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
