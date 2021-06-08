// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: util/tracing/tracingservicepb/tracing_service.proto

package tracingservicepb

import (
	context "context"
	fmt "fmt"
	tracingpb "github.com/cockroachdb/cockroach/pkg/util/tracing/tracingpb"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type SpanRecordingRequest struct {
	TraceID uint64 `protobuf:"varint,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
}

func (m *SpanRecordingRequest) Reset()         { *m = SpanRecordingRequest{} }
func (m *SpanRecordingRequest) String() string { return proto.CompactTextString(m) }
func (*SpanRecordingRequest) ProtoMessage()    {}
func (*SpanRecordingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_29b78bec82996ca3, []int{0}
}
func (m *SpanRecordingRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SpanRecordingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *SpanRecordingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpanRecordingRequest.Merge(m, src)
}
func (m *SpanRecordingRequest) XXX_Size() int {
	return m.Size()
}
func (m *SpanRecordingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SpanRecordingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SpanRecordingRequest proto.InternalMessageInfo

type SpanRecordingResponse struct {
	SpanRecordings []tracingpb.RecordedSpan `protobuf:"bytes,1,rep,name=span_recordings,json=spanRecordings,proto3" json:"span_recordings"`
}

func (m *SpanRecordingResponse) Reset()         { *m = SpanRecordingResponse{} }
func (m *SpanRecordingResponse) String() string { return proto.CompactTextString(m) }
func (*SpanRecordingResponse) ProtoMessage()    {}
func (*SpanRecordingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_29b78bec82996ca3, []int{1}
}
func (m *SpanRecordingResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SpanRecordingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *SpanRecordingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpanRecordingResponse.Merge(m, src)
}
func (m *SpanRecordingResponse) XXX_Size() int {
	return m.Size()
}
func (m *SpanRecordingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SpanRecordingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SpanRecordingResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SpanRecordingRequest)(nil), "cockroach.util.tracing.SpanRecordingRequest")
	proto.RegisterType((*SpanRecordingResponse)(nil), "cockroach.util.tracing.SpanRecordingResponse")
}

func init() {
	proto.RegisterFile("util/tracing/tracingservicepb/tracing_service.proto", fileDescriptor_29b78bec82996ca3)
}

var fileDescriptor_29b78bec82996ca3 = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0x4d, 0x4b, 0xc3, 0x30,
	0x00, 0x4d, 0x70, 0x38, 0xc9, 0xc0, 0x8f, 0x32, 0x65, 0xec, 0x90, 0x8d, 0x1d, 0x64, 0x88, 0x66,
	0xd0, 0xdd, 0x3d, 0x14, 0x41, 0x76, 0xad, 0x9e, 0x04, 0x29, 0x6d, 0x1a, 0x6a, 0x50, 0x92, 0x98,
	0x74, 0xbb, 0xf9, 0x1f, 0xfc, 0x59, 0x3d, 0xee, 0xb8, 0xd3, 0xd0, 0xf6, 0x8f, 0x48, 0xfa, 0x21,
	0x3a, 0x2a, 0x78, 0x6a, 0x79, 0x1f, 0x79, 0xef, 0x25, 0x68, 0xbe, 0x4c, 0xf9, 0xcb, 0x2c, 0xd5,
	0x21, 0xe5, 0x22, 0x69, 0xbe, 0x86, 0xe9, 0x15, 0xa7, 0x4c, 0x45, 0x0d, 0x10, 0xd4, 0x08, 0x51,
	0x5a, 0xa6, 0xd2, 0x39, 0xa3, 0x92, 0x3e, 0x6b, 0x19, 0xd2, 0x27, 0x62, 0xed, 0xa4, 0x56, 0x0d,
	0xfb, 0x89, 0x4c, 0x64, 0x29, 0x99, 0xd9, 0xbf, 0x4a, 0x3d, 0xbc, 0x68, 0x8b, 0x50, 0xd1, 0x4c,
	0x33, 0x2a, 0x75, 0xcc, 0xe2, 0xc0, 0xa8, 0x50, 0x54, 0xda, 0xc9, 0x35, 0xea, 0xdf, 0xa9, 0x50,
	0xf8, 0x25, 0xc5, 0x45, 0xe2, 0xb3, 0xd7, 0x25, 0x33, 0xa9, 0x73, 0x8e, 0x0e, 0xac, 0x91, 0x05,
	0x3c, 0x1e, 0xc0, 0x31, 0x9c, 0x76, 0xbc, 0x5e, 0xbe, 0x1d, 0x75, 0xef, 0x2d, 0xb6, 0xb8, 0xf1,
	0xbb, 0x25, 0xb9, 0x88, 0x27, 0x2b, 0x74, 0xba, 0xe3, 0x37, 0x4a, 0x0a, 0xc3, 0x9c, 0x47, 0x74,
	0x64, 0x63, 0x02, 0xdd, 0x30, 0x66, 0x00, 0xc7, 0x7b, 0xd3, 0x9e, 0x4b, 0x48, 0xfb, 0x18, 0xf2,
	0x5d, 0x94, 0xf8, 0x75, 0x51, 0x7b, 0xb2, 0xd7, 0xc9, 0xb6, 0x23, 0xe0, 0x1f, 0x9a, 0x9f, 0x29,
	0xc6, 0x7d, 0x43, 0x65, 0x17, 0x2e, 0x12, 0x47, 0xa3, 0x93, 0x5b, 0x96, 0xfe, 0x6a, 0x61, 0x9c,
	0xcb, 0xbf, 0x52, 0xda, 0xd6, 0x0e, 0xaf, 0xfe, 0xa9, 0xae, 0xb6, 0x4d, 0x80, 0xe7, 0x66, 0x9f,
	0x18, 0x64, 0x39, 0x86, 0xeb, 0x1c, 0xc3, 0x4d, 0x8e, 0xe1, 0x47, 0x8e, 0xe1, 0x7b, 0x81, 0xc1,
	0xba, 0xc0, 0x60, 0x53, 0x60, 0xf0, 0x70, 0xbc, 0xfb, 0xb4, 0xd1, 0x7e, 0x79, 0xe3, 0xf3, 0xaf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xe4, 0xe8, 0x32, 0x41, 0x02, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TracingClient is the client API for Tracing service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TracingClient interface {
	GetSpanRecordings(ctx context.Context, in *SpanRecordingRequest, opts ...grpc.CallOption) (*SpanRecordingResponse, error)
}

type tracingClient struct {
	cc *grpc.ClientConn
}

func NewTracingClient(cc *grpc.ClientConn) TracingClient {
	return &tracingClient{cc}
}

func (c *tracingClient) GetSpanRecordings(ctx context.Context, in *SpanRecordingRequest, opts ...grpc.CallOption) (*SpanRecordingResponse, error) {
	out := new(SpanRecordingResponse)
	err := c.cc.Invoke(ctx, "/cockroach.util.tracing.Tracing/GetSpanRecordings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TracingServer is the server API for Tracing service.
type TracingServer interface {
	GetSpanRecordings(context.Context, *SpanRecordingRequest) (*SpanRecordingResponse, error)
}

// UnimplementedTracingServer can be embedded to have forward compatible implementations.
type UnimplementedTracingServer struct {
}

func (*UnimplementedTracingServer) GetSpanRecordings(ctx context.Context, req *SpanRecordingRequest) (*SpanRecordingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSpanRecordings not implemented")
}

func RegisterTracingServer(s *grpc.Server, srv TracingServer) {
	s.RegisterService(&_Tracing_serviceDesc, srv)
}

func _Tracing_GetSpanRecordings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SpanRecordingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TracingServer).GetSpanRecordings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cockroach.util.tracing.Tracing/GetSpanRecordings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TracingServer).GetSpanRecordings(ctx, req.(*SpanRecordingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Tracing_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cockroach.util.tracing.Tracing",
	HandlerType: (*TracingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSpanRecordings",
			Handler:    _Tracing_GetSpanRecordings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "util/tracing/tracingservicepb/tracing_service.proto",
}

func (m *SpanRecordingRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SpanRecordingRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SpanRecordingRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TraceID != 0 {
		i = encodeVarintTracingService(dAtA, i, uint64(m.TraceID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *SpanRecordingResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SpanRecordingResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SpanRecordingResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SpanRecordings) > 0 {
		for iNdEx := len(m.SpanRecordings) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SpanRecordings[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTracingService(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintTracingService(dAtA []byte, offset int, v uint64) int {
	offset -= sovTracingService(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SpanRecordingRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TraceID != 0 {
		n += 1 + sovTracingService(uint64(m.TraceID))
	}
	return n
}

func (m *SpanRecordingResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.SpanRecordings) > 0 {
		for _, e := range m.SpanRecordings {
			l = e.Size()
			n += 1 + l + sovTracingService(uint64(l))
		}
	}
	return n
}

func sovTracingService(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTracingService(x uint64) (n int) {
	return sovTracingService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SpanRecordingRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTracingService
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
			return fmt.Errorf("proto: SpanRecordingRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SpanRecordingRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TraceID", wireType)
			}
			m.TraceID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTracingService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TraceID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTracingService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTracingService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SpanRecordingResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTracingService
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
			return fmt.Errorf("proto: SpanRecordingResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SpanRecordingResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpanRecordings", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTracingService
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
				return ErrInvalidLengthTracingService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTracingService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SpanRecordings = append(m.SpanRecordings, tracingpb.RecordedSpan{})
			if err := m.SpanRecordings[len(m.SpanRecordings)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTracingService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTracingService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTracingService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTracingService
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTracingService
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTracingService
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTracingService
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTracingService
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTracingService
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTracingService        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTracingService          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTracingService = fmt.Errorf("proto: unexpected end of group")
)