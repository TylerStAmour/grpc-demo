package proto

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/runtime/protoimpl"
)

type Client interface {
	Ping(context.Context, *PingRequest, ...grpc.CallOption) (*PongReply, error)
}

// Struct only exists to hold this conn
type client struct {
	conn grpc.ClientConnInterface
}

type Server interface {
	Ping(context.Context, *PingRequest) (*PongReply, error)
}

type PingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

type PongReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (c *client) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PongReply, error) {
	reply := new(PongReply)
	err := c.conn.Invoke(ctx, "Ping", in, reply, opts...)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func pingHandler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}

	if interceptor == nil {
		return srv.(Server).Ping(ctx, in)
	}

	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "Ping",
	}

	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Server).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func NewClient(conn grpc.ClientConnInterface) Client {
	return &client{conn}
}

func RegisterServer(s grpc.ServiceRegistrar, srv Server) {
	s.RegisterService(&grpc.ServiceDesc{
		ServiceName: "grpc-demo.demo",
		HandlerType: (*Server)(nil),
		Methods: []grpc.MethodDesc{
			{
				MethodName: "Ping",
				Handler:    pingHandler,
			},
		},
		Streams:  []grpc.StreamDesc{},
		Metadata: "proto/demo.proto",
	}, srv)
}
