// Code generated by protoc-gen-go-lite. DO NOT EDIT.
// protoc-gen-go-lite version: v0.7.0
// source: github.com/aperturerobotics/starpc/srpc/rpcproto.proto

package srpc

import (
	base64 "encoding/base64"
	fmt "fmt"
	io "io"
	strconv "strconv"
	strings "strings"

	protobuf_go_lite "github.com/aperturerobotics/protobuf-go-lite"
	json "github.com/aperturerobotics/protobuf-go-lite/json"
)

// Packet is a message sent over a srpc packet connection.
type Packet struct {
	unknownFields []byte
	// Body is the packet body.
	//
	// Types that are assignable to Body:
	//
	//	*Packet_CallStart
	//	*Packet_CallData
	//	*Packet_CallCancel
	Body isPacket_Body `protobuf_oneof:"body"`
}

func (x *Packet) Reset() {
	*x = Packet{}
}

func (*Packet) ProtoMessage() {}

func (m *Packet) GetBody() isPacket_Body {
	if m != nil {
		return m.Body
	}
	return nil
}

func (x *Packet) GetCallStart() *CallStart {
	if x, ok := x.GetBody().(*Packet_CallStart); ok {
		return x.CallStart
	}
	return nil
}

func (x *Packet) GetCallData() *CallData {
	if x, ok := x.GetBody().(*Packet_CallData); ok {
		return x.CallData
	}
	return nil
}

func (x *Packet) GetCallCancel() bool {
	if x, ok := x.GetBody().(*Packet_CallCancel); ok {
		return x.CallCancel
	}
	return false
}

type isPacket_Body interface {
	isPacket_Body()
}

type Packet_CallStart struct {
	// CallStart initiates a new call.
	CallStart *CallStart `protobuf:"bytes,1,opt,name=call_start,json=callStart,proto3,oneof"`
}

type Packet_CallData struct {
	// CallData is a message in a streaming RPC sequence.
	CallData *CallData `protobuf:"bytes,2,opt,name=call_data,json=callData,proto3,oneof"`
}

type Packet_CallCancel struct {
	// CallCancel cancels the call.
	CallCancel bool `protobuf:"varint,3,opt,name=call_cancel,json=callCancel,proto3,oneof"`
}

func (*Packet_CallStart) isPacket_Body() {}

func (*Packet_CallData) isPacket_Body() {}

func (*Packet_CallCancel) isPacket_Body() {}

// CallStart requests starting a new RPC call.
type CallStart struct {
	unknownFields []byte
	// RpcService is the service to contact.
	// Must be set.
	RpcService string `protobuf:"bytes,1,opt,name=rpc_service,json=rpcService,proto3" json:"rpcService,omitempty"`
	// RpcMethod is the RPC method to call.
	// Must be set.
	RpcMethod string `protobuf:"bytes,2,opt,name=rpc_method,json=rpcMethod,proto3" json:"rpcMethod,omitempty"`
	// Data contains the request or the first message in the stream.
	// Optional if streaming.
	Data []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	// DataIsZero indicates Data is set with an empty message.
	DataIsZero bool `protobuf:"varint,4,opt,name=data_is_zero,json=dataIsZero,proto3" json:"dataIsZero,omitempty"`
}

func (x *CallStart) Reset() {
	*x = CallStart{}
}

func (*CallStart) ProtoMessage() {}

func (x *CallStart) GetRpcService() string {
	if x != nil {
		return x.RpcService
	}
	return ""
}

func (x *CallStart) GetRpcMethod() string {
	if x != nil {
		return x.RpcMethod
	}
	return ""
}

func (x *CallStart) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *CallStart) GetDataIsZero() bool {
	if x != nil {
		return x.DataIsZero
	}
	return false
}

// CallData contains a message in a streaming RPC sequence.
type CallData struct {
	unknownFields []byte
	// Data contains the packet in the sequence.
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	// DataIsZero indicates Data is set with an empty message.
	DataIsZero bool `protobuf:"varint,2,opt,name=data_is_zero,json=dataIsZero,proto3" json:"dataIsZero,omitempty"`
	// Complete indicates the RPC call is completed.
	Complete bool `protobuf:"varint,3,opt,name=complete,proto3" json:"complete,omitempty"`
	// Error contains any error that caused the RPC to fail.
	// If set, implies complete=true.
	Error string `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *CallData) Reset() {
	*x = CallData{}
}

func (*CallData) ProtoMessage() {}

func (x *CallData) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *CallData) GetDataIsZero() bool {
	if x != nil {
		return x.DataIsZero
	}
	return false
}

func (x *CallData) GetComplete() bool {
	if x != nil {
		return x.Complete
	}
	return false
}

func (x *CallData) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (m *Packet) CloneVT() *Packet {
	if m == nil {
		return (*Packet)(nil)
	}
	r := new(Packet)
	if m.Body != nil {
		r.Body = m.Body.(interface{ CloneOneofVT() isPacket_Body }).CloneOneofVT()
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *Packet) CloneMessageVT() protobuf_go_lite.CloneMessage {
	return m.CloneVT()
}

func (m *Packet_CallStart) CloneVT() *Packet_CallStart {
	if m == nil {
		return (*Packet_CallStart)(nil)
	}
	r := new(Packet_CallStart)
	r.CallStart = m.CallStart.CloneVT()
	return r
}

func (m *Packet_CallStart) CloneOneofVT() isPacket_Body {
	return m.CloneVT()
}

func (m *Packet_CallData) CloneVT() *Packet_CallData {
	if m == nil {
		return (*Packet_CallData)(nil)
	}
	r := new(Packet_CallData)
	r.CallData = m.CallData.CloneVT()
	return r
}

func (m *Packet_CallData) CloneOneofVT() isPacket_Body {
	return m.CloneVT()
}

func (m *Packet_CallCancel) CloneVT() *Packet_CallCancel {
	if m == nil {
		return (*Packet_CallCancel)(nil)
	}
	r := new(Packet_CallCancel)
	r.CallCancel = m.CallCancel
	return r
}

func (m *Packet_CallCancel) CloneOneofVT() isPacket_Body {
	return m.CloneVT()
}

func (m *CallStart) CloneVT() *CallStart {
	if m == nil {
		return (*CallStart)(nil)
	}
	r := new(CallStart)
	r.RpcService = m.RpcService
	r.RpcMethod = m.RpcMethod
	r.DataIsZero = m.DataIsZero
	if rhs := m.Data; rhs != nil {
		tmpBytes := make([]byte, len(rhs))
		copy(tmpBytes, rhs)
		r.Data = tmpBytes
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *CallStart) CloneMessageVT() protobuf_go_lite.CloneMessage {
	return m.CloneVT()
}

func (m *CallData) CloneVT() *CallData {
	if m == nil {
		return (*CallData)(nil)
	}
	r := new(CallData)
	r.DataIsZero = m.DataIsZero
	r.Complete = m.Complete
	r.Error = m.Error
	if rhs := m.Data; rhs != nil {
		tmpBytes := make([]byte, len(rhs))
		copy(tmpBytes, rhs)
		r.Data = tmpBytes
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *CallData) CloneMessageVT() protobuf_go_lite.CloneMessage {
	return m.CloneVT()
}

func (this *Packet) EqualVT(that *Packet) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Body == nil && that.Body != nil {
		return false
	} else if this.Body != nil {
		if that.Body == nil {
			return false
		}
		if !this.Body.(interface{ EqualVT(isPacket_Body) bool }).EqualVT(that.Body) {
			return false
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *Packet) EqualMessageVT(thatMsg any) bool {
	that, ok := thatMsg.(*Packet)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *Packet_CallStart) EqualVT(thatIface isPacket_Body) bool {
	that, ok := thatIface.(*Packet_CallStart)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.CallStart, that.CallStart; p != q {
		if p == nil {
			p = &CallStart{}
		}
		if q == nil {
			q = &CallStart{}
		}
		if !p.EqualVT(q) {
			return false
		}
	}
	return true
}

func (this *Packet_CallData) EqualVT(thatIface isPacket_Body) bool {
	that, ok := thatIface.(*Packet_CallData)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.CallData, that.CallData; p != q {
		if p == nil {
			p = &CallData{}
		}
		if q == nil {
			q = &CallData{}
		}
		if !p.EqualVT(q) {
			return false
		}
	}
	return true
}

func (this *Packet_CallCancel) EqualVT(thatIface isPacket_Body) bool {
	that, ok := thatIface.(*Packet_CallCancel)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if this.CallCancel != that.CallCancel {
		return false
	}
	return true
}

func (this *CallStart) EqualVT(that *CallStart) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.RpcService != that.RpcService {
		return false
	}
	if this.RpcMethod != that.RpcMethod {
		return false
	}
	if string(this.Data) != string(that.Data) {
		return false
	}
	if this.DataIsZero != that.DataIsZero {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *CallStart) EqualMessageVT(thatMsg any) bool {
	that, ok := thatMsg.(*CallStart)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *CallData) EqualVT(that *CallData) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if string(this.Data) != string(that.Data) {
		return false
	}
	if this.DataIsZero != that.DataIsZero {
		return false
	}
	if this.Complete != that.Complete {
		return false
	}
	if this.Error != that.Error {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *CallData) EqualMessageVT(thatMsg any) bool {
	that, ok := thatMsg.(*CallData)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}

// MarshalProtoJSON marshals the Packet message to JSON.
func (x *Packet) MarshalProtoJSON(s *json.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.Body != nil {
		switch ov := x.Body.(type) {
		case *Packet_CallStart:
			s.WriteMoreIf(&wroteField)
			s.WriteObjectField("callStart")
			ov.CallStart.MarshalProtoJSON(s.WithField("callStart"))
		case *Packet_CallData:
			s.WriteMoreIf(&wroteField)
			s.WriteObjectField("callData")
			ov.CallData.MarshalProtoJSON(s.WithField("callData"))
		case *Packet_CallCancel:
			s.WriteMoreIf(&wroteField)
			s.WriteObjectField("callCancel")
			s.WriteBool(ov.CallCancel)
		}
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the Packet to JSON.
func (x *Packet) MarshalJSON() ([]byte, error) {
	return json.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the Packet message from JSON.
func (x *Packet) UnmarshalProtoJSON(s *json.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.Skip() // ignore unknown field
		case "call_start", "callStart":
			ov := &Packet_CallStart{}
			x.Body = ov
			if s.ReadNil() {
				ov.CallStart = nil
				return
			}
			ov.CallStart = &CallStart{}
			ov.CallStart.UnmarshalProtoJSON(s.WithField("call_start", true))
		case "call_data", "callData":
			ov := &Packet_CallData{}
			x.Body = ov
			if s.ReadNil() {
				ov.CallData = nil
				return
			}
			ov.CallData = &CallData{}
			ov.CallData.UnmarshalProtoJSON(s.WithField("call_data", true))
		case "call_cancel", "callCancel":
			s.AddField("call_cancel")
			ov := &Packet_CallCancel{}
			x.Body = ov
			ov.CallCancel = s.ReadBool()
		}
	})
}

// UnmarshalJSON unmarshals the Packet from JSON.
func (x *Packet) UnmarshalJSON(b []byte) error {
	return json.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the CallStart message to JSON.
func (x *CallStart) MarshalProtoJSON(s *json.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.RpcService != "" || s.HasField("rpcService") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("rpcService")
		s.WriteString(x.RpcService)
	}
	if x.RpcMethod != "" || s.HasField("rpcMethod") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("rpcMethod")
		s.WriteString(x.RpcMethod)
	}
	if len(x.Data) > 0 || s.HasField("data") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("data")
		s.WriteBytes(x.Data)
	}
	if x.DataIsZero || s.HasField("dataIsZero") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("dataIsZero")
		s.WriteBool(x.DataIsZero)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the CallStart to JSON.
func (x *CallStart) MarshalJSON() ([]byte, error) {
	return json.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the CallStart message from JSON.
func (x *CallStart) UnmarshalProtoJSON(s *json.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.Skip() // ignore unknown field
		case "rpc_service", "rpcService":
			s.AddField("rpc_service")
			x.RpcService = s.ReadString()
		case "rpc_method", "rpcMethod":
			s.AddField("rpc_method")
			x.RpcMethod = s.ReadString()
		case "data":
			s.AddField("data")
			x.Data = s.ReadBytes()
		case "data_is_zero", "dataIsZero":
			s.AddField("data_is_zero")
			x.DataIsZero = s.ReadBool()
		}
	})
}

// UnmarshalJSON unmarshals the CallStart from JSON.
func (x *CallStart) UnmarshalJSON(b []byte) error {
	return json.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the CallData message to JSON.
func (x *CallData) MarshalProtoJSON(s *json.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if len(x.Data) > 0 || s.HasField("data") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("data")
		s.WriteBytes(x.Data)
	}
	if x.DataIsZero || s.HasField("dataIsZero") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("dataIsZero")
		s.WriteBool(x.DataIsZero)
	}
	if x.Complete || s.HasField("complete") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("complete")
		s.WriteBool(x.Complete)
	}
	if x.Error != "" || s.HasField("error") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("error")
		s.WriteString(x.Error)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the CallData to JSON.
func (x *CallData) MarshalJSON() ([]byte, error) {
	return json.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the CallData message from JSON.
func (x *CallData) UnmarshalProtoJSON(s *json.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.Skip() // ignore unknown field
		case "data":
			s.AddField("data")
			x.Data = s.ReadBytes()
		case "data_is_zero", "dataIsZero":
			s.AddField("data_is_zero")
			x.DataIsZero = s.ReadBool()
		case "complete":
			s.AddField("complete")
			x.Complete = s.ReadBool()
		case "error":
			s.AddField("error")
			x.Error = s.ReadString()
		}
	})
}

// UnmarshalJSON unmarshals the CallData from JSON.
func (x *CallData) UnmarshalJSON(b []byte) error {
	return json.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

func (m *Packet) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Packet) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *Packet) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if vtmsg, ok := m.Body.(interface {
		MarshalToSizedBufferVT([]byte) (int, error)
	}); ok {
		size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	return len(dAtA) - i, nil
}

func (m *Packet_CallStart) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *Packet_CallStart) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.CallStart != nil {
		size, err := m.CallStart.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protobuf_go_lite.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0xa
	} else {
		i = protobuf_go_lite.EncodeVarint(dAtA, i, 0)
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *Packet_CallData) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *Packet_CallData) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.CallData != nil {
		size, err := m.CallData.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protobuf_go_lite.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x12
	} else {
		i = protobuf_go_lite.EncodeVarint(dAtA, i, 0)
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *Packet_CallCancel) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *Packet_CallCancel) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	i--
	if m.CallCancel {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i--
	dAtA[i] = 0x18
	return len(dAtA) - i, nil
}
func (m *CallStart) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CallStart) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *CallStart) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if m.DataIsZero {
		i--
		if m.DataIsZero {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = protobuf_go_lite.EncodeVarint(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.RpcMethod) > 0 {
		i -= len(m.RpcMethod)
		copy(dAtA[i:], m.RpcMethod)
		i = protobuf_go_lite.EncodeVarint(dAtA, i, uint64(len(m.RpcMethod)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.RpcService) > 0 {
		i -= len(m.RpcService)
		copy(dAtA[i:], m.RpcService)
		i = protobuf_go_lite.EncodeVarint(dAtA, i, uint64(len(m.RpcService)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CallData) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CallData) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *CallData) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if len(m.Error) > 0 {
		i -= len(m.Error)
		copy(dAtA[i:], m.Error)
		i = protobuf_go_lite.EncodeVarint(dAtA, i, uint64(len(m.Error)))
		i--
		dAtA[i] = 0x22
	}
	if m.Complete {
		i--
		if m.Complete {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if m.DataIsZero {
		i--
		if m.DataIsZero {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = protobuf_go_lite.EncodeVarint(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Packet) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if vtmsg, ok := m.Body.(interface{ SizeVT() int }); ok {
		n += vtmsg.SizeVT()
	}
	n += len(m.unknownFields)
	return n
}

func (m *Packet_CallStart) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CallStart != nil {
		l = m.CallStart.SizeVT()
		n += 1 + l + protobuf_go_lite.SizeOfVarint(uint64(l))
	} else {
		n += 2
	}
	return n
}
func (m *Packet_CallData) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CallData != nil {
		l = m.CallData.SizeVT()
		n += 1 + l + protobuf_go_lite.SizeOfVarint(uint64(l))
	} else {
		n += 2
	}
	return n
}
func (m *Packet_CallCancel) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 2
	return n
}
func (m *CallStart) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RpcService)
	if l > 0 {
		n += 1 + l + protobuf_go_lite.SizeOfVarint(uint64(l))
	}
	l = len(m.RpcMethod)
	if l > 0 {
		n += 1 + l + protobuf_go_lite.SizeOfVarint(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + protobuf_go_lite.SizeOfVarint(uint64(l))
	}
	if m.DataIsZero {
		n += 2
	}
	n += len(m.unknownFields)
	return n
}

func (m *CallData) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + protobuf_go_lite.SizeOfVarint(uint64(l))
	}
	if m.DataIsZero {
		n += 2
	}
	if m.Complete {
		n += 2
	}
	l = len(m.Error)
	if l > 0 {
		n += 1 + l + protobuf_go_lite.SizeOfVarint(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (x *Packet) MarshalProtoText() string {
	var sb strings.Builder
	sb.WriteString("Packet {")
	switch body := x.Body.(type) {
	case *Packet_CallStart:
		if body.CallStart != nil {
			if sb.Len() > 8 {
				sb.WriteString(" ")
			}
			sb.WriteString("call_start: ")
			sb.WriteString(body.CallStart.MarshalProtoText())
		}
	case *Packet_CallData:
		if body.CallData != nil {
			if sb.Len() > 8 {
				sb.WriteString(" ")
			}
			sb.WriteString("call_data: ")
			sb.WriteString(body.CallData.MarshalProtoText())
		}
	case *Packet_CallCancel:
		if body.CallCancel != false {
			if sb.Len() > 8 {
				sb.WriteString(" ")
			}
			sb.WriteString("call_cancel: ")
			sb.WriteString(strconv.FormatBool(body.CallCancel))
		}
	}
	sb.WriteString("}")
	return sb.String()
}

func (x *Packet) String() string {
	return x.MarshalProtoText()
}
func (x *CallStart) MarshalProtoText() string {
	var sb strings.Builder
	sb.WriteString("CallStart {")
	if x.RpcService != "" {
		if sb.Len() > 11 {
			sb.WriteString(" ")
		}
		sb.WriteString("rpc_service: ")
		sb.WriteString(strconv.Quote(x.RpcService))
	}
	if x.RpcMethod != "" {
		if sb.Len() > 11 {
			sb.WriteString(" ")
		}
		sb.WriteString("rpc_method: ")
		sb.WriteString(strconv.Quote(x.RpcMethod))
	}
	if x.Data != nil {
		if sb.Len() > 11 {
			sb.WriteString(" ")
		}
		sb.WriteString("data: ")
		sb.WriteString("\"")
		sb.WriteString(base64.StdEncoding.EncodeToString(x.Data))
		sb.WriteString("\"")
	}
	if x.DataIsZero != false {
		if sb.Len() > 11 {
			sb.WriteString(" ")
		}
		sb.WriteString("data_is_zero: ")
		sb.WriteString(strconv.FormatBool(x.DataIsZero))
	}
	sb.WriteString("}")
	return sb.String()
}

func (x *CallStart) String() string {
	return x.MarshalProtoText()
}
func (x *CallData) MarshalProtoText() string {
	var sb strings.Builder
	sb.WriteString("CallData {")
	if x.Data != nil {
		if sb.Len() > 10 {
			sb.WriteString(" ")
		}
		sb.WriteString("data: ")
		sb.WriteString("\"")
		sb.WriteString(base64.StdEncoding.EncodeToString(x.Data))
		sb.WriteString("\"")
	}
	if x.DataIsZero != false {
		if sb.Len() > 10 {
			sb.WriteString(" ")
		}
		sb.WriteString("data_is_zero: ")
		sb.WriteString(strconv.FormatBool(x.DataIsZero))
	}
	if x.Complete != false {
		if sb.Len() > 10 {
			sb.WriteString(" ")
		}
		sb.WriteString("complete: ")
		sb.WriteString(strconv.FormatBool(x.Complete))
	}
	if x.Error != "" {
		if sb.Len() > 10 {
			sb.WriteString(" ")
		}
		sb.WriteString("error: ")
		sb.WriteString(strconv.Quote(x.Error))
	}
	sb.WriteString("}")
	return sb.String()
}

func (x *CallData) String() string {
	return x.MarshalProtoText()
}
func (m *Packet) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return protobuf_go_lite.ErrIntOverflow
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Packet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Packet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CallStart", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protobuf_go_lite.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return protobuf_go_lite.ErrInvalidLength
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return protobuf_go_lite.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if oneof, ok := m.Body.(*Packet_CallStart); ok {
				if err := oneof.CallStart.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
					return err
				}
			} else {
				v := &CallStart{}
				if err := v.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
					return err
				}
				m.Body = &Packet_CallStart{CallStart: v}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CallData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protobuf_go_lite.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return protobuf_go_lite.ErrInvalidLength
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return protobuf_go_lite.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if oneof, ok := m.Body.(*Packet_CallData); ok {
				if err := oneof.CallData.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
					return err
				}
			} else {
				v := &CallData{}
				if err := v.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
					return err
				}
				m.Body = &Packet_CallData{CallData: v}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CallCancel", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protobuf_go_lite.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.Body = &Packet_CallCancel{CallCancel: b}
		default:
			iNdEx = preIndex
			skippy, err := protobuf_go_lite.Skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return protobuf_go_lite.ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CallStart) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return protobuf_go_lite.ErrIntOverflow
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CallStart: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CallStart: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RpcService", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protobuf_go_lite.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return protobuf_go_lite.ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return protobuf_go_lite.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RpcService = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RpcMethod", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protobuf_go_lite.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return protobuf_go_lite.ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return protobuf_go_lite.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RpcMethod = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protobuf_go_lite.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return protobuf_go_lite.ErrInvalidLength
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return protobuf_go_lite.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataIsZero", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protobuf_go_lite.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.DataIsZero = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := protobuf_go_lite.Skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return protobuf_go_lite.ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CallData) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return protobuf_go_lite.ErrIntOverflow
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CallData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CallData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protobuf_go_lite.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return protobuf_go_lite.ErrInvalidLength
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return protobuf_go_lite.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataIsZero", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protobuf_go_lite.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.DataIsZero = bool(v != 0)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Complete", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protobuf_go_lite.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Complete = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protobuf_go_lite.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return protobuf_go_lite.ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return protobuf_go_lite.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Error = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := protobuf_go_lite.Skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return protobuf_go_lite.ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
