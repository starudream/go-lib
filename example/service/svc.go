package main

import (
	"context"
	"fmt"

	"github.com/starudream/go-lib/core/v2/config/version"
	"github.com/starudream/go-lib/core/v2/slog"
	"github.com/starudream/go-lib/server/v2/grpc"
	"github.com/starudream/go-lib/server/v2/hggw"
	"github.com/starudream/go-lib/server/v2/http"
	"github.com/starudream/go-lib/server/v2/http/middlewares/cors"
	"github.com/starudream/go-lib/server/v2/ictx"
	"github.com/starudream/go-lib/server/v2/ierr"

	"github.com/starudream/go-lib/example/v2/api/admin"
	"github.com/starudream/go-lib/example/v2/api/admin/user"
	"github.com/starudream/go-lib/example/v2/api/common"
)

func NewHTTPServer() *hggw.Server {
	hs := hggw.NewServer()
	hs.Use(cors.AllowAll())
	hs.Get("/admin/user/add", func(c *http.Context) error { return c.JSON(200, version.GetVersionInfo()) })
	hs.RegisterHandler(admin.RegisterAdminUserServiceHandler)
	return hs
}

func NewGRPCServer() *grpc.Server {
	gs := grpc.NewServer()
	gs.RegisterServer(admin.RegisterAdminUserServiceServer, &AdminUserService{})
	return gs
}

type AdminUserService struct {
	admin.UnimplementedAdminUserServiceServer
}

var _ admin.AdminUserServiceServer = (*AdminUserService)(nil)

func (s *AdminUserService) Health(ctx context.Context, _ *common.Empty) (*common.Struct, error) {
	ictx.FromContext(ctx).Range(func(k string, vs []string) bool { fmt.Printf("\t%-30s - %s\n", k, vs); return true })
	return common.NewStruct(map[string]any{"version": version.GetVersionInfo().GitVersion})
}

func (s *AdminUserService) AddUser(ctx context.Context, req *user.AddUserReq) (*common.Id, error) {
	slog.Info("password: %s", req.Password, slog.GetAttrs(ctx))
	return nil, ierr.Forbidden(1, "no permission")
}

func (s *AdminUserService) GetUser(context.Context, *user.GetUserReq) (*user.User, error) {
	return &user.User{Username: "admin", Password: "password"}, nil
}
