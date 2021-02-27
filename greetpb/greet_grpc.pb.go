
package greetpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

const _ = grpc.SupportPackageIsVersion7


type GreetServiceClient interface {
	Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error)
	SendEcho(ctx context.Context, in *EchoMessage, opts ...grpc.CallOption) (GreetService_SendEchoClient, error)
	SendMultipleEcho(ctx context.Context, opts ...grpc.CallOption) (GreetService_SendMultipleEchoClient, error)
}

type greetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGreetServiceClient(cc grpc.ClientConnInterface) GreetServiceClient {
	return &greetServiceClient{cc}
}

func (c *greetServiceClient) Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error) {
	out := new(GreetResponse)
	err := c.cc.Invoke(ctx, "/greet.GreetService/Greet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greetServiceClient) SendEcho(ctx context.Context, in *EchoMessage, opts ...grpc.CallOption) (GreetService_SendEchoClient, error) {
	stream, err := c.cc.NewStream(ctx, &GreetService_ServiceDesc.Streams[0], "/greet.GreetService/SendEcho", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceSendEchoClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GreetService_SendEchoClient interface {
	Recv() (*EchoResponse, error)
	grpc.ClientStream
}

type greetServiceSendEchoClient struct {
	grpc.ClientStream
}

func (x *greetServiceSendEchoClient) Recv() (*EchoResponse, error) {
	m := new(EchoResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greetServiceClient) SendMultipleEcho(ctx context.Context, opts ...grpc.CallOption) (GreetService_SendMultipleEchoClient, error) {
	stream, err := c.cc.NewStream(ctx, &GreetService_ServiceDesc.Streams[1], "/greet.GreetService/SendMultipleEcho", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceSendMultipleEchoClient{stream}
	return x, nil
}

type GreetService_SendMultipleEchoClient interface {
	Send(*EchoMessage) error
	Recv() (*EchoResponse, error)
	grpc.ClientStream
}

type greetServiceSendMultipleEchoClient struct {
	grpc.ClientStream
}

func (x *greetServiceSendMultipleEchoClient) Send(m *EchoMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greetServiceSendMultipleEchoClient) Recv() (*EchoResponse, error) {
	m := new(EchoResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

type GreetServiceServer interface {
	Greet(context.Context, *GreetRequest) (*GreetResponse, error)
	SendEcho(*EchoMessage, GreetService_SendEchoServer) error
	SendMultipleEcho(GreetService_SendMultipleEchoServer) error
	mustEmbedUnimplementedGreetServiceServer()
}

// UnimplementedGreetServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGreetServiceServer struct {
}

func (UnimplementedGreetServiceServer) Greet(context.Context, *GreetRequest) (*GreetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Greet not implemented")
}
func (UnimplementedGreetServiceServer) SendEcho(*EchoMessage, GreetService_SendEchoServer) error {
	return status.Errorf(codes.Unimplemented, "method SendEcho not implemented")
}
func (UnimplementedGreetServiceServer) SendMultipleEcho(GreetService_SendMultipleEchoServer) error {
	return status.Errorf(codes.Unimplemented, "method SendMultipleEcho not implemented")
}
func (UnimplementedGreetServiceServer) mustEmbedUnimplementedGreetServiceServer() {}

// UnsafeGreetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreetServiceServer will
// result in compilation errors.
type UnsafeGreetServiceServer interface {
	mustEmbedUnimplementedGreetServiceServer()
}

func RegisterGreetServiceServer(s grpc.ServiceRegistrar, srv GreetServiceServer) {
	s.RegisterService(&GreetService_ServiceDesc, srv)
}

func _GreetService_Greet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetServiceServer).Greet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greet.GreetService/Greet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetServiceServer).Greet(ctx, req.(*GreetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GreetService_SendEcho_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EchoMessage)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreetServiceServer).SendEcho(m, &greetServiceSendEchoServer{stream})
}

type GreetService_SendEchoServer interface {
	Send(*EchoResponse) error
	grpc.ServerStream
}

type greetServiceSendEchoServer struct {
	grpc.ServerStream
}

func (x *greetServiceSendEchoServer) Send(m *EchoResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _GreetService_SendMultipleEcho_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetServiceServer).SendMultipleEcho(&greetServiceSendMultipleEchoServer{stream})
}

type GreetService_SendMultipleEchoServer interface {
	Send(*EchoResponse) error
	Recv() (*EchoMessage, error)
	grpc.ServerStream
}

type greetServiceSendMultipleEchoServer struct {
	grpc.ServerStream
}

func (x *greetServiceSendMultipleEchoServer) Send(m *EchoResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greetServiceSendMultipleEchoServer) Recv() (*EchoMessage, error) {
	m := new(EchoMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreetService_ServiceDesc is the grpc.ServiceDesc for GreetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GreetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "greet.GreetService",
	HandlerType: (*GreetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Greet",
			Handler:    _GreetService_Greet_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendEcho",
			Handler:       _GreetService_SendEcho_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SendMultipleEcho",
			Handler:       _GreetService_SendMultipleEcho_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "greet/greetpb/greet.proto",
}

// CalculatorServiceClient is the client API for CalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalculatorServiceClient interface {
	GetSum(ctx context.Context, in *NumbersRequest, opts ...grpc.CallOption) (*NumbersResponse, error)
	PrimeNumberDecomposition(ctx context.Context, in *IntegerRequest, opts ...grpc.CallOption) (CalculatorService_PrimeNumberDecompositionClient, error)
	ComputeAverage(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_ComputeAverageClient, error)
}

type calculatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculatorServiceClient(cc grpc.ClientConnInterface) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) GetSum(ctx context.Context, in *NumbersRequest, opts ...grpc.CallOption) (*NumbersResponse, error) {
	out := new(NumbersResponse)
	err := c.cc.Invoke(ctx, "/greet.CalculatorService/GetSum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) PrimeNumberDecomposition(ctx context.Context, in *IntegerRequest, opts ...grpc.CallOption) (CalculatorService_PrimeNumberDecompositionClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalculatorService_ServiceDesc.Streams[0], "/greet.CalculatorService/PrimeNumberDecomposition", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServicePrimeNumberDecompositionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalculatorService_PrimeNumberDecompositionClient interface {
	Recv() (*IntegerResponse, error)
	grpc.ClientStream
}

type calculatorServicePrimeNumberDecompositionClient struct {
	grpc.ClientStream
}

func (x *calculatorServicePrimeNumberDecompositionClient) Recv() (*IntegerResponse, error) {
	m := new(IntegerResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) ComputeAverage(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_ComputeAverageClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalculatorService_ServiceDesc.Streams[1], "/greet.CalculatorService/ComputeAverage", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceComputeAverageClient{stream}
	return x, nil
}

type CalculatorService_ComputeAverageClient interface {
	Send(*IntegerRequest) error
	CloseAndRecv() (*AverageResponse, error)
	grpc.ClientStream
}

type calculatorServiceComputeAverageClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceComputeAverageClient) Send(m *IntegerRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorServiceComputeAverageClient) CloseAndRecv() (*AverageResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(AverageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalculatorServiceServer is the server API for CalculatorService service.
// All implementations must embed UnimplementedCalculatorServiceServer
// for forward compatibility
type CalculatorServiceServer interface {
	GetSum(context.Context, *NumbersRequest) (*NumbersResponse, error)
	PrimeNumberDecomposition(*IntegerRequest, CalculatorService_PrimeNumberDecompositionServer) error
	ComputeAverage(CalculatorService_ComputeAverageServer) error
	mustEmbedUnimplementedCalculatorServiceServer()
}

// UnimplementedCalculatorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCalculatorServiceServer struct {
}

func (UnimplementedCalculatorServiceServer) GetSum(context.Context, *NumbersRequest) (*NumbersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSum not implemented")
}
func (UnimplementedCalculatorServiceServer) PrimeNumberDecomposition(*IntegerRequest, CalculatorService_PrimeNumberDecompositionServer) error {
	return status.Errorf(codes.Unimplemented, "method PrimeNumberDecomposition not implemented")
}
func (UnimplementedCalculatorServiceServer) ComputeAverage(CalculatorService_ComputeAverageServer) error {
	return status.Errorf(codes.Unimplemented, "method ComputeAverage not implemented")
}
func (UnimplementedCalculatorServiceServer) mustEmbedUnimplementedCalculatorServiceServer() {}

// UnsafeCalculatorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalculatorServiceServer will
// result in compilation errors.
type UnsafeCalculatorServiceServer interface {
	mustEmbedUnimplementedCalculatorServiceServer()
}

func RegisterCalculatorServiceServer(s grpc.ServiceRegistrar, srv CalculatorServiceServer) {
	s.RegisterService(&CalculatorService_ServiceDesc, srv)
}

func _CalculatorService_GetSum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NumbersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).GetSum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greet.CalculatorService/GetSum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).GetSum(ctx, req.(*NumbersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_PrimeNumberDecomposition_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(IntegerRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculatorServiceServer).PrimeNumberDecomposition(m, &calculatorServicePrimeNumberDecompositionServer{stream})
}

type CalculatorService_PrimeNumberDecompositionServer interface {
	Send(*IntegerResponse) error
	grpc.ServerStream
}

type calculatorServicePrimeNumberDecompositionServer struct {
	grpc.ServerStream
}

func (x *calculatorServicePrimeNumberDecompositionServer) Send(m *IntegerResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CalculatorService_ComputeAverage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).ComputeAverage(&calculatorServiceComputeAverageServer{stream})
}

type CalculatorService_ComputeAverageServer interface {
	SendAndClose(*AverageResponse) error
	Recv() (*IntegerRequest, error)
	grpc.ServerStream
}

type calculatorServiceComputeAverageServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceComputeAverageServer) SendAndClose(m *AverageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorServiceComputeAverageServer) Recv() (*IntegerRequest, error) {
	m := new(IntegerRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalculatorService_ServiceDesc is the grpc.ServiceDesc for CalculatorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CalculatorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "greet.CalculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSum",
			Handler:    _CalculatorService_GetSum_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PrimeNumberDecomposition",
			Handler:       _CalculatorService_PrimeNumberDecomposition_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ComputeAverage",
			Handler:       _CalculatorService_ComputeAverage_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "greet/greetpb/greet.proto",
}
