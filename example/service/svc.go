package main

import (
	"context"
	"embed"
	"fmt"
	"io/fs"

	"github.com/starudream/go-lib/core/v2/config/version"
	"github.com/starudream/go-lib/core/v2/gh"
	"github.com/starudream/go-lib/core/v2/slog"
	"github.com/starudream/go-lib/core/v2/utils/osutil"
	"github.com/starudream/go-lib/server/v2/grpc"
	"github.com/starudream/go-lib/server/v2/hggw"
	"github.com/starudream/go-lib/server/v2/http"
	"github.com/starudream/go-lib/server/v2/http/middlewares"
	"github.com/starudream/go-lib/server/v2/ictx"
	"github.com/starudream/go-lib/server/v2/ierr"
	"github.com/starudream/go-lib/server/v2/jwt"

	"github.com/starudream/go-lib/example/v2/api/admin"
	"github.com/starudream/go-lib/example/v2/api/admin/user"
	"github.com/starudream/go-lib/example/v2/api/common"
)

//go:embed dist
var dist embed.FS

func NewHTTPServer() *hggw.Server {
	hs := hggw.NewServer(hggw.WithMountPath("/api"))
	hs.Use(middlewares.CORS())
	hs.RegisterHandler(admin.RegisterAdminUserServiceHandler)
	hs.With(middlewares.Logger()).Group(func(r http.Router) {
		r.HandleCtx("GET /token", func(c *http.Context) error {
			token, err := jwt.New("*", "starudream", "*").Sign()
			if err != nil {
				return err
			}
			return c.JSON(200, gh.M{"token": token})
		})
		r.HandleCtx("GET /panic", func(c *http.Context) error {
			panic("panic")
		})
		r.HandleCtx("GET /logger", func(c *http.Context) error {
			if c.GetQuery("error").Bool() {
				return ierr.BadRequest(9, "request error")
			}
			return c.JSON(200, gh.M{"foo": "bar"})
		})
		r.With(middlewares.JWT()).HandleCtx("GET /admin/user/add", func(c *http.Context) error {
			jc := jwt.MustFromContext(c)
			slog.Info("subject: %s", jc.SUB(), slog.GetAttrs(c))
			return c.JSON(200, "ok")
		})
	})
	http.FileServer(hs, "/", osutil.Must1(fs.Sub(dist, "dist")))
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
	jc := jwt.MustFromContext(ctx)
	slog.Info("password: %s, subject: %s", req.Password, jc.SUB(), slog.GetAttrs(ctx))
	return nil, ierr.Forbidden(1, "no permission")
}

func (s *AdminUserService) GetUser(context.Context, *user.GetUserReq) (*user.User, error) {
	return &user.User{Username: "admin", Password: "password"}, nil
}
