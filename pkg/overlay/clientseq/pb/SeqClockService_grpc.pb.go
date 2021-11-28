// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

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

// MessageQueueSeqClient is the client API for MessageQueueSeq service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessageQueueSeqClient interface {
	Enqueue(ctx context.Context, in *MessageSeq, opts ...grpc.CallOption) (*EnqueueReply, error)
}

type messageQueueSeqClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageQueueSeqClient(cc grpc.ClientConnInterface) MessageQueueSeqClient {
	return &messageQueueSeqClient{cc}
}

func (c *messageQueueSeqClient) Enqueue(ctx context.Context, in *MessageSeq, opts ...grpc.CallOption) (*EnqueueReply, error) {
	out := new(EnqueueReply)
	err := c.cc.Invoke(ctx, "/MessageQueueSeq/Enqueue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageQueueSeqServer is the server API for MessageQueueSeq service.
// All implementations must embed UnimplementedMessageQueueSeqServer
// for forward compatibility
type MessageQueueSeqServer interface {
	Enqueue(context.Context, *MessageSeq) (*EnqueueReply, error)
	mustEmbedUnimplementedMessageQueueSeqServer()
}

// UnimplementedMessageQueueSeqServer must be embedded to have forward compatible implementations.
type UnimplementedMessageQueueSeqServer struct {
}

func (UnimplementedMessageQueueSeqServer) Enqueue(context.Context, *MessageSeq) (*EnqueueReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Enqueue not implemented")
}
func (UnimplementedMessageQueueSeqServer) mustEmbedUnimplementedMessageQueueSeqServer() {}

// UnsafeMessageQueueSeqServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessageQueueSeqServer will
// result in compilation errors.
type UnsafeMessageQueueSeqServer interface {
	mustEmbedUnimplementedMessageQueueSeqServer()
}

func RegisterMessageQueueSeqServer(s grpc.ServiceRegistrar, srv MessageQueueSeqServer) {
	s.RegisterService(&MessageQueueSeq_ServiceDesc, srv)
}

func _MessageQueueSeq_Enqueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageSeq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageQueueSeqServer).Enqueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MessageQueueSeq/Enqueue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageQueueSeqServer).Enqueue(ctx, req.(*MessageSeq))
	}
	return interceptor(ctx, in, info, handler)
}

// MessageQueueSeq_ServiceDesc is the grpc.ServiceDesc for MessageQueueSeq service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MessageQueueSeq_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "MessageQueueSeq",
	HandlerType: (*MessageQueueSeqServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Enqueue",
			Handler:    _MessageQueueSeq_Enqueue_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "SeqClockService.proto",
}
