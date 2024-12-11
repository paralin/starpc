// Code generated by protoc-gen-srpc. DO NOT EDIT.
// protoc-gen-srpc version: v0.32.15
// source: github.com/aperturerobotics/starpc/echo/echo.proto

package echo

import (
	context "context"

	emptypb "github.com/aperturerobotics/protobuf-go-lite/types/known/emptypb"
	rpcstream "github.com/aperturerobotics/starpc/rpcstream"
	srpc "github.com/aperturerobotics/starpc/srpc"
)

type SRPCEchoerClient interface {
	// SRPCClient returns the underlying SRPC client.
	SRPCClient() srpc.Client

	// Echo returns the given message.
	Echo(ctx context.Context, in *EchoMsg) (*EchoMsg, error)
	// EchoServerStream is an example of a server -> client one-way stream.
	EchoServerStream(ctx context.Context, in *EchoMsg) (SRPCEchoer_EchoServerStreamClient, error)
	// EchoClientStream is an example of client->server one-way stream.
	EchoClientStream(ctx context.Context) (SRPCEchoer_EchoClientStreamClient, error)
	// EchoBidiStream is an example of a two-way stream.
	EchoBidiStream(ctx context.Context) (SRPCEchoer_EchoBidiStreamClient, error)
	// RpcStream opens a nested rpc call stream.
	RpcStream(ctx context.Context) (SRPCEchoer_RpcStreamClient, error)
	// DoNothing does nothing.
	DoNothing(ctx context.Context, in *emptypb.Empty) (*emptypb.Empty, error)
}

type srpcEchoerClient struct {
	cc        srpc.Client
	serviceID string
}

func NewSRPCEchoerClient(cc srpc.Client) SRPCEchoerClient {
	return &srpcEchoerClient{cc: cc, serviceID: SRPCEchoerServiceID}
}

func NewSRPCEchoerClientWithServiceID(cc srpc.Client, serviceID string) SRPCEchoerClient {
	if serviceID == "" {
		serviceID = SRPCEchoerServiceID
	}
	return &srpcEchoerClient{cc: cc, serviceID: serviceID}
}

func (c *srpcEchoerClient) SRPCClient() srpc.Client { return c.cc }

func (c *srpcEchoerClient) Echo(ctx context.Context, in *EchoMsg) (*EchoMsg, error) {
	out := new(EchoMsg)
	err := c.cc.ExecCall(ctx, c.serviceID, "Echo", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *srpcEchoerClient) EchoServerStream(ctx context.Context, in *EchoMsg) (SRPCEchoer_EchoServerStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, c.serviceID, "EchoServerStream", in)
	if err != nil {
		return nil, err
	}
	strm := &srpcEchoer_EchoServerStreamClient{stream}
	if err := strm.CloseSend(); err != nil {
		return nil, err
	}
	return strm, nil
}

type SRPCEchoer_EchoServerStreamClient interface {
	srpc.Stream
	Recv() (*EchoMsg, error)
	RecvTo(*EchoMsg) error
}

type srpcEchoer_EchoServerStreamClient struct {
	srpc.Stream
}

func (x *srpcEchoer_EchoServerStreamClient) Recv() (*EchoMsg, error) {
	m := new(EchoMsg)
	if err := x.MsgRecv(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (x *srpcEchoer_EchoServerStreamClient) RecvTo(m *EchoMsg) error {
	return x.MsgRecv(m)
}

func (c *srpcEchoerClient) EchoClientStream(ctx context.Context) (SRPCEchoer_EchoClientStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, c.serviceID, "EchoClientStream", nil)
	if err != nil {
		return nil, err
	}
	strm := &srpcEchoer_EchoClientStreamClient{stream}
	return strm, nil
}

type SRPCEchoer_EchoClientStreamClient interface {
	srpc.Stream
	Send(*EchoMsg) error
	CloseAndRecv() (*EchoMsg, error)
}

type srpcEchoer_EchoClientStreamClient struct {
	srpc.Stream
}

func (x *srpcEchoer_EchoClientStreamClient) Send(m *EchoMsg) error {
	if m == nil {
		return nil
	}
	return x.MsgSend(m)
}

func (x *srpcEchoer_EchoClientStreamClient) CloseAndRecv() (*EchoMsg, error) {
	if err := x.CloseSend(); err != nil {
		return nil, err
	}
	m := new(EchoMsg)
	if err := x.MsgRecv(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (x *srpcEchoer_EchoClientStreamClient) CloseAndMsgRecv(m *EchoMsg) error {
	if err := x.CloseSend(); err != nil {
		return err
	}
	return x.MsgRecv(m)
}

func (c *srpcEchoerClient) EchoBidiStream(ctx context.Context) (SRPCEchoer_EchoBidiStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, c.serviceID, "EchoBidiStream", nil)
	if err != nil {
		return nil, err
	}
	strm := &srpcEchoer_EchoBidiStreamClient{stream}
	return strm, nil
}

type SRPCEchoer_EchoBidiStreamClient interface {
	srpc.Stream
	Send(*EchoMsg) error
	Recv() (*EchoMsg, error)
	RecvTo(*EchoMsg) error
}

type srpcEchoer_EchoBidiStreamClient struct {
	srpc.Stream
}

func (x *srpcEchoer_EchoBidiStreamClient) Send(m *EchoMsg) error {
	if m == nil {
		return nil
	}
	return x.MsgSend(m)
}

func (x *srpcEchoer_EchoBidiStreamClient) Recv() (*EchoMsg, error) {
	m := new(EchoMsg)
	if err := x.MsgRecv(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (x *srpcEchoer_EchoBidiStreamClient) RecvTo(m *EchoMsg) error {
	return x.MsgRecv(m)
}

func (c *srpcEchoerClient) RpcStream(ctx context.Context) (SRPCEchoer_RpcStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, c.serviceID, "RpcStream", nil)
	if err != nil {
		return nil, err
	}
	strm := &srpcEchoer_RpcStreamClient{stream}
	return strm, nil
}

type SRPCEchoer_RpcStreamClient interface {
	srpc.Stream
	Send(*rpcstream.RpcStreamPacket) error
	Recv() (*rpcstream.RpcStreamPacket, error)
	RecvTo(*rpcstream.RpcStreamPacket) error
}

type srpcEchoer_RpcStreamClient struct {
	srpc.Stream
}

func (x *srpcEchoer_RpcStreamClient) Send(m *rpcstream.RpcStreamPacket) error {
	if m == nil {
		return nil
	}
	return x.MsgSend(m)
}

func (x *srpcEchoer_RpcStreamClient) Recv() (*rpcstream.RpcStreamPacket, error) {
	m := new(rpcstream.RpcStreamPacket)
	if err := x.MsgRecv(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (x *srpcEchoer_RpcStreamClient) RecvTo(m *rpcstream.RpcStreamPacket) error {
	return x.MsgRecv(m)
}

func (c *srpcEchoerClient) DoNothing(ctx context.Context, in *emptypb.Empty) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.ExecCall(ctx, c.serviceID, "DoNothing", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type SRPCEchoerServer interface {
	// Echo returns the given message.
	Echo(context.Context, *EchoMsg) (*EchoMsg, error)
	// EchoServerStream is an example of a server -> client one-way stream.
	EchoServerStream(*EchoMsg, SRPCEchoer_EchoServerStreamStream) error
	// EchoClientStream is an example of client->server one-way stream.
	EchoClientStream(SRPCEchoer_EchoClientStreamStream) (*EchoMsg, error)
	// EchoBidiStream is an example of a two-way stream.
	EchoBidiStream(SRPCEchoer_EchoBidiStreamStream) error
	// RpcStream opens a nested rpc call stream.
	RpcStream(SRPCEchoer_RpcStreamStream) error
	// DoNothing does nothing.
	DoNothing(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
}

const SRPCEchoerServiceID = "echo.Echoer"

type SRPCEchoerHandler struct {
	serviceID string
	impl      SRPCEchoerServer
}

// NewSRPCEchoerHandler constructs a new RPC handler.
// serviceID: if empty, uses default: echo.Echoer
func NewSRPCEchoerHandler(impl SRPCEchoerServer, serviceID string) srpc.Handler {
	if serviceID == "" {
		serviceID = SRPCEchoerServiceID
	}
	return &SRPCEchoerHandler{impl: impl, serviceID: serviceID}
}

// SRPCRegisterEchoer registers the implementation with the mux.
// Uses the default serviceID: echo.Echoer
func SRPCRegisterEchoer(mux srpc.Mux, impl SRPCEchoerServer) error {
	return mux.Register(NewSRPCEchoerHandler(impl, ""))
}

func (d *SRPCEchoerHandler) GetServiceID() string { return d.serviceID }

func (SRPCEchoerHandler) GetMethodIDs() []string {
	return []string{
		"Echo",
		"EchoServerStream",
		"EchoClientStream",
		"EchoBidiStream",
		"RpcStream",
		"DoNothing",
	}
}

func (d *SRPCEchoerHandler) InvokeMethod(
	serviceID, methodID string,
	strm srpc.Stream,
) (bool, error) {
	if serviceID != "" && serviceID != d.GetServiceID() {
		return false, nil
	}

	switch methodID {
	case "Echo":
		return true, d.InvokeMethod_Echo(d.impl, strm)
	case "EchoServerStream":
		return true, d.InvokeMethod_EchoServerStream(d.impl, strm)
	case "EchoClientStream":
		return true, d.InvokeMethod_EchoClientStream(d.impl, strm)
	case "EchoBidiStream":
		return true, d.InvokeMethod_EchoBidiStream(d.impl, strm)
	case "RpcStream":
		return true, d.InvokeMethod_RpcStream(d.impl, strm)
	case "DoNothing":
		return true, d.InvokeMethod_DoNothing(d.impl, strm)
	default:
		return false, nil
	}
}

func (SRPCEchoerHandler) InvokeMethod_Echo(impl SRPCEchoerServer, strm srpc.Stream) error {
	req := new(EchoMsg)
	if err := strm.MsgRecv(req); err != nil {
		return err
	}
	out, err := impl.Echo(strm.Context(), req)
	if err != nil {
		return err
	}
	return strm.MsgSend(out)
}

func (SRPCEchoerHandler) InvokeMethod_EchoServerStream(impl SRPCEchoerServer, strm srpc.Stream) error {
	req := new(EchoMsg)
	if err := strm.MsgRecv(req); err != nil {
		return err
	}
	serverStrm := &srpcEchoer_EchoServerStreamStream{strm}
	return impl.EchoServerStream(req, serverStrm)
}

func (SRPCEchoerHandler) InvokeMethod_EchoClientStream(impl SRPCEchoerServer, strm srpc.Stream) error {
	clientStrm := &srpcEchoer_EchoClientStreamStream{strm}
	out, err := impl.EchoClientStream(clientStrm)
	if err != nil {
		return err
	}
	return strm.MsgSend(out)
}

func (SRPCEchoerHandler) InvokeMethod_EchoBidiStream(impl SRPCEchoerServer, strm srpc.Stream) error {
	clientStrm := &srpcEchoer_EchoBidiStreamStream{strm}
	return impl.EchoBidiStream(clientStrm)
}

func (SRPCEchoerHandler) InvokeMethod_RpcStream(impl SRPCEchoerServer, strm srpc.Stream) error {
	clientStrm := &srpcEchoer_RpcStreamStream{strm}
	return impl.RpcStream(clientStrm)
}

func (SRPCEchoerHandler) InvokeMethod_DoNothing(impl SRPCEchoerServer, strm srpc.Stream) error {
	req := new(emptypb.Empty)
	if err := strm.MsgRecv(req); err != nil {
		return err
	}
	out, err := impl.DoNothing(strm.Context(), req)
	if err != nil {
		return err
	}
	return strm.MsgSend(out)
}

type SRPCEchoer_EchoStream interface {
	srpc.Stream
}

type srpcEchoer_EchoStream struct {
	srpc.Stream
}

type SRPCEchoer_EchoServerStreamStream interface {
	srpc.Stream
	Send(*EchoMsg) error
	SendAndClose(*EchoMsg) error
}

type srpcEchoer_EchoServerStreamStream struct {
	srpc.Stream
}

func (x *srpcEchoer_EchoServerStreamStream) Send(m *EchoMsg) error {
	return x.MsgSend(m)
}

func (x *srpcEchoer_EchoServerStreamStream) SendAndClose(m *EchoMsg) error {
	if m != nil {
		if err := x.MsgSend(m); err != nil {
			return err
		}
	}
	return x.CloseSend()
}

type SRPCEchoer_EchoClientStreamStream interface {
	srpc.Stream
	Recv() (*EchoMsg, error)
	RecvTo(*EchoMsg) error
}

type srpcEchoer_EchoClientStreamStream struct {
	srpc.Stream
}

func (x *srpcEchoer_EchoClientStreamStream) Recv() (*EchoMsg, error) {
	m := new(EchoMsg)
	if err := x.MsgRecv(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (x *srpcEchoer_EchoClientStreamStream) RecvTo(m *EchoMsg) error {
	return x.MsgRecv(m)
}

type SRPCEchoer_EchoBidiStreamStream interface {
	srpc.Stream
	Send(*EchoMsg) error
	SendAndClose(*EchoMsg) error
	Recv() (*EchoMsg, error)
	RecvTo(*EchoMsg) error
}

type srpcEchoer_EchoBidiStreamStream struct {
	srpc.Stream
}

func (x *srpcEchoer_EchoBidiStreamStream) Send(m *EchoMsg) error {
	return x.MsgSend(m)
}

func (x *srpcEchoer_EchoBidiStreamStream) SendAndClose(m *EchoMsg) error {
	if m != nil {
		if err := x.MsgSend(m); err != nil {
			return err
		}
	}
	return x.CloseSend()
}

func (x *srpcEchoer_EchoBidiStreamStream) Recv() (*EchoMsg, error) {
	m := new(EchoMsg)
	if err := x.MsgRecv(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (x *srpcEchoer_EchoBidiStreamStream) RecvTo(m *EchoMsg) error {
	return x.MsgRecv(m)
}

type SRPCEchoer_RpcStreamStream interface {
	srpc.Stream
	Send(*rpcstream.RpcStreamPacket) error
	SendAndClose(*rpcstream.RpcStreamPacket) error
	Recv() (*rpcstream.RpcStreamPacket, error)
	RecvTo(*rpcstream.RpcStreamPacket) error
}

type srpcEchoer_RpcStreamStream struct {
	srpc.Stream
}

func (x *srpcEchoer_RpcStreamStream) Send(m *rpcstream.RpcStreamPacket) error {
	return x.MsgSend(m)
}

func (x *srpcEchoer_RpcStreamStream) SendAndClose(m *rpcstream.RpcStreamPacket) error {
	if m != nil {
		if err := x.MsgSend(m); err != nil {
			return err
		}
	}
	return x.CloseSend()
}

func (x *srpcEchoer_RpcStreamStream) Recv() (*rpcstream.RpcStreamPacket, error) {
	m := new(rpcstream.RpcStreamPacket)
	if err := x.MsgRecv(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (x *srpcEchoer_RpcStreamStream) RecvTo(m *rpcstream.RpcStreamPacket) error {
	return x.MsgRecv(m)
}

type SRPCEchoer_DoNothingStream interface {
	srpc.Stream
}

type srpcEchoer_DoNothingStream struct {
	srpc.Stream
}
