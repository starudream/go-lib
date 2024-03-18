package ierr_test

import (
	"fmt"
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/genproto/googleapis/rpc/errdetails"

	"github.com/starudream/go-lib/core/v2/codec/json"
	"github.com/starudream/go-lib/core/v2/utils/testutil"
	"github.com/starudream/go-lib/server/v2/ierr"
)

func TestError(t *testing.T) {
	e1 := ierr.BadRequest("empty account", "hello %s", "foo")
	testutil.Log(t, e1.String())
	testutil.Log(t, e1.GRPCStatus().String())
	testutil.Log(t, json.MustMarshalString(e1))

	e2 := e1.AppendMessage("hello %s", "bar")
	testutil.Log(t, e2)

	e3 := fmt.Errorf("hello %s", "world")

	e4 := ierr.FromError(e3)
	testutil.Log(t, e4)

	e5, _ := status.New(codes.NotFound, "not fount").WithDetails(&errdetails.ErrorInfo{Reason: "???"})

	e6 := ierr.FromError(e5.Err())
	testutil.Log(t, e6, ierr.Code(e6), ierr.Reason(e6))
}
