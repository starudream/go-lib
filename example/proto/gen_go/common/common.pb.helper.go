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
