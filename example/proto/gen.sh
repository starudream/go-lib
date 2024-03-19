#!/usr/bin/env bash

set -e

echo "Cleaning"
rm -rf gen_*

echo "Generating proto files"
buf generate

echo "Generating go mod files"
pushd gen_go >/dev/null
rm -rf doc google protoc-gen-openapiv2 validate
cat<<EOF > "go.mod"
module github.com/starudream/go-lib/example/v2/api

go 1.22

require (
	github.com/envoyproxy/protoc-gen-validate v1.0.4
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.19.1
	google.golang.org/genproto/googleapis/api v0.0.0-20240318140521-94a12d6c2237
	google.golang.org/grpc v1.62.1
	google.golang.org/protobuf v1.33.0
)

require (
	github.com/golang/protobuf v1.5.4 // indirect
	golang.org/x/net v0.22.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240318140521-94a12d6c2237 // indirect
)
EOF
go mod tidy

echo "Generating common helper files"
cat<<EOF > "common/common.pb.helper.go"
package common

import (
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var (
	NewDuration    = durationpb.New
	NewTimestamp   = timestamppb.New
	NewStruct      = structpb.NewStruct
	NewValue       = structpb.NewValue
	NewNullValue   = structpb.NewNullValue
	NewBoolValue   = structpb.NewBoolValue
	NewNumberValue = structpb.NewNumberValue
	NewStringValue = structpb.NewStringValue
	NewStructValue = structpb.NewStructValue
	NewListValue   = structpb.NewListValue
	NewList        = structpb.NewList

	Double = wrapperspb.Double
	Float  = wrapperspb.Float
	Int64  = wrapperspb.Int64
	UInt64 = wrapperspb.UInt64
	Int32  = wrapperspb.Int32
	UInt32 = wrapperspb.UInt32
	Bool   = wrapperspb.Bool
	String = wrapperspb.String
	Bytes  = wrapperspb.Bytes

	NewAny          = anypb.New
	AnyMarshalFrom  = anypb.MarshalFrom
	AnyUnmarshalTo  = anypb.UnmarshalTo
	AnyUnmarshalNew = anypb.UnmarshalNew
)
EOF
popd >/dev/null

echo "Done"
