package main

import (
	"context"
	"strconv"

	"github.com/oklog/ulid/v2"

	"github.com/starudream/go-lib/core/v2/codec/json"
	"github.com/starudream/go-lib/core/v2/slog"
	"github.com/starudream/go-lib/core/v2/utils/osutil"
	"github.com/starudream/go-lib/sqlite/v2"
)

type User struct {
	Account  string `gorm:"type:varchar(36);uniqueIndex:udx_account_deleted_at" json:"account"`
	Password string `gorm:"type:varchar(128)"                                   json:"password"`

	DeletedAt sqlite.DeletedAt `gorm:"uniqueIndex:udx_account_deleted_at" json:"deletedAt,omitempty"`

	sqlite.Meta
}

func init() {
	sqlite.GenId = func(ctx context.Context) string { return ulid.Make().String() }
	sqlite.GenUid = func(ctx context.Context) string { return "xxx" }
}

func main() {
	osutil.PanicErr(sqlite.DB().AutoMigrate(&User{}))

	user := &User{Account: "admin", Password: "admin"}
	osutil.PanicErr(sqlite.DB().FirstOrCreate(user).Error)

	for i := 1; i <= 3; i++ {
		user = &User{Account: "admin" + strconv.Itoa(i), Password: "admin" + strconv.Itoa(5-i)}
		osutil.PanicErr(sqlite.DB().Where("account=?", user.Account).FirstOrCreate(user).Error)
	}

	user = &User{}
	osutil.PanicErr(sqlite.DB().First(user, "account=?", "admin").Error)
	slog.Info(json.MustMarshalIndentString(user))

	ps := &sqlite.ListParams{
		Pagination: &sqlite.Pagination{
			Page:     1,
			PageSize: 3,
		},
		Sort: []string{"-password"},
	}
	slog.Info(json.MustMarshalIndentString(osutil.Must1(sqlite.QueryList[*User](sqlite.DB().Model(&User{}), ps))))
}
