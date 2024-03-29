// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
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
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: eventmesh-client.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PublisherServiceClient is the client API for PublisherService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PublisherServiceClient interface {
	// Async event publish
	Publish(ctx context.Context, in *SimpleMessage, opts ...grpc.CallOption) (*Response, error)
	// Sync event publish
	RequestReply(ctx context.Context, in *SimpleMessage, opts ...grpc.CallOption) (*SimpleMessage, error)
	// Async batch event publish
	BatchPublish(ctx context.Context, in *BatchMessage, opts ...grpc.CallOption) (*Response, error)
}

type publisherServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPublisherServiceClient(cc grpc.ClientConnInterface) PublisherServiceClient {
	return &publisherServiceClient{cc}
}

func (c *publisherServiceClient) Publish(ctx context.Context, in *SimpleMessage, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/eventmesh.common.protocol.grpc.PublisherService/publish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publisherServiceClient) RequestReply(ctx context.Context, in *SimpleMessage, opts ...grpc.CallOption) (*SimpleMessage, error) {
	out := new(SimpleMessage)
	err := c.cc.Invoke(ctx, "/eventmesh.common.protocol.grpc.PublisherService/requestReply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publisherServiceClient) BatchPublish(ctx context.Context, in *BatchMessage, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/eventmesh.common.protocol.grpc.PublisherService/batchPublish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PublisherServiceServer is the server API for PublisherService service.
// All implementations must embed UnimplementedPublisherServiceServer
// for forward compatibility
type PublisherServiceServer interface {
	// Async event publish
	Publish(context.Context, *SimpleMessage) (*Response, error)
	// Sync event publish
	RequestReply(context.Context, *SimpleMessage) (*SimpleMessage, error)
	// Async batch event publish
	BatchPublish(context.Context, *BatchMessage) (*Response, error)
	mustEmbedUnimplementedPublisherServiceServer()
}

// UnimplementedPublisherServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPublisherServiceServer struct {
}

func (UnimplementedPublisherServiceServer) Publish(context.Context, *SimpleMessage) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (UnimplementedPublisherServiceServer) RequestReply(context.Context, *SimpleMessage) (*SimpleMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestReply not implemented")
}
func (UnimplementedPublisherServiceServer) BatchPublish(context.Context, *BatchMessage) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchPublish not implemented")
}
func (UnimplementedPublisherServiceServer) mustEmbedUnimplementedPublisherServiceServer() {}

// UnsafePublisherServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PublisherServiceServer will
// result in compilation errors.
type UnsafePublisherServiceServer interface {
	mustEmbedUnimplementedPublisherServiceServer()
}

func RegisterPublisherServiceServer(s grpc.ServiceRegistrar, srv PublisherServiceServer) {
	s.RegisterService(&PublisherService_ServiceDesc, srv)
}

func _PublisherService_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimpleMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublisherServiceServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventmesh.common.protocol.grpc.PublisherService/publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublisherServiceServer).Publish(ctx, req.(*SimpleMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _PublisherService_RequestReply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimpleMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublisherServiceServer).RequestReply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventmesh.common.protocol.grpc.PublisherService/requestReply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublisherServiceServer).RequestReply(ctx, req.(*SimpleMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _PublisherService_BatchPublish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublisherServiceServer).BatchPublish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventmesh.common.protocol.grpc.PublisherService/batchPublish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublisherServiceServer).BatchPublish(ctx, req.(*BatchMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// PublisherService_ServiceDesc is the grpc.ServiceDesc for PublisherService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PublisherService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eventmesh.common.protocol.grpc.PublisherService",
	HandlerType: (*PublisherServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "publish",
			Handler:    _PublisherService_Publish_Handler,
		},
		{
			MethodName: "requestReply",
			Handler:    _PublisherService_RequestReply_Handler,
		},
		{
			MethodName: "batchPublish",
			Handler:    _PublisherService_BatchPublish_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eventmesh-client.proto",
}

// ConsumerServiceClient is the client API for ConsumerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConsumerServiceClient interface {
	// The subscribed event will be delivered by invoking the webhook url in the Subscription
	Subscribe(ctx context.Context, in *Subscription, opts ...grpc.CallOption) (*Response, error)
	//  The subscribed event will be delivered through stream of Message
	SubscribeStream(ctx context.Context, opts ...grpc.CallOption) (ConsumerService_SubscribeStreamClient, error)
	Unsubscribe(ctx context.Context, in *Subscription, opts ...grpc.CallOption) (*Response, error)
}

type consumerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConsumerServiceClient(cc grpc.ClientConnInterface) ConsumerServiceClient {
	return &consumerServiceClient{cc}
}

func (c *consumerServiceClient) Subscribe(ctx context.Context, in *Subscription, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/eventmesh.common.protocol.grpc.ConsumerService/subscribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consumerServiceClient) SubscribeStream(ctx context.Context, opts ...grpc.CallOption) (ConsumerService_SubscribeStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &ConsumerService_ServiceDesc.Streams[0], "/eventmesh.common.protocol.grpc.ConsumerService/subscribeStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &consumerServiceSubscribeStreamClient{stream}
	return x, nil
}

type ConsumerService_SubscribeStreamClient interface {
	Send(*Subscription) error
	Recv() (*SimpleMessage, error)
	grpc.ClientStream
}

type consumerServiceSubscribeStreamClient struct {
	grpc.ClientStream
}

func (x *consumerServiceSubscribeStreamClient) Send(m *Subscription) error {
	return x.ClientStream.SendMsg(m)
}

func (x *consumerServiceSubscribeStreamClient) Recv() (*SimpleMessage, error) {
	m := new(SimpleMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *consumerServiceClient) Unsubscribe(ctx context.Context, in *Subscription, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/eventmesh.common.protocol.grpc.ConsumerService/unsubscribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConsumerServiceServer is the server API for ConsumerService service.
// All implementations must embed UnimplementedConsumerServiceServer
// for forward compatibility
type ConsumerServiceServer interface {
	// The subscribed event will be delivered by invoking the webhook url in the Subscription
	Subscribe(context.Context, *Subscription) (*Response, error)
	//  The subscribed event will be delivered through stream of Message
	SubscribeStream(ConsumerService_SubscribeStreamServer) error
	Unsubscribe(context.Context, *Subscription) (*Response, error)
	mustEmbedUnimplementedConsumerServiceServer()
}

// UnimplementedConsumerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedConsumerServiceServer struct {
}

func (UnimplementedConsumerServiceServer) Subscribe(context.Context, *Subscription) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedConsumerServiceServer) SubscribeStream(ConsumerService_SubscribeStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeStream not implemented")
}
func (UnimplementedConsumerServiceServer) Unsubscribe(context.Context, *Subscription) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unsubscribe not implemented")
}
func (UnimplementedConsumerServiceServer) mustEmbedUnimplementedConsumerServiceServer() {}

// UnsafeConsumerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConsumerServiceServer will
// result in compilation errors.
type UnsafeConsumerServiceServer interface {
	mustEmbedUnimplementedConsumerServiceServer()
}

func RegisterConsumerServiceServer(s grpc.ServiceRegistrar, srv ConsumerServiceServer) {
	s.RegisterService(&ConsumerService_ServiceDesc, srv)
}

func _ConsumerService_Subscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Subscription)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsumerServiceServer).Subscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventmesh.common.protocol.grpc.ConsumerService/subscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsumerServiceServer).Subscribe(ctx, req.(*Subscription))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConsumerService_SubscribeStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ConsumerServiceServer).SubscribeStream(&consumerServiceSubscribeStreamServer{stream})
}

type ConsumerService_SubscribeStreamServer interface {
	Send(*SimpleMessage) error
	Recv() (*Subscription, error)
	grpc.ServerStream
}

type consumerServiceSubscribeStreamServer struct {
	grpc.ServerStream
}

func (x *consumerServiceSubscribeStreamServer) Send(m *SimpleMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *consumerServiceSubscribeStreamServer) Recv() (*Subscription, error) {
	m := new(Subscription)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ConsumerService_Unsubscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Subscription)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsumerServiceServer).Unsubscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventmesh.common.protocol.grpc.ConsumerService/unsubscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsumerServiceServer).Unsubscribe(ctx, req.(*Subscription))
	}
	return interceptor(ctx, in, info, handler)
}

// ConsumerService_ServiceDesc is the grpc.ServiceDesc for ConsumerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConsumerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eventmesh.common.protocol.grpc.ConsumerService",
	HandlerType: (*ConsumerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "subscribe",
			Handler:    _ConsumerService_Subscribe_Handler,
		},
		{
			MethodName: "unsubscribe",
			Handler:    _ConsumerService_Unsubscribe_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "subscribeStream",
			Handler:       _ConsumerService_SubscribeStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "eventmesh-client.proto",
}

// HeartbeatServiceClient is the client API for HeartbeatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HeartbeatServiceClient interface {
	Heartbeat(ctx context.Context, in *Heartbeat, opts ...grpc.CallOption) (*Response, error)
}

type heartbeatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHeartbeatServiceClient(cc grpc.ClientConnInterface) HeartbeatServiceClient {
	return &heartbeatServiceClient{cc}
}

func (c *heartbeatServiceClient) Heartbeat(ctx context.Context, in *Heartbeat, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/eventmesh.common.protocol.grpc.HeartbeatService/heartbeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HeartbeatServiceServer is the server API for HeartbeatService service.
// All implementations must embed UnimplementedHeartbeatServiceServer
// for forward compatibility
type HeartbeatServiceServer interface {
	Heartbeat(context.Context, *Heartbeat) (*Response, error)
	mustEmbedUnimplementedHeartbeatServiceServer()
}

// UnimplementedHeartbeatServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHeartbeatServiceServer struct {
}

func (UnimplementedHeartbeatServiceServer) Heartbeat(context.Context, *Heartbeat) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Heartbeat not implemented")
}
func (UnimplementedHeartbeatServiceServer) mustEmbedUnimplementedHeartbeatServiceServer() {}

// UnsafeHeartbeatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HeartbeatServiceServer will
// result in compilation errors.
type UnsafeHeartbeatServiceServer interface {
	mustEmbedUnimplementedHeartbeatServiceServer()
}

func RegisterHeartbeatServiceServer(s grpc.ServiceRegistrar, srv HeartbeatServiceServer) {
	s.RegisterService(&HeartbeatService_ServiceDesc, srv)
}

func _HeartbeatService_Heartbeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Heartbeat)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeartbeatServiceServer).Heartbeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventmesh.common.protocol.grpc.HeartbeatService/heartbeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeartbeatServiceServer).Heartbeat(ctx, req.(*Heartbeat))
	}
	return interceptor(ctx, in, info, handler)
}

// HeartbeatService_ServiceDesc is the grpc.ServiceDesc for HeartbeatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HeartbeatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eventmesh.common.protocol.grpc.HeartbeatService",
	HandlerType: (*HeartbeatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "heartbeat",
			Handler:    _HeartbeatService_Heartbeat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eventmesh-client.proto",
}
