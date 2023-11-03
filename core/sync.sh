#!/usr/bin/env bash

# config

wget -qO config/version/version.go https://github.com/kubernetes-sigs/release-utils/raw/main/version/version.go

wget -qO config/internal/parsers/yaml/yaml.go https://raw.githubusercontent.com/knadh/koanf/master/parsers/yaml/yaml.go

wget -qO config/internal/providers/confmap/confmap.go https://github.com/knadh/koanf/raw/master/providers/confmap/confmap.go
wget -qO config/internal/providers/env/env.go https://raw.githubusercontent.com/knadh/koanf/master/providers/env/env.go
wget -qO config/internal/providers/file/file.go https://raw.githubusercontent.com/knadh/koanf/master/providers/file/file.go
wget -qO config/internal/providers/posflag/posflag.go https://raw.githubusercontent.com/knadh/koanf/master/providers/posflag/posflag.go

# slog

wget -qO slog/internal/handler/lumberjack.go https://raw.githubusercontent.com/natefinch/lumberjack/v2.0/lumberjack.go

wget -qO slog/internal/handler/chown.go https://github.com/natefinch/lumberjack/raw/v2.0/chown.go
wget -qO slog/internal/handler/chown_linux.go https://github.com/natefinch/lumberjack/raw/v2.0/chown_linux.go
