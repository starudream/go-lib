module github.com/starudream/go-lib/example/v2

go 1.22

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
	github.com/starudream/go-lib/cobra/v2 v2.0.6
	github.com/starudream/go-lib/core/v2 v2.0.20
	github.com/starudream/go-lib/cron/v2 v2.0.6
	github.com/starudream/go-lib/ntfy/v2 v2.0.9
	github.com/starudream/go-lib/resty/v2 v2.0.9
	github.com/starudream/go-lib/selfupdate/v2 v2.0.4
	github.com/starudream/go-lib/server/v2 v2.0.0-rc.3
	github.com/starudream/go-lib/service/v2 v2.0.3
	github.com/starudream/go-lib/sqlite/v2 v2.0.2
	github.com/starudream/go-lib/tablew/v2 v2.0.5
)

require (
	github.com/Masterminds/semver/v3 v3.2.1
	github.com/oklog/ulid/v2 v2.1.0
	github.com/samber/lo v1.39.0
	golang.org/x/mod v0.16.0
)

require (
	github.com/VividCortex/ewma v1.2.0 // indirect
	github.com/cheggaaa/pb/v3 v3.1.5 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/envoyproxy/protoc-gen-validate v1.0.4 // indirect
	github.com/fatih/color v1.16.0 // indirect
	github.com/go-chi/chi/v5 v5.0.12 // indirect
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-ping/ping v1.1.0 // indirect
	github.com/go-resty/resty/v2 v2.11.0 // indirect
	github.com/go-viper/mapstructure/v2 v2.0.0-alpha.1 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/goccy/go-yaml v1.11.3 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.19.1 // indirect
	github.com/hashicorp/golang-lru/v2 v2.0.7 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/kardianos/service v1.2.2 // indirect
	github.com/knadh/koanf/maps v0.1.1 // indirect
	github.com/knadh/koanf/v2 v2.1.0 // indirect
	github.com/lmittmann/tint v1.0.4 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/ncruces/go-strftime v0.1.9 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/robfig/cron/v3 v3.0.1 // indirect
	github.com/soheilhy/cmux v0.1.5 // indirect
	github.com/spf13/cast v1.6.0 // indirect
	github.com/spf13/cobra v1.8.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.49.0 // indirect
	go.opentelemetry.io/contrib/propagators/b3 v1.24.0 // indirect
	go.opentelemetry.io/otel v1.24.0 // indirect
	go.opentelemetry.io/otel/metric v1.24.0 // indirect
	go.opentelemetry.io/otel/trace v1.24.0 // indirect
	golang.org/x/exp v0.0.0-20240314144324-c7f7c6466f7f // indirect
	golang.org/x/net v0.22.0 // indirect
	golang.org/x/sync v0.6.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	golang.org/x/xerrors v0.0.0-20231012003039-104605ab7028 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240314234333-6e1732d8331c // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240314234333-6e1732d8331c // indirect
	google.golang.org/grpc v1.62.1 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
	gorm.io/gorm v1.25.7 // indirect
	gorm.io/plugin/soft_delete v1.2.1 // indirect
	modernc.org/gc/v3 v3.0.0-20240304020402-f0dba7c97c2b // indirect
	modernc.org/libc v1.44.0 // indirect
	modernc.org/mathutil v1.6.0 // indirect
	modernc.org/memory v1.7.2 // indirect
	modernc.org/sqlite v1.29.3 // indirect
	modernc.org/strutil v1.2.0 // indirect
	modernc.org/token v1.1.0 // indirect
)
