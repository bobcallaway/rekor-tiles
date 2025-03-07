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
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Rekor_Echo_FullMethodName = "/dev.sigstore.rekor.v2.Rekor/Echo"
)

// RekorClient is the client API for Rekor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RekorClient interface {
	Echo(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (*StringMessage, error)
}

type rekorClient struct {
	cc grpc.ClientConnInterface
}

func NewRekorClient(cc grpc.ClientConnInterface) RekorClient {
	return &rekorClient{cc}
}

func (c *rekorClient) Echo(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (*StringMessage, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StringMessage)
	err := c.cc.Invoke(ctx, Rekor_Echo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RekorServer is the server API for Rekor service.
// All implementations must embed UnimplementedRekorServer
// for forward compatibility.
type RekorServer interface {
	Echo(context.Context, *StringMessage) (*StringMessage, error)
	mustEmbedUnimplementedRekorServer()
}

// UnimplementedRekorServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRekorServer struct{}

func (UnimplementedRekorServer) Echo(context.Context, *StringMessage) (*StringMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
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

func _Rekor_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RekorServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rekor_Echo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RekorServer).Echo(ctx, req.(*StringMessage))
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
			MethodName: "Echo",
			Handler:    _Rekor_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rekor_service.proto",
}
