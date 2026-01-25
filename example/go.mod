module github.com/starudream/go-lib/example/v2

go 1.24.0

replace github.com/starudream/go-lib/example/v2/api => ./proto/gen_go

require github.com/starudream/go-lib/example/v2/api v0.0.0-00010101000000-000000000000

replace (
	github.com/starudream/go-lib/cobra/v2 => ../cobra
	github.com/starudream/go-lib/core/v2 => ../core
	github.com/starudream/go-lib/cron/v2 => ../cron
	github.com/starudream/go-lib/ntfy/v2 => ../ntfy
	github.com/starudream/go-lib/resty/v2 => ../resty
	github.com/starudream/go-lib/selfupdate/v2 => ../selfupdate
	github.com/starudream/go-lib/server/v2 => ../server
	github.com/starudream/go-lib/service/v2 => ../service
	github.com/starudream/go-lib/sqlite/v2 => ../sqlite
	github.com/starudream/go-lib/tablew/v2 => ../tablew
)

require (
	github.com/starudream/go-lib/cobra/v2 v2.0.22
	github.com/starudream/go-lib/core/v2 v2.1.13
	github.com/starudream/go-lib/cron/v2 v2.0.22
	github.com/starudream/go-lib/ntfy/v2 v2.0.25
	github.com/starudream/go-lib/resty/v2 v2.0.24
	github.com/starudream/go-lib/selfupdate/v2 v2.0.17
	github.com/starudream/go-lib/server/v2 v2.0.7
	github.com/starudream/go-lib/service/v2 v2.0.17
	github.com/starudream/go-lib/sqlite/v2 v2.0.19
	github.com/starudream/go-lib/tablew/v2 v2.0.10
)

require (
	github.com/Masterminds/semver/v3 v3.4.0
	github.com/oklog/ulid/v2 v2.1.1
	github.com/samber/lo v1.52.0
	golang.org/x/mod v0.32.0
)

require (
	github.com/VividCortex/ewma v1.2.0 // indirect
	github.com/cheggaaa/pb/v3 v3.1.7 // indirect
	github.com/clipperhouse/stringish v0.1.1 // indirect
	github.com/clipperhouse/uax29/v2 v2.4.0 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/envoyproxy/protoc-gen-validate v1.3.0 // indirect
	github.com/fatih/color v1.18.0 // indirect
	github.com/go-chi/chi/v5 v5.2.4 // indirect
	github.com/go-resty/resty/v2 v2.17.1 // indirect
	github.com/go-viper/mapstructure/v2 v2.5.0 // indirect
	github.com/goccy/go-json v0.10.5 // indirect
	github.com/goccy/go-yaml v1.19.2 // indirect
	github.com/golang-jwt/jwt/v5 v5.3.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.27.5 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/kardianos/service v1.2.4 // indirect
	github.com/knadh/koanf/maps v0.1.2 // indirect
	github.com/knadh/koanf/v2 v2.3.2 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/lmittmann/tint v1.1.2 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.19 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/ncruces/go-strftime v1.0.0 // indirect
	github.com/prometheus-community/pro-bing v0.7.0 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec // indirect
	github.com/robfig/cron/v3 v3.0.1 // indirect
	github.com/rogpeppe/go-internal v1.14.1 // indirect
	github.com/rs/cors v1.11.1 // indirect
	github.com/soheilhy/cmux v0.1.5 // indirect
	github.com/spf13/cast v1.10.0 // indirect
	github.com/spf13/cobra v1.10.2 // indirect
	github.com/spf13/pflag v1.0.10 // indirect
	golang.org/x/exp v0.0.0-20260112195511-716be5621a96 // indirect
	golang.org/x/net v0.49.0 // indirect
	golang.org/x/sync v0.19.0 // indirect
	golang.org/x/sys v0.40.0 // indirect
	golang.org/x/text v0.33.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20260122232226-8e98ce8d340d // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260122232226-8e98ce8d340d // indirect
	google.golang.org/grpc v1.78.0 // indirect
	google.golang.org/protobuf v1.36.11 // indirect
	gorm.io/gorm v1.31.1 // indirect
	gorm.io/plugin/soft_delete v1.2.1 // indirect
	modernc.org/libc v1.67.6 // indirect
	modernc.org/mathutil v1.7.1 // indirect
	modernc.org/memory v1.11.0 // indirect
	modernc.org/sqlite v1.44.3 // indirect
)
