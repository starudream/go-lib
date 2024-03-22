package main

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/codec/json"
	"github.com/starudream/go-lib/core/v2/utils/testutil"
	"github.com/starudream/go-lib/server/v2/grpc"
	"github.com/starudream/go-lib/server/v2/ictx"

	"github.com/starudream/go-lib/example/v2/api/admin"
	"github.com/starudream/go-lib/example/v2/api/admin/user"
	"github.com/starudream/go-lib/example/v2/api/common"
)

var (
	ctx = ictx.SYSTEM()

	cliAdminUser admin.AdminUserServiceClient
)

func TestMain(m *testing.M) {
	conn, wait := grpc.NewMock(NewGRPCServer())
	cliAdminUser = admin.NewAdminUserServiceClient(conn)
	wait(m)
}

func TestAdminUserService_Health(t *testing.T) {
	resp, err := cliAdminUser.Health(ctx, &common.Empty{})
	testutil.LogNoErr(t, err, json.MustMarshalString(resp))
}

func TestAdminUserService_AddUser(t *testing.T) {
	req := &user.AddUserReq{
		Username:    "admin",
		Password:    "p1ssw0rd",
		DisplayName: "administrator",
	}
	_, err := cliAdminUser.AddUser(ctx, req)
	testutil.NotNil(t, err)
}

func TestAdminUserService_GetUser(t *testing.T) {
	req := &user.GetUserReq{
		X: &user.GetUserReq_Id{
			Id: "abcd1234",
		},
	}
	resp, err := cliAdminUser.GetUser(ctx, req)
	testutil.LogNoErr(t, err, json.MustMarshalString(resp))
}
