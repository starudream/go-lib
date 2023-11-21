#!/usr/bin/env bash

# https://github.com/go-gorm/sqlite

wget -qO internal/driver/ddlmod.go https://raw.githubusercontent.com/go-gorm/sqlite/master/ddlmod.go
wget -qO internal/driver/error_translator.go https://raw.githubusercontent.com/go-gorm/sqlite/master/error_translator.go
wget -qO internal/driver/migrator.go https://raw.githubusercontent.com/go-gorm/sqlite/master/migrator.go
wget -qO internal/driver/sqlite.go https://raw.githubusercontent.com/go-gorm/sqlite/master/sqlite.go
