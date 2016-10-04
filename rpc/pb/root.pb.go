// Code generated by protoc-gen-go.
// source: root.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	root.proto

It has these top-level messages:
	Benchmark
	WriteStatus
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Benchmark struct {
	Project string `protobuf:"bytes,1,opt,name=project" json:"project,omitempty"`
	// Types that are valid to be assigned to Kind:
	//	*Benchmark_GoTestBench_
	Kind isBenchmark_Kind `protobuf_oneof:"kind"`
}

func (m *Benchmark) Reset()                    { *m = Benchmark{} }
func (m *Benchmark) String() string            { return proto.CompactTextString(m) }
func (*Benchmark) ProtoMessage()               {}
func (*Benchmark) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isBenchmark_Kind interface {
	isBenchmark_Kind()
}

type Benchmark_GoTestBench_ struct {
	GoTestBench *Benchmark_GoTestBench `protobuf:"bytes,2,opt,name=goTestBench,oneof"`
}

func (*Benchmark_GoTestBench_) isBenchmark_Kind() {}

func (m *Benchmark) GetKind() isBenchmark_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (m *Benchmark) GetGoTestBench() *Benchmark_GoTestBench {
	if x, ok := m.GetKind().(*Benchmark_GoTestBench_); ok {
		return x.GoTestBench
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Benchmark) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Benchmark_OneofMarshaler, _Benchmark_OneofUnmarshaler, _Benchmark_OneofSizer, []interface{}{
		(*Benchmark_GoTestBench_)(nil),
	}
}

func _Benchmark_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Benchmark)
	// kind
	switch x := m.Kind.(type) {
	case *Benchmark_GoTestBench_:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.GoTestBench); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Benchmark.Kind has unexpected type %T", x)
	}
	return nil
}

func _Benchmark_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Benchmark)
	switch tag {
	case 2: // kind.goTestBench
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Benchmark_GoTestBench)
		err := b.DecodeMessage(msg)
		m.Kind = &Benchmark_GoTestBench_{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Benchmark_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Benchmark)
	// kind
	switch x := m.Kind.(type) {
	case *Benchmark_GoTestBench_:
		s := proto.Size(x.GoTestBench)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// metric types
type Benchmark_GoTestBench struct {
	Name              string  `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Package           string  `protobuf:"bytes,2,opt,name=package" json:"package,omitempty"`
	N                 uint64  `protobuf:"varint,3,opt,name=n" json:"n,omitempty"`
	NsPerOp           float64 `protobuf:"fixed64,4,opt,name=nsPerOp" json:"nsPerOp,omitempty"`
	AllocedBytesPerOp uint64  `protobuf:"varint,5,opt,name=allocedBytesPerOp" json:"allocedBytesPerOp,omitempty"`
	AllocsPerOp       uint64  `protobuf:"varint,6,opt,name=allocsPerOp" json:"allocsPerOp,omitempty"`
	MbPerS            float64 `protobuf:"fixed64,7,opt,name=mbPerS" json:"mbPerS,omitempty"`
	Measured          int64   `protobuf:"varint,8,opt,name=measured" json:"measured,omitempty"`
}

func (m *Benchmark_GoTestBench) Reset()                    { *m = Benchmark_GoTestBench{} }
func (m *Benchmark_GoTestBench) String() string            { return proto.CompactTextString(m) }
func (*Benchmark_GoTestBench) ProtoMessage()               {}
func (*Benchmark_GoTestBench) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type WriteStatus struct {
	// Types that are valid to be assigned to Status:
	//	*WriteStatus_Stats_
	//	*WriteStatus_Error
	Status isWriteStatus_Status `protobuf_oneof:"status"`
}

func (m *WriteStatus) Reset()                    { *m = WriteStatus{} }
func (m *WriteStatus) String() string            { return proto.CompactTextString(m) }
func (*WriteStatus) ProtoMessage()               {}
func (*WriteStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type isWriteStatus_Status interface {
	isWriteStatus_Status()
}

type WriteStatus_Stats_ struct {
	Stats *WriteStatus_Stats `protobuf:"bytes,1,opt,name=stats,oneof"`
}
type WriteStatus_Error struct {
	Error string `protobuf:"bytes,2,opt,name=error,oneof"`
}

func (*WriteStatus_Stats_) isWriteStatus_Status() {}
func (*WriteStatus_Error) isWriteStatus_Status()  {}

func (m *WriteStatus) GetStatus() isWriteStatus_Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *WriteStatus) GetStats() *WriteStatus_Stats {
	if x, ok := m.GetStatus().(*WriteStatus_Stats_); ok {
		return x.Stats
	}
	return nil
}

func (m *WriteStatus) GetError() string {
	if x, ok := m.GetStatus().(*WriteStatus_Error); ok {
		return x.Error
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*WriteStatus) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _WriteStatus_OneofMarshaler, _WriteStatus_OneofUnmarshaler, _WriteStatus_OneofSizer, []interface{}{
		(*WriteStatus_Stats_)(nil),
		(*WriteStatus_Error)(nil),
	}
}

func _WriteStatus_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*WriteStatus)
	// status
	switch x := m.Status.(type) {
	case *WriteStatus_Stats_:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Stats); err != nil {
			return err
		}
	case *WriteStatus_Error:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Error)
	case nil:
	default:
		return fmt.Errorf("WriteStatus.Status has unexpected type %T", x)
	}
	return nil
}

func _WriteStatus_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*WriteStatus)
	switch tag {
	case 1: // status.stats
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(WriteStatus_Stats)
		err := b.DecodeMessage(msg)
		m.Status = &WriteStatus_Stats_{msg}
		return true, err
	case 2: // status.error
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Status = &WriteStatus_Error{x}
		return true, err
	default:
		return false, nil
	}
}

func _WriteStatus_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*WriteStatus)
	// status
	switch x := m.Status.(type) {
	case *WriteStatus_Stats_:
		s := proto.Size(x.Stats)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *WriteStatus_Error:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Error)))
		n += len(x.Error)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type WriteStatus_Stats struct {
	Written uint64 `protobuf:"varint,1,opt,name=written" json:"written,omitempty"`
}

func (m *WriteStatus_Stats) Reset()                    { *m = WriteStatus_Stats{} }
func (m *WriteStatus_Stats) String() string            { return proto.CompactTextString(m) }
func (*WriteStatus_Stats) ProtoMessage()               {}
func (*WriteStatus_Stats) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func init() {
	proto.RegisterType((*Benchmark)(nil), "pb.Benchmark")
	proto.RegisterType((*Benchmark_GoTestBench)(nil), "pb.Benchmark.GoTestBench")
	proto.RegisterType((*WriteStatus)(nil), "pb.WriteStatus")
	proto.RegisterType((*WriteStatus_Stats)(nil), "pb.WriteStatus.Stats")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Ingester service

type IngesterClient interface {
	AddBenchmark(ctx context.Context, opts ...grpc.CallOption) (Ingester_AddBenchmarkClient, error)
}

type ingesterClient struct {
	cc *grpc.ClientConn
}

func NewIngesterClient(cc *grpc.ClientConn) IngesterClient {
	return &ingesterClient{cc}
}

func (c *ingesterClient) AddBenchmark(ctx context.Context, opts ...grpc.CallOption) (Ingester_AddBenchmarkClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Ingester_serviceDesc.Streams[0], c.cc, "/pb.Ingester/AddBenchmark", opts...)
	if err != nil {
		return nil, err
	}
	x := &ingesterAddBenchmarkClient{stream}
	return x, nil
}

type Ingester_AddBenchmarkClient interface {
	Send(*Benchmark) error
	CloseAndRecv() (*WriteStatus, error)
	grpc.ClientStream
}

type ingesterAddBenchmarkClient struct {
	grpc.ClientStream
}

func (x *ingesterAddBenchmarkClient) Send(m *Benchmark) error {
	return x.ClientStream.SendMsg(m)
}

func (x *ingesterAddBenchmarkClient) CloseAndRecv() (*WriteStatus, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(WriteStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Ingester service

type IngesterServer interface {
	AddBenchmark(Ingester_AddBenchmarkServer) error
}

func RegisterIngesterServer(s *grpc.Server, srv IngesterServer) {
	s.RegisterService(&_Ingester_serviceDesc, srv)
}

func _Ingester_AddBenchmark_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(IngesterServer).AddBenchmark(&ingesterAddBenchmarkServer{stream})
}

type Ingester_AddBenchmarkServer interface {
	SendAndClose(*WriteStatus) error
	Recv() (*Benchmark, error)
	grpc.ServerStream
}

type ingesterAddBenchmarkServer struct {
	grpc.ServerStream
}

func (x *ingesterAddBenchmarkServer) SendAndClose(m *WriteStatus) error {
	return x.ServerStream.SendMsg(m)
}

func (x *ingesterAddBenchmarkServer) Recv() (*Benchmark, error) {
	m := new(Benchmark)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Ingester_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Ingester",
	HandlerType: (*IngesterServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "AddBenchmark",
			Handler:       _Ingester_AddBenchmark_Handler,
			ClientStreams: true,
		},
	},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("root.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 397 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x92, 0x41, 0xaf, 0x93, 0x40,
	0x10, 0xc7, 0xdf, 0xf2, 0x80, 0xc7, 0x1b, 0x9e, 0x31, 0x6e, 0x62, 0x45, 0xa2, 0x09, 0x72, 0x22,
	0x46, 0x69, 0xac, 0x37, 0x13, 0x0f, 0x72, 0xb1, 0x9e, 0xda, 0x50, 0x13, 0x0f, 0x9e, 0x16, 0x98,
	0x20, 0xb6, 0xec, 0x92, 0xdd, 0xad, 0xc6, 0xa3, 0x7e, 0x05, 0x3f, 0x9a, 0x89, 0x9f, 0xc0, 0xc4,
	0xaf, 0x61, 0x58, 0xa0, 0x45, 0x3d, 0xc1, 0x7f, 0xe6, 0x37, 0xf3, 0x87, 0x99, 0x01, 0x90, 0x42,
	0xe8, 0xb4, 0x93, 0x42, 0x0b, 0x6a, 0x75, 0x45, 0xf8, 0xa0, 0x16, 0xa2, 0x3e, 0xe0, 0x92, 0x75,
	0xcd, 0x92, 0x71, 0x2e, 0x34, 0xd3, 0x8d, 0xe0, 0x6a, 0x20, 0xe2, 0x9f, 0x16, 0x5c, 0x67, 0xc8,
	0xcb, 0x0f, 0x2d, 0x93, 0x7b, 0x1a, 0xc0, 0x55, 0x27, 0xc5, 0x47, 0x2c, 0x75, 0x40, 0x22, 0x92,
	0x5c, 0xe7, 0x93, 0xa4, 0x2f, 0xc1, 0xaf, 0xc5, 0x5b, 0x54, 0xda, 0xc0, 0x81, 0x15, 0x91, 0xc4,
	0x5f, 0xdd, 0x4f, 0xbb, 0x22, 0x3d, 0x55, 0xa7, 0xaf, 0xcf, 0xc0, 0xfa, 0x22, 0x9f, 0xf3, 0xe1,
	0x6f, 0x02, 0xfe, 0x2c, 0x4d, 0x29, 0xd8, 0x9c, 0xb5, 0x38, 0xba, 0x98, 0x77, 0x63, 0xce, 0xca,
	0x3d, 0xab, 0xd1, 0xb4, 0xef, 0xcd, 0x07, 0x49, 0x6f, 0x80, 0xf0, 0xe0, 0x32, 0x22, 0x89, 0x9d,
	0x13, 0xde, 0x73, 0x5c, 0x6d, 0x51, 0x6e, 0xba, 0xc0, 0x8e, 0x48, 0x42, 0xf2, 0x49, 0xd2, 0x27,
	0x70, 0x87, 0x1d, 0x0e, 0xa2, 0xc4, 0x2a, 0xfb, 0xa2, 0x71, 0x64, 0x1c, 0x53, 0xf7, 0x7f, 0x82,
	0x46, 0xe0, 0x9b, 0xe0, 0xc8, 0xb9, 0x86, 0x9b, 0x87, 0xe8, 0x02, 0xdc, 0xb6, 0xd8, 0xa2, 0xdc,
	0x05, 0x57, 0xc6, 0x68, 0x54, 0x34, 0x04, 0xaf, 0x45, 0xa6, 0x8e, 0x12, 0xab, 0xc0, 0x8b, 0x48,
	0x72, 0x99, 0x9f, 0x74, 0xe6, 0x82, 0xbd, 0x6f, 0x78, 0x15, 0x7f, 0x25, 0xe0, 0xbf, 0x93, 0x8d,
	0xc6, 0x9d, 0x66, 0xfa, 0xa8, 0xe8, 0x53, 0x70, 0x94, 0x66, 0x5a, 0x99, 0x5f, 0xf6, 0x57, 0x77,
	0xfb, 0xd1, 0xcd, 0xf2, 0x69, 0xff, 0x50, 0xeb, 0x8b, 0x7c, 0xa0, 0xe8, 0x02, 0x1c, 0x94, 0x52,
	0xc8, 0x61, 0x14, 0x7d, 0xdc, 0xc8, 0xf0, 0x11, 0x38, 0x86, 0xec, 0xa7, 0xf0, 0x59, 0x36, 0x5a,
	0x23, 0x37, 0x1d, 0xed, 0x7c, 0x92, 0x99, 0x07, 0xae, 0x32, 0x3d, 0x57, 0xef, 0xc1, 0x7b, 0xc3,
	0x6b, 0x54, 0x1a, 0x25, 0xdd, 0xc0, 0xcd, 0xab, 0xaa, 0x3a, 0xaf, 0xfa, 0xd6, 0x5f, 0xbb, 0x0b,
	0x6f, 0xff, 0xf3, 0x3d, 0xf1, 0xc3, 0x6f, 0x3f, 0x7e, 0x7d, 0xb7, 0xee, 0xc5, 0xd4, 0x1c, 0xce,
	0xa7, 0x67, 0xcb, 0x62, 0x62, 0xd5, 0x0b, 0xf2, 0x38, 0x21, 0x85, 0x6b, 0x0e, 0xe8, 0xf9, 0x9f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x1c, 0x6b, 0x95, 0x0f, 0x70, 0x02, 0x00, 0x00,
}
