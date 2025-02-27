// Copyright 2025 The Sigstore Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: rekor_service.proto

package protobuf

import (
	context "context"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Rekor_CreateEntry_FullMethodName = "/dev.sigstore.rekor.v2.Rekor/CreateEntry"
)

// RekorClient is the client API for Rekor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// A service for sigstore clients to connect to to create log entries
type RekorClient interface {
	// Create an entry in the log
	CreateEntry(ctx context.Context, in *CreateEntryRequest, opts ...grpc.CallOption) (*TransparencyLogEntry, error)
}

type rekorClient struct {
	cc grpc.ClientConnInterface
}

func NewRekorClient(cc grpc.ClientConnInterface) RekorClient {
	return &rekorClient{cc}
}

func (c *rekorClient) CreateEntry(ctx context.Context, in *CreateEntryRequest, opts ...grpc.CallOption) (*TransparencyLogEntry, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TransparencyLogEntry)
	err := c.cc.Invoke(ctx, Rekor_CreateEntry_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RekorServer is the server API for Rekor service.
// All implementations must embed UnimplementedRekorServer
// for forward compatibility.
//
// A service for sigstore clients to connect to to create log entries
type RekorServer interface {
	// Create an entry in the log
	CreateEntry(context.Context, *CreateEntryRequest) (*TransparencyLogEntry, error)
	mustEmbedUnimplementedRekorServer()
}

// UnimplementedRekorServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRekorServer struct{}

func (UnimplementedRekorServer) CreateEntry(context.Context, *CreateEntryRequest) (*TransparencyLogEntry, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEntry not implemented")
}
func (UnimplementedRekorServer) mustEmbedUnimplementedRekorServer() {}
func (UnimplementedRekorServer) testEmbeddedByValue()               {}

// UnsafeRekorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RekorServer will
// result in compilation errors.
type UnsafeRekorServer interface {
	mustEmbedUnimplementedRekorServer()
}

func RegisterRekorServer(s grpc.ServiceRegistrar, srv RekorServer) {
	// If the following call pancis, it indicates UnimplementedRekorServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Rekor_ServiceDesc, srv)
}

func _Rekor_CreateEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RekorServer).CreateEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rekor_CreateEntry_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RekorServer).CreateEntry(ctx, req.(*CreateEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Rekor_ServiceDesc is the grpc.ServiceDesc for Rekor service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Rekor_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dev.sigstore.rekor.v2.Rekor",
	HandlerType: (*RekorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEntry",
			Handler:    _Rekor_CreateEntry_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rekor_service.proto",
}

const (
	RekorLog_GetTile_FullMethodName               = "/dev.sigstore.rekor.v2.RekorLog/GetTile"
	RekorLog_GetPartialTile_FullMethodName        = "/dev.sigstore.rekor.v2.RekorLog/GetPartialTile"
	RekorLog_GetEntryBundle_FullMethodName        = "/dev.sigstore.rekor.v2.RekorLog/GetEntryBundle"
	RekorLog_GetPartialEntryBundle_FullMethodName = "/dev.sigstore.rekor.v2.RekorLog/GetPartialEntryBundle"
)

// RekorLogClient is the client API for RekorLog service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// A service for log monitors and witnesses to audit/inspect the log
type RekorLogClient interface {
	// Get a tile from the log (cacheable)
	GetTile(ctx context.Context, in *TileRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error)
	// Get a partial tile from the log (not cacheable)
	GetPartialTile(ctx context.Context, in *PartialTileRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error)
	// Get an entry bundle from the log (cacheable)
	GetEntryBundle(ctx context.Context, in *EntryBundleRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error)
	// Get a partial entry bundle from the log (not cacheable)
	// TODO: I might have misunderstood the cacheability of this object
	GetPartialEntryBundle(ctx context.Context, in *PartialEntryBundleRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error)
}

type rekorLogClient struct {
	cc grpc.ClientConnInterface
}

func NewRekorLogClient(cc grpc.ClientConnInterface) RekorLogClient {
	return &rekorLogClient{cc}
}

func (c *rekorLogClient) GetTile(ctx context.Context, in *TileRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(httpbody.HttpBody)
	err := c.cc.Invoke(ctx, RekorLog_GetTile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rekorLogClient) GetPartialTile(ctx context.Context, in *PartialTileRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(httpbody.HttpBody)
	err := c.cc.Invoke(ctx, RekorLog_GetPartialTile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rekorLogClient) GetEntryBundle(ctx context.Context, in *EntryBundleRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(httpbody.HttpBody)
	err := c.cc.Invoke(ctx, RekorLog_GetEntryBundle_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rekorLogClient) GetPartialEntryBundle(ctx context.Context, in *PartialEntryBundleRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(httpbody.HttpBody)
	err := c.cc.Invoke(ctx, RekorLog_GetPartialEntryBundle_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RekorLogServer is the server API for RekorLog service.
// All implementations must embed UnimplementedRekorLogServer
// for forward compatibility.
//
// A service for log monitors and witnesses to audit/inspect the log
type RekorLogServer interface {
	// Get a tile from the log (cacheable)
	GetTile(context.Context, *TileRequest) (*httpbody.HttpBody, error)
	// Get a partial tile from the log (not cacheable)
	GetPartialTile(context.Context, *PartialTileRequest) (*httpbody.HttpBody, error)
	// Get an entry bundle from the log (cacheable)
	GetEntryBundle(context.Context, *EntryBundleRequest) (*httpbody.HttpBody, error)
	// Get a partial entry bundle from the log (not cacheable)
	// TODO: I might have misunderstood the cacheability of this object
	GetPartialEntryBundle(context.Context, *PartialEntryBundleRequest) (*httpbody.HttpBody, error)
	mustEmbedUnimplementedRekorLogServer()
}

// UnimplementedRekorLogServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRekorLogServer struct{}

func (UnimplementedRekorLogServer) GetTile(context.Context, *TileRequest) (*httpbody.HttpBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTile not implemented")
}
func (UnimplementedRekorLogServer) GetPartialTile(context.Context, *PartialTileRequest) (*httpbody.HttpBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPartialTile not implemented")
}
func (UnimplementedRekorLogServer) GetEntryBundle(context.Context, *EntryBundleRequest) (*httpbody.HttpBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEntryBundle not implemented")
}
func (UnimplementedRekorLogServer) GetPartialEntryBundle(context.Context, *PartialEntryBundleRequest) (*httpbody.HttpBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPartialEntryBundle not implemented")
}
func (UnimplementedRekorLogServer) mustEmbedUnimplementedRekorLogServer() {}
func (UnimplementedRekorLogServer) testEmbeddedByValue()                  {}

// UnsafeRekorLogServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RekorLogServer will
// result in compilation errors.
type UnsafeRekorLogServer interface {
	mustEmbedUnimplementedRekorLogServer()
}

func RegisterRekorLogServer(s grpc.ServiceRegistrar, srv RekorLogServer) {
	// If the following call pancis, it indicates UnimplementedRekorLogServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&RekorLog_ServiceDesc, srv)
}

func _RekorLog_GetTile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RekorLogServer).GetTile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RekorLog_GetTile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RekorLogServer).GetTile(ctx, req.(*TileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RekorLog_GetPartialTile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PartialTileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RekorLogServer).GetPartialTile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RekorLog_GetPartialTile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RekorLogServer).GetPartialTile(ctx, req.(*PartialTileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RekorLog_GetEntryBundle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EntryBundleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RekorLogServer).GetEntryBundle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RekorLog_GetEntryBundle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RekorLogServer).GetEntryBundle(ctx, req.(*EntryBundleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RekorLog_GetPartialEntryBundle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PartialEntryBundleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RekorLogServer).GetPartialEntryBundle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RekorLog_GetPartialEntryBundle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RekorLogServer).GetPartialEntryBundle(ctx, req.(*PartialEntryBundleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RekorLog_ServiceDesc is the grpc.ServiceDesc for RekorLog service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RekorLog_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dev.sigstore.rekor.v2.RekorLog",
	HandlerType: (*RekorLogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTile",
			Handler:    _RekorLog_GetTile_Handler,
		},
		{
			MethodName: "GetPartialTile",
			Handler:    _RekorLog_GetPartialTile_Handler,
		},
		{
			MethodName: "GetEntryBundle",
			Handler:    _RekorLog_GetEntryBundle_Handler,
		},
		{
			MethodName: "GetPartialEntryBundle",
			Handler:    _RekorLog_GetPartialEntryBundle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rekor_service.proto",
}
